package sql

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db = InitDatabase("database.db")
var Tpl *template.Template

func InitDatabase(name string) *sql.DB {

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
					viewlist TEXT NOT NULL,
					likelist TEXT NOT NULL,
					FOREIGN KEY (user_id) REFERENCES users(id)
			   );
				`
	_, err = database.Exec(sqltbl)

	if err != nil {
		log.Fatal(err)
	}

	return database
}

func InsertIntoUsers(db *sql.DB, name string, email string, password string, like int, post int) (int64, error) {

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

func InsertIntoPosts(db *sql.DB, like int, views int, content string, name string, tags string, user_id int, viewlist string, likelist string) (int64, error) {

	result, _ := db.Exec(`
				INSERT INTO posts (like, views, content, name, tags, user_id, viewlist, likelist) VALUES (?,?,?,?,?,?,?,?)
						`,
		like, views, content, name, tags, user_id, viewlist, likelist)

	return result.LastInsertId()
}

func UpdatePosts(db *sql.DB, id int, content string, name string, user_id string) {
	sage := strconv.Itoa(id)
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update testTable set id=?,content=?,name=? user_id=?")
	_, err := stmt.Exec(sage, content, name, user_id)
	CheckError(err)
	tx.Commit()
}

func DeletePost(db *sql.DB, id int) {
	sid := strconv.Itoa(id) // int to string
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from testTable where id=?")
	_, err := stmt.Exec(sid)
	CheckError(err)
	tx.Commit()
}

func CheckError(err error) {
	if err != nil {
	}
}

func SelectAllFromTable(db *sql.DB, table string) *sql.Rows {

	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

func DisplayUsersRow(rows *sql.Rows) {
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
		arr = append(arr, p.Password)
	}
	return (arr)
}

func DisplayPostsRow(rows *sql.Rows) {
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
}

func ReturnPostLike(rows *sql.Rows) []int {
	arr := []int{}
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, p.Like)
	}
	return arr
}

func ReturnPostId(rows *sql.Rows) []int {
	arr := []int{}
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, p.User_id)
	}
	return arr
}

func SelectUserNameWithPattern(db *sql.DB, pattern string) *sql.Rows {
	query := "SELECT * FROM users WHERE name LIKE '%" + pattern + "%'"
	result, _ := db.Query(query)
	return result
}
