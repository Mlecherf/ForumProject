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

	_ "github.com/mattn/go-sqlite3"
)

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

func main() {
	tpl, _ = tpl.ParseGlob("template/*.html")

	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/userprofile", userprofile)
	http.HandleFunc("/userpost", userpost)
	http.HandleFunc("/usertheme", usertheme)

	style := http.FileServer(http.Dir("asset/style/"))
	image := http.FileServer(http.Dir("asset/image/"))
	js := http.FileServer(http.Dir("asset/js/"))

	http.Handle("/static/style/", http.StripPrefix("/static/style/", style))
	http.Handle("/static/image/", http.StripPrefix("/static/image/", image))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", js))

	http.ListenAndServe(":8070", nil)
}

func home(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/home.html", nil)
}

var db = initDatabase("database.db")

func register(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("Email__input")
	username := request.FormValue("Username__input")
	password := request.FormValue("Password__input")
	passwordverif := request.FormValue("Password_Verif__input")
	like := 0
	post := 0
	println("=================")
	println("EMAIL :", email)
	println("USERNAME :", username)
	println("PASSWORD :", password)
	println("PASSWORDVERIF :", passwordverif)
	println("=================")

	D := "{ 'user': " + username + " 'mail': " + email + " 'nb posts': " + strconv.Itoa(post) + " 'nb likes': " + strconv.Itoa(like) + " }"

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "Login", Value: D, Expires: expiration}
	http.SetCookie(response, &cookie)

	insertIntoUsers(db, username, email, password, like, post)

	alltable := selectAllFromTable(db, "users")
	displayUsersRow(alltable)

	tpl.ExecuteTemplate(response, "register.html", nil)
}

func login(response http.ResponseWriter, request *http.Request) {

	email := request.FormValue("Insert_a_Pseudo")
	Password := request.FormValue("Password__input")
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

	if emailtrue == true && passwordtrue == true {
		println("LOGIN AUTORISEE")
		tpl.ExecuteTemplate(response, "index.html", nil)
	} else {
		println("LOGIN REFUSER")
		tpl.ExecuteTemplate(response, "login.html", nil)
	}

}

func userprofile(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/userprofile.html", nil)
}

func userpost(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/userpost.html", nil)
}

func usertheme(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/usertheme.html", nil)
}

// func managepost(response http.ResponseWriter, request *http.Request) {
// 	http.ServeFile(response, request, "template/managepost.html")
// }

// func modify(response http.ResponseWriter, request *http.Request) {
// 	http.ServeFile(response, request, "template/modify.html")
// }

// func manageprofile(response http.ResponseWriter, request *http.Request) {
// 	http.ServeFile(response, request, "template/manageprofile.html")
// }

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

// func database() {

// 	// Creation de la DB.

// 	db := initDatabase("database.db")

// 	insertIntoUsers(db, Name, mail, m)

// 	insertIntoPosts(db, 2, 5, "CONTENT", "Post1", 1)

// 	alltable := selectAllFromTable(db, "users")
// 	alltable2 := selectAllFromTable(db, "posts")

// 	fmt.Print("|---------------------------------------| \n")
// 	fmt.Print("  USER : ")
// 	displayUsersRow(alltable)
// 	fmt.Print("|---------------------------------------| \n")
// 	fmt.Print("  POST : ")
// 	displayPostsRow(alltable2)
// 	fmt.Print("|---------------------------------------| \n")

// 	db.Close()

// }
