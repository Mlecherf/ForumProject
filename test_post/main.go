package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

type User struct {
	Username string `json:"firstname"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/recup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(0)
		fmt.Println(r.FormValue("inp_username"))
		// var user User
		// json.NewDecoder(r.Body).Decode(&user)

		// fmt.Fprintf(w, "%s %s", user.Username, user.Password)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", nil)

	})

}

func index(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")
}
