package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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

func main() {

	// Creation de la DB.

	db := initDatabase("database.db")

	insertIntoUsers(db, "Mathieu", "m.m@gmail.com", "kenshilebouffon")

	insertIntoPosts(db, 2, 5, "CONTENT", "Post1", 1)

	alltable := selectAllFromTable(db, "users")
	alltable2 := selectAllFromTable(db, "posts")
	// alltablepost := selectAllFromTable(db, "posts")
	// usernamepattern := selectUserNameWithPattern(db, "as")
	fmt.Print("|---------------------------------------| \n")
	fmt.Print("  USER : ")
	displayUsersRow(alltable)
	fmt.Print("|---------------------------------------| \n")
	fmt.Print("  POST : ")
	displayPostsRow(alltable2)
	fmt.Print("|---------------------------------------| \n")
	// displayUsersRow(usernamepattern)
	// displayPostsRow(alltablepost)

	db.Close()
	//

}
