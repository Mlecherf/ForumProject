package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Tableau  []string `json:"tableau"`
	Time     string   `json:"time"`
}

func main() {
	tpl, _ = tpl.ParseGlob("*.html")
	http.HandleFunc("/", index)
	http.HandleFunc("/recup", recup)
	fmt.Println("Server UP")
	http.ListenAndServe(":8080", nil)
}

func index(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "index.html", nil)
}

func recup(response http.ResponseWriter, request *http.Request) {
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	fmt.Println(user)
	fmt.Println("user", user.Username)
	fmt.Println("pass", user.Password)
	fmt.Println("tab", user.Tableau)
	fmt.Println("time", user.Time)
	fmt.Println()
	// fmt.Printf(" user %s pass %s \n", user.Username, user.Password)
}
