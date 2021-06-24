package sql

import "net/http"

func Logout(response http.ResponseWriter, request *http.Request) {
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
	Tpl.ExecuteTemplate(response, "home.html", nil)
}
