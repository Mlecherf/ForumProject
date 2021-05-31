package main

import (
	"net/http"
)

func serveur() {
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/userprofile", userprofile)
	http.HandleFunc("/userpost", userpost)
	http.HandleFunc("/usertheme", usertheme)
	// http.HandleFunc("/modify", modify)
	// http.HandleFunc("/manageprofile", manageprofile)
	// http.HandleFunc("/managepost", managepost)
	http.ListenAndServe(":80", nil)
}

func home(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "template/home.html")
}

func register(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "template/register.html")
}

func login(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "template/login.html")
}

func userprofile(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "template/userprofile.html")
}

func userpost(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "template/userpost.html")
}

func usertheme(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "template/usertheme.html")
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
