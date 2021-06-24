package sql

import (
	"fmt"
	"net/http"
)

func Register(response http.ResponseWriter, request *http.Request) {
	email := request.FormValue("Email__input")
	username := request.FormValue("Username__input")
	password := request.FormValue("Password__input")
	passwordverif := request.FormValue("Password_Verif__input")

	like := 0
	post := 0

	if len(email) != 0 || len(username) != 0 || len(password) != 0 || len(passwordverif) != 0 {
		alltable := SelectAllFromTable(db, "users")
		Arr := VerifRegister(alltable)
		Wrong := false

		if len(Arr) != 0 {
			for i := 0; i < len(Arr); i++ {
				if Arr[i] == username || Arr[i] == email || Arr[i] == password {
					Wrong = true
				}
			}

			if Wrong == false {
				fmt.Println("Created... new enter in the DB.")
				InsertIntoUsers(db, username, email, password, like, post)
				alltable2 := SelectAllFromTable(db, "users")
				DisplayUsersRow(alltable2)
			} else {
				fmt.Println("Cant create... Already exist in the DB.")
			}

		} else {
			InsertIntoUsers(db, username, email, password, like, post)
			alltable1 := SelectAllFromTable(db, "users")
			DisplayUsersRow(alltable1)
		}
	}

	Tpl.ExecuteTemplate(response, "register.html", nil)
}
