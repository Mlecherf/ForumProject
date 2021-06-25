package sql

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func Userprofile(response http.ResponseWriter, request *http.Request) {
	Modif := request.FormValue("Input__info")

	INFO1 := request.FormValue("Info")

	ChangeEmail := false
	ChangeUsername := false
	ChangePassword := false
	session, _ := store.Get(request, "Logged...")
	Verif := session.Values["Authentificated "]
	// Check if user is authenticated
	if Verif != true {
		http.Error(response, "Non connecté.", http.StatusForbidden)
		return
	}

	if INFO1 == "Email" {
		ChangeEmail = true
	} else if INFO1 == "Name" {
		ChangeUsername = true
	} else if INFO1 == "Password" {
		ChangePassword = true
	}

	Psswd := request.FormValue("Input__curt_pwd")
	hsha2psswd := sha256.Sum256([]byte(Psswd))
	stringpsswd := ""
	for i := 0; i < len(hsha2psswd); i++ {
		stringpsswd += string(hsha2psswd[i])
	}

	passwordtrue := false

	alltable := SelectAllFromTable(db, "users")
	Arr := NewUsername(alltable)
	for i := 0; i < len(Arr); i++ {
		if Arr[i] == stringpsswd {
			passwordtrue = true

		}
	}

	if passwordtrue == true && ChangeUsername == true {

		upStmt := "UPDATE `users` SET `name` = ? WHERE ( `password` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			fmt.Println(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(Modif, stringpsswd)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			fmt.Println(err)
		}
		sessionname, _ := store.Get(request, "Logged...")
		sessionname.Values["Authentificated "] = true
		sessionname.Values["Name "] = Modif
		sessionname.Save(request, response)

		http.Redirect(response, request, "/", http.StatusFound)
		return

	} else if passwordtrue == true && ChangeEmail == true {

		upStmt1 := "UPDATE `users` SET `email` = ? WHERE ( `password` = ?);"
		stmt1, err1 := db.Prepare(upStmt1)

		if err1 != nil {
			fmt.Println(err1)
		}

		defer stmt1.Close()
		var res1 sql.Result
		res1, err1 = stmt1.Exec(Modif, stringpsswd)

		rowsAff1, _ := res1.RowsAffected()
		if err1 != nil || rowsAff1 != 1 {
			fmt.Println(err1)
		}

		sessionemail, _ := store.Get(request, "Logged...")
		sessionemail.Values["Authentificated "] = true
		sessionemail.Values["Email "] = Modif
		sessionemail.Save(request, response)

		http.Redirect(response, request, "/", http.StatusFound)
		return
	} else if ChangePassword == true {

		upStmt := "UPDATE `users` SET `password` = ? WHERE ( `name` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			fmt.Println(err)
		}
		hsha2psswd := sha256.Sum256([]byte(Modif))
		stringpsswd2 := ""
		for i := 0; i < len(hsha2psswd); i++ {
			stringpsswd2 += string(hsha2psswd[i])
		}

		NameGo := session.Values["Name "]
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(stringpsswd2, NameGo)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			fmt.Println(err)
		}
		sessionpassword, _ := store.Get(request, "Logged...")
		sessionpassword.Values["Authentificated "] = true
		sessionpassword.Values["Password "] = Modif
		sessionpassword.Save(request, response)

		http.Redirect(response, request, "/", http.StatusFound)
		return
	}

	TablePost := SelectAllFromTable(db, "posts")
	IdPost := ReturnPostId(TablePost)

	rows := db.QueryRow("SELECT * FROM users WHERE name = ?;", session.Values["Name "])
	var s User
	err := rows.Scan(&s.Id, &s.Name, &s.Email, &s.Password, &s.Like, &s.Post)
	if err != nil {
		fmt.Println("err1", err)
	}

	type Final struct {
		Post_info Post
		Tag_info  []string
	}
	TOTALLIKE := 0
	TOTALPOST := 0
	TabPost := []Final{}
	for i := 0; i < len(IdPost); i++ {
		if IdPost[i] == s.Id {
			row := db.QueryRow("SELECT * FROM `posts` WHERE (`id` = ?);", i+1)
			var p Post
			err1 := row.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)
			if err1 != nil {
				fmt.Println("mypost", err1)
			}
			Allinfo := Final{Post_info: p, Tag_info: strings.Split(p.Tags, "$")[1:]}

			TOTALLIKE += p.Like
			TOTALPOST += 1
			TabPost = append(TabPost, Allinfo)
		}
	}

	type Data struct {
		Username     interface{}
		Email_Adress interface{}
		Display      []Final
		Like         int
		Postes       int
	}
	sessionfinal, _ := store.Get(request, "Logged...")
	Userss := sessionfinal.Values["Name "]
	Emailss := sessionfinal.Values["Email "]

	SEND := Data{Username: Userss, Email_Adress: Emailss, Display: TabPost, Like: TOTALLIKE, Postes: TOTALPOST}
	Tpl.ExecuteTemplate(response, "profile.html", SEND)
}
