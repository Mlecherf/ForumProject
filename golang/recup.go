package sql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Recup(response http.ResponseWriter, request *http.Request) {
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

	InsertIntoPosts(db, post.Like, post.Views, post.Content, post.Name, post.Tags, post.User_id, "", "")

	type Final struct {
		Post_info Post
	}

	fmt.Println("New enter in database / posts...")
	Tpl.ExecuteTemplate(response, "home.html", nil)
}
