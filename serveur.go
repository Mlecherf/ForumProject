package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var db = initDatabase("database.db")

type Cookie struct {
	User       string
	Mail       string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

var tpl *template.Template

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	tpl, _ = tpl.ParseGlob("template/*.html")
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/profile", userprofile)
	http.HandleFunc("/post", userpost)
	http.HandleFunc("/theme", usertheme)
	style := http.FileServer(http.Dir("asset/style/"))
	image := http.FileServer(http.Dir("asset/image/"))
	js := http.FileServer(http.Dir("asset/js/"))

	http.Handle("/static/style/", http.StripPrefix("/static/style/", style))
	http.Handle("/static/image/", http.StripPrefix("/static/image/", image))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", js))

	http.ListenAndServe(":8060", nil)
}

func home(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "/home.html", nil)
}

func register(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("Email__input")
	username := request.FormValue("Username__input")
	password := request.FormValue("Password__input")
	passwordverif := request.FormValue("Password_Verif__input")

	like := 0
	post := 0

	if len(email) != 0 || len(username) != 0 || len(password) != 0 || len(passwordverif) != 0 {
		println("=================")
		println("EMAIL :", email)
		println("USERNAME :", username)
		println("PASSWORD :", password)
		println("PASSWORDVERIF :", passwordverif)
		println("=================")
		alltable := selectAllFromTable(db, "users")
		Arr := VerifRegister(alltable)
		Wrong := false

		if len(Arr) != 0 {
			for i := 0; i < len(Arr); i++ {
				if Arr[i] == username || Arr[i] == email || Arr[i] == password {
					Wrong = true
				}
			}

			if Wrong == false {
				fmt.Println("Created... new enter in the DB.")
				insertIntoUsers(db, username, email, password, like, post)
				alltable2 := selectAllFromTable(db, "users")
				displayUsersRow(alltable2)
			} else {
				fmt.Println("Cant create... Already exist in the DB.")
			}

		} else {
			insertIntoUsers(db, username, email, password, like, post)
			alltable1 := selectAllFromTable(db, "users")
			displayUsersRow(alltable1)
		}
	}

	tpl.ExecuteTemplate(response, "register.html", nil)
}

func login(response http.ResponseWriter, request *http.Request) {
	connected := false
	http.SetCookie(response, &http.Cookie{
		Name:     "Connect",
		Value:    strconv.FormatBool(connected),
		HttpOnly: true,
		Path:     "/",
	})

	email := request.FormValue("Insert_a_mail")
	Password := request.FormValue("Insert_a_password")
	alltable := selectAllFromTable(db, "users")

	hsha2psswd := sha256.Sum256([]byte(Password))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}

	var passwordtrue bool
	var emailtrue bool
	passwordtrue = false
	emailtrue = false
	Arr := VerifLogin(alltable)
	for i := 0; i < len(Arr); i++ {
		if Arr[i] == stringpsswd {
			passwordtrue = true
		}
		if Arr[i] == email {
			emailtrue = true
		}
	}
	fmt.Println(emailtrue, passwordtrue)
	if emailtrue == true && passwordtrue == true {
		println("LOGIN AUTORISEE")
		session, _ := store.Get(request, "Logged...")
		session.Values["Authentificated "] = true
		session.Values["Email "] = email
		session.Values["Password "] = Password
		fmt.Println(session.Values)
		session.Save(request, response)
		connected = true
		http.SetCookie(response, &http.Cookie{
			Name:     "Connect",
			Value:    strconv.FormatBool(connected),
			HttpOnly: true,
			Path:     "/",
		})
		tpl.ExecuteTemplate(response, "home.html", nil)
	} else {
		println("LOGIN REFUSER")
		tpl.ExecuteTemplate(response, "login.html", nil)
	}

}

func userprofile(response http.ResponseWriter, request *http.Request) {
	fmt.Println("InUser")
	Modif := request.FormValue("Input__info")
	INFO1 := request.FormValue("Tochange")

	ChangeEmail := false
	ChangeUsername := false

	session, _ := store.Get(request, "Logged...")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(response, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Println(INFO1)

	if INFO1 == "Email" {
		ChangeEmail = true
	} else if INFO1 == "Name" {
		ChangeUsername = true
	}

	Psswd := request.FormValue("Input__curt_pwd")
	fmt.Println("-------")
	fmt.Print("Password :  ")
	fmt.Println(Psswd)
	fmt.Print("New Value :  ")
	fmt.Println(Modif)
	fmt.Println("-------")

	hsha2psswd := sha256.Sum256([]byte(Psswd))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}

	passwordtrue := false

	alltable := selectAllFromTable(db, "users")
	Arr := NewUsername(alltable)
	for i := 0; i < len(Arr); i++ {
		if Arr[i] == stringpsswd {
			passwordtrue = true
		}
	}

	if passwordtrue == true && ChangeUsername == true {
		row := db.QueryRow("SELECT * FROM users WHERE password = ?;", stringpsswd)
		var p User
		err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			fmt.Println(err)
		}

		upStmt := "UPDATE `users` SET `name` = ? WHERE ( `password` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(Modif, stringpsswd)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		alltable1 := selectAllFromTable(db, "users")
		displayUsersRow(alltable1)
	} else if passwordtrue == true && ChangeEmail == true {

		row := db.QueryRow("SELECT * FROM users WHERE password = ?;", stringpsswd)
		var p User
		err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			fmt.Println(err)
		}

		upStmt := "UPDATE `users` SET `email` = ? WHERE ( `password` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(Modif, stringpsswd)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		alltable1 := selectAllFromTable(db, "users")
		displayUsersRow(alltable1)
	}

	tpl.ExecuteTemplate(response, "profile.html", nil) // Nil devient l'ensemble des posts (DB POST.)
}

func userpost(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/post.html", nil)
}

func usertheme(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/theme.html", nil)
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Like     int
	Post     int
}

type Post struct {
	Id      int
	Like    int
	Views   int
	Content string
	Name    string
	User_id int
}

func initDatabase(name string) *sql.DB {

	database, err := sql.Open("sqlite3", name)

	if err != nil {
		log.Fatal(err)
	}

	sqltbl := `
			PRAGMA foreign_keys = ON;
				CREATE TABLE IF NOT EXISTS users (
					 id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
					 name TEXT NOT NULL,
					 email TEXT NOT NULL UNIQUE,
					 password TEXT NOT NULL,
					 like INTEGER NOT NULL,
					 post INTEGER NOT NULL
				);
				
				CREATE TABLE IF NOT EXISTS posts (
					id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
					like INTEGER NOT NULL,
					views INTEGER NOT NULL,
					content TEXT NOT NULL,
					name TEXT NOT NULL,
					user_id INTEGER NOT NULL,
					FOREIGN KEY (user_id) REFERENCES users(id)
			   );
				`
	_, err = database.Exec(sqltbl)

	if err != nil {
		log.Fatal(err)
	}

	return database
}

func insertIntoUsers(db *sql.DB, name string, email string, password string, like int, post int) (int64, error) {

	hsha2psswd := sha256.Sum256([]byte(password))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}

	result, _ := db.Exec(`
				INSERT INTO users (name,email,password,like,post) VALUES (?,?,?,?,?)
						`,
		name, email, stringpsswd, like, post)

	return result.LastInsertId()
}

func insertIntoPosts(db *sql.DB, like int, views int, content string, name string, user_id int) (int64, error) {

	result2, _ := db.Exec(`
				INSERT INTO posts (like, views, content, name, user_id) VALUES (?,?,?,?,?)
						`,
		like, views, content, name, user_id)

	return result2.LastInsertId()
}

func selectAllFromTable(db *sql.DB, table string) *sql.Rows {

	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

func displayUsersRow(rows *sql.Rows) {
	for rows.Next() {
		var p User
		err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
}

func VerifRegister(rows *sql.Rows) []string {
	arr := []string{}
	for rows.Next() {
		var p User
		err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, p.Name)
		arr = append(arr, p.Email)
		arr = append(arr, p.Password)

	}
	return (arr)
}

func VerifLogin(rows *sql.Rows) []string {
	arr := []string{}
	for rows.Next() {
		var p User
		err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, p.Password)
		arr = append(arr, p.Email)
	}
	return (arr)
}

func NewUsername(rows *sql.Rows) []string {
	arr := []string{}
	for rows.Next() {
		var p User
		err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, p.Password)
		arr = append(arr, p.Name)
	}
	return (arr)
}

func displayPostsRow(rows *sql.Rows) {
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.User_id)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
}

func selectUserNameWithPattern(db *sql.DB, pattern string) *sql.Rows {
	query := "SELECT * FROM users WHERE name LIKE '%" + pattern + "%'"
	result, _ := db.Query(query)
	return result
}
