package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var db = initDatabase("database.db")

var tpl *template.Template

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	var err error
	tpl, err = tpl.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/profile", userprofile)
	http.HandleFunc("/post", userpost)
	http.HandleFunc("/theme", usertheme)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/recup", recup)
	style := http.FileServer(http.Dir("asset/style/"))
	image := http.FileServer(http.Dir("asset/image/"))
	js := http.FileServer(http.Dir("asset/js/"))

	http.Handle("/static/style/", http.StripPrefix("/static/style/", style))
	http.Handle("/static/image/", http.StripPrefix("/static/image/", image))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", js))

	http.ListenAndServe(":8020", nil)
}

func home(response http.ResponseWriter, request *http.Request) {

	TabPost := []Post{}
	alltable := selectAllFromTable(db, "posts")
	Tabint := returnPostLike(alltable)
	if len(Tabint) != 0 {
		sort.Ints(Tabint)
		ToSearch := 0
		for i := len(Tabint) - 1; i >= 0; i-- {
			if i > len(Tabint)-4 {
				ToSearch = Tabint[i]
				row1 := db.QueryRow("SELECT * FROM posts WHERE like = ?;", ToSearch)
				var s Post
				err := row1.Scan(&s.Id, &s.Like, &s.Views, &s.Content, &s.Name, &s.Tags, &s.User_id)
				if err != nil {
					fmt.Println(err)
				}
				TabPost = append(TabPost, s)
			}
		}
	}
	tpl.ExecuteTemplate(response, "home.html", TabPost)
}

type Test struct {
	Tableau []string
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
	Tags    string
	User_id int
}

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

func recup(response http.ResponseWriter, request *http.Request) {
	var post Post
	post.Like = 0
	post.Views = 0
	session, _ := store.Get(request, "Logged...")
	NameUser := session.Values["Name "]
	json.NewDecoder(request.Body).Decode(&post)
	row := db.QueryRow("SELECT * FROM users WHERE name = ?;", NameUser)
	var p User
	err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)
	if err != nil {
		fmt.Println(err)
	}
	post.User_id = p.Id

	insertIntoPosts(db, post.Like, post.Views, post.Content, post.Name, post.Tags, post.User_id)

	type Final struct {
		Post_info Post
	}

	X := Final{post}
	fmt.Println("New enter in database / posts...")
	fmt.Println(X)
	tpl.ExecuteTemplate(response, "home.html", nil)
}

func register(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("Email__input")
	username := request.FormValue("Username__input")
	password := request.FormValue("Password__input")
	passwordverif := request.FormValue("Password_Verif__input")

	like := 0
	post := 0

	if len(email) != 0 || len(username) != 0 || len(password) != 0 || len(passwordverif) != 0 {
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
		HttpOnly: false,
		Path:     "/",
		MaxAge:   150,
	})

	email := request.FormValue("Insert_a_mail")
	Password := request.FormValue("Insert_a_password")
	alltable := selectAllFromTable(db, "users")

	hsha2psswd := sha256.Sum256([]byte(Password))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}
	Name := ""
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

		row := db.QueryRow("SELECT * FROM users WHERE password = ?;", stringpsswd)
		var p User
		err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)
		Name = p.Name
		if err != nil {
			fmt.Println(err)
		}
	}

	if emailtrue == true && passwordtrue == true {
		fmt.Println("Login...")
		session, _ := store.Get(request, "Logged...")
		session.Values["Authentificated "] = true
		session.Values["Email "] = email
		session.Values["Password "] = Password
		session.Values["Name "] = Name

		session.Save(request, response)
		connected = true
		http.SetCookie(response, &http.Cookie{
			Name:     "Connect",
			Value:    strconv.FormatBool(connected),
			HttpOnly: false,
			Path:     "/",
			MaxAge:   999999,
		})
		tpl.ExecuteTemplate(response, "home.html", nil)
	} else {
		fmt.Println("Can't login...")
		tpl.ExecuteTemplate(response, "login.html", nil)
	}

}

func logout(response http.ResponseWriter, request *http.Request) {
	http.SetCookie(response, &http.Cookie{
		Name:     "Connect",
		HttpOnly: false,
		Path:     "/",
		MaxAge:   -1,
	})
	http.SetCookie(response, &http.Cookie{
		Name:     "Logged...",
		HttpOnly: false,
		Path:     "/",
		MaxAge:   -1,
	})
	tpl.ExecuteTemplate(response, "home.html", nil)
}

func delete(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "Logged...")
	Psswd := session.Values["Password "]
	mycode := Psswd.(string)
	hsha2psswd := sha256.Sum256([]byte(mycode))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}
	upStmt := "DELETE FROM users WHERE ( `password` = ?);"
	stmt, err := db.Prepare(upStmt)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(stringpsswd)

	fmt.Println("User Deleted...")

	http.SetCookie(response, &http.Cookie{
		Name:     "Connect",
		HttpOnly: false,
		Path:     "/",
		MaxAge:   -1,
	})
	http.SetCookie(response, &http.Cookie{
		Name:     "Logged...",
		HttpOnly: false,
		Path:     "/",
		MaxAge:   -1,
	})
	tpl.ExecuteTemplate(response, "home.html", nil)
}

func userprofile(response http.ResponseWriter, request *http.Request) {
	Modif := request.FormValue("Input__info")

	INFO1 := request.FormValue("Info")

	ChangeEmail := false
	ChangeUsername := false
	ChangePassword := false
	session, _ := store.Get(request, "Logged...")

	Verif := session.Values["Authentificated "]
	// Check if user is authenticated
	if Verif != true {
		http.Error(response, "Non connecté.", http.StatusForbidden)
		return
	}

	if INFO1 == "Email" {
		ChangeEmail = true
	} else if INFO1 == "Name" {
		ChangeUsername = true
	} else if INFO1 == "Password" {
		ChangePassword = true
	}

	Psswd := request.FormValue("Input__curt_pwd")
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
		session.Values["Name "] = Modif
		session.Save(request, response)

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

		session.Values["Email "] = Modif
		session.Save(request, response)

	} else if passwordtrue == true && ChangePassword == true {
		row := db.QueryRow("SELECT * FROM users WHERE password = ?;", stringpsswd)
		var p User
		err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			fmt.Println(err)
		}

		upStmt := "UPDATE `users` SET `password` = ? WHERE ( `password` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}
		hsha2psswd := sha256.Sum256([]byte(Modif))
		stringpsswd2 := ""
		for i := 0; i < len(hsha2psswd); i++ {
			stringpsswd2 += string(hsha2psswd[i])
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(stringpsswd2, stringpsswd)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}
		session.Values["Password "] = Modif
		session.Save(request, response)

	}
	TablePost := selectAllFromTable(db, "posts")
	IdPost := returnPostId(TablePost)

	rows := db.QueryRow("SELECT * FROM users WHERE name = ?;", session.Values["Name "])
	var s User
	err := rows.Scan(&s.Id, &s.Name, &s.Email, &s.Password, &s.Like, &s.Post)
	if err != nil {
		fmt.Println("err1", err)
	}

	type Final struct {
		Post_info Post
		Tag_info  []string
	}

	fmt.Println(IdPost)
	TabPost := []Final{}
	for i := 0; i < len(IdPost); i++ {
		if IdPost[i] == s.Id {
			row := db.QueryRow("SELECT * FROM `posts` WHERE (`id` = ?);", i+1)
			var p Post
			err1 := row.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id)
			if err1 != nil {
				fmt.Println("mypost", err1)
			}
			fmt.Println("My tag", p)
			fmt.Println("Tag info", p.Tags, strings.Split(p.Tags, "$")[1:])

			Allinfo := Final{Post_info: p, Tag_info: strings.Split(p.Tags, "$")[1:]}
			TabPost = append(TabPost, Allinfo)
		}
	}

	type Data struct {
		Username     interface{}
		Email_Adress interface{}
		Display      []Final
	}
	Userss := session.Values["Name "]
	Emailss := session.Values["Email "]

	SEND := Data{Username: Userss, Email_Adress: Emailss, Display: TabPost}
	tpl.ExecuteTemplate(response, "profile.html", SEND)
}

func userpost(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/post.html", nil)
}

func usertheme(response http.ResponseWriter, request *http.Request) {
	// recup url
	// ? -> fastfood
	// -> dbb recup tout les tags fastfood
	URL := request.URL
	name, ok := URL.Query()["name"]
	if !ok || len(name[0]) < 1 {
		log.Println("Url Param 'name' is missing")
		return
	}
	ALLTABLE := selectAllFromTable(db, "posts")
	ArrTagsBrut := []Post{}
	for ALLTABLE.Next() {
		var p Post
		err := ALLTABLE.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id)

		if err != nil {
			log.Fatal(err)
		}
		ArrTagsBrut = append(ArrTagsBrut, p)
	}
	SelectedPosts := []Post{}
	NewArr2 := []string{}
	for i := 0; i < len(ArrTagsBrut); i++ {
		if strings.Contains(string(ArrTagsBrut[i].Tags), name[0]) {
			SelectedPosts = append(SelectedPosts, ArrTagsBrut[i])
		}
		NewArr2 = append(NewArr2, ArrTagsBrut[i].Tags)
		NewArr2 = append(NewArr2, "||")
	}
	row := db.QueryRow("SELECT `tags` FROM `posts`;")
	var p Post
	err1 := row.Scan(&p.Tags)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(SelectedPosts)
	tpl.ExecuteTemplate(response, "view-theme.html", SelectedPosts)
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
					tags TEXT NOT NULL,
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

func insertIntoPosts(db *sql.DB, like int, views int, content string, name string, tags string, user_id int) (int64, error) {

	result, _ := db.Exec(`
				INSERT INTO posts (like, views, content, name, tags, user_id) VALUES (?,?,?,?,?,?)
						`,
		like, views, content, name, tags, user_id)

	return result.LastInsertId()
}

func updatePosts(db *sql.DB, id int, content string, name string, user_id string) {
	sage := strconv.Itoa(id)
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update testTable set id=?,content=?,name=? user_id=?")
	_, err := stmt.Exec(sage, content, name, user_id)
	checkError(err)
	tx.Commit()
}

func deletePost(db *sql.DB, id int) {
	sid := strconv.Itoa(id) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from testTable where id=?")
	_, err := stmt.Exec(sid)
	checkError(err)
	tx.Commit()
}

func checkError(err error) {
	if err != nil {
	}
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
		arr = append(arr, p.Name)
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

	}
	return (arr)
}

func displayPostsRow(rows *sql.Rows) {
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
}

func returnPostLike(rows *sql.Rows) []int {
	arr := []int{}
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id)

		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, p.Like)
	}
	return arr
}

func returnPostId(rows *sql.Rows) []int {
	arr := []int{}
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id)

		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, p.User_id)
	}
	return arr
}

func selectUserNameWithPattern(db *sql.DB, pattern string) *sql.Rows {
	query := "SELECT * FROM users WHERE name LIKE '%" + pattern + "%'"
	result, _ := db.Query(query)
	return result
}
