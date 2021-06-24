package sql

import (
	"crypto/sha256"
	"log"
	"net/http"
	"strconv"
)

func Login(response http.ResponseWriter, request *http.Request) {
	connected := false
	http.SetCookie(response, &http.Cookie{
		Name:     "Connect",
		Value:    strconv.FormatBool(connected),
		HttpOnly: false,
		Path:     "/",
		MaxAge:   150,
	})

	email := request.FormValue("Insert_a_mail")
	Password := request.FormValue("Insert_a_password")
	alltable := SelectAllFromTable(db, "users")

	hsha2psswd := sha256.Sum256([]byte(Password))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}

	var passwordtrue bool
	var emailtrue bool
	passwordtrue = false
	emailtrue = false
	Arr := VerifLogin(alltable)
	for i := 0; i < len(Arr); i++ {

		if Arr[i] == stringpsswd {
			passwordtrue = true
		}
		if Arr[i] == email {
			emailtrue = true
		}

	}

	if emailtrue == true && passwordtrue == true {

		ALLTABLE := SelectAllFromTable(db, "users")
		ArrTagsBrut := []User{}
		for ALLTABLE.Next() {
			var p User
			err := ALLTABLE.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

			if err != nil {
				log.Fatal(err)
			}
			ArrTagsBrut = append(ArrTagsBrut, p)
		}

		Name := ""

		for i := 0; i < len(ArrTagsBrut); i++ {
			if ArrTagsBrut[i].Email == email && ArrTagsBrut[i].Password == stringpsswd {
				Name = ArrTagsBrut[i].Name
			}
		}

		session, _ := store.Get(request, "Logged...")
		session.Values["Authentificated "] = true
		session.Values["Email "] = email
		session.Values["Password "] = Password
		session.Values["Name "] = Name

		session.Save(request, response)
		connected = true
		http.SetCookie(response, &http.Cookie{
			Name:     "Connect",
			Value:    strconv.FormatBool(connected),
			HttpOnly: false,
			Path:     "/",
			MaxAge:   999999,
		})
		http.Redirect(response, request, "/", http.StatusSeeOther)

	} else {
		Tpl.ExecuteTemplate(response, "login.html", nil)
	}

}
