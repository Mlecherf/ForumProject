package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("*.html")

	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/userprofile", userprofile)
	http.HandleFunc("/userpost", userpost)
	http.HandleFunc("/usertheme", usertheme)

	http.ListenAndServe(":8070", nil)
}

func home(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/home.html", nil)
}

func register(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("Email__input")
	username := request.FormValue("Username__input")
	password := request.FormValue("Password__input")
	passwordverif := request.FormValue("Password_Verif__input")
	println("=================")
	println("EMAIL :", email)
	println("USERNAME :", username)
	println("PASSWORD :", password)
	println("PASSWORDVERIF :", passwordverif)
	println("=================")
	tpl.ExecuteTemplate(response, "register.html", nil)

	db := initDatabase("database" + username + ".db")

	insertIntoUsers(db, username, email, password)

	alltable := selectAllFromTable(db, "users")
	displayUsersRow(alltable)
}

func login(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "template/login.html", nil)
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
					 password TEXT NOT NULL
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

func insertIntoUsers(db *sql.DB, name string, email string, password string) (int64, error) {

	hsha2psswd := sha256.Sum256([]byte(password))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}

	result, _ := db.Exec(`
				INSERT INTO users (name,email,password) VALUES (?,?,?)
						`,
		name, email, stringpsswd)

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
		err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Password)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
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
