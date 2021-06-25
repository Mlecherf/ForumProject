package main

import (
	sql "angrycreative/golang"
	"fmt"
	"net/http"
)

func main() {
	var err error
	sql.Tpl, err = sql.Tpl.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", sql.Home)
	http.HandleFunc("/register", sql.Register)
	http.HandleFunc("/login", sql.Login)
	http.HandleFunc("/profile", sql.Userprofile)
	http.HandleFunc("/viewpost", sql.Userpost)
	http.HandleFunc("/theme", sql.Usertheme)
	http.HandleFunc("/logout", sql.Logout)
	http.HandleFunc("/delete", sql.Delete)
	http.HandleFunc("/recup", sql.Recup)
	http.HandleFunc("/like", sql.Likes)
	http.HandleFunc("/deletepost", sql.DeletePostModif)
	style := http.FileServer(http.Dir("asset/style/"))
	image := http.FileServer(http.Dir("asset/image/"))
	js := http.FileServer(http.Dir("asset/js/"))

	http.Handle("/static/style/", http.StripPrefix("/static/style/", style))
	http.Handle("/static/image/", http.StripPrefix("/static/image/", image))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", js))

	http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
}
