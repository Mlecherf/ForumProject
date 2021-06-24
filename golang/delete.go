package sql

import (
	"crypto/sha256"
	"fmt"
	"net/http"
)

func Delete(response http.ResponseWriter, request *http.Request) {
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
	Tpl.ExecuteTemplate(response, "home.html", nil)
}
