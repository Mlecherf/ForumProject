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
		row := db.QueryRow("SELECT * FROM users WHERE password = ?;", stringpsswd)
		var p User
		err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			fmt.Println(err)
		}

		upStmt := "UPDATE `users` SET `name` = ? WHERE ( `password` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(Modif, stringpsswd)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}
		session.Values["Name "] = Modif
		session.Save(request, response)

	} else if passwordtrue == true && ChangeEmail == true {

		row := db.QueryRow("SELECT * FROM users WHERE password = ?;", stringpsswd)
		var p User
		err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			fmt.Println(err)
		}

		upStmt := "UPDATE `users` SET `email` = ? WHERE ( `password` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(Modif, stringpsswd)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		session.Values["Email "] = Modif
		session.Save(request, response)

	} else if passwordtrue == true && ChangePassword == true {
		row := db.QueryRow("SELECT * FROM users WHERE password = ?;", stringpsswd)
		var p User
		err := row.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			fmt.Println(err)
		}

		upStmt := "UPDATE `users` SET `password` = ? WHERE ( `password` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}
		hsha2psswd := sha256.Sum256([]byte(Modif))
		stringpsswd2 := ""
		for i := 0; i < len(hsha2psswd); i++ {
			stringpsswd2 += string(hsha2psswd[i])
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(stringpsswd2, stringpsswd)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}
		session.Values["Password "] = Modif
		session.Save(request, response)

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
	Userss := session.Values["Name "]
	Emailss := session.Values["Email "]

	SEND := Data{Username: Userss, Email_Adress: Emailss, Display: TabPost, Like: TOTALLIKE, Postes: TOTALPOST}
	fmt.Println(SEND)
	Tpl.ExecuteTemplate(response, "profile.html", SEND)
}
