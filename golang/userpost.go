package sql

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Userpost(response http.ResponseWriter, request *http.Request) {

	NameModif := request.FormValue("Name_mod")
	ContentModif := request.FormValue("Content_mod")
	Modifpost(response, request, NameModif, ContentModif)

	URL := request.URL
	name, ok := URL.Query()["id"]
	if !ok || len(name[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}

	ALLTABLE := SelectAllFromTable(db, "posts")
	ArrTagsBrut := []Post{}
	for ALLTABLE.Next() {
		var p Post
		err := ALLTABLE.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

		if err != nil {
			log.Fatal(err)
		}
		if name[0] == strconv.Itoa(p.Id) {
			ArrTagsBrut = append(ArrTagsBrut, p)
		}
	}

	type Final struct {
		Post_info Post
		Tag_info  []string
		Owner     bool
	}

	ALLTABLE2 := SelectAllFromTable(db, "users")
	UserTab := []User{}
	for ALLTABLE2.Next() {
		var p User
		err := ALLTABLE2.Scan(&p.Id, &p.Name, &p.Email, &p.Password, &p.Like, &p.Post)

		if err != nil {
			log.Fatal(err)
		}
		UserTab = append(UserTab, p)
	}
	Connected_ID := 0
	session, _ := store.Get(request, "Logged...")
	for i := 0; i < len(UserTab); i++ {
		if UserTab[i].Name == session.Values["Name "] {
			Connected_ID = UserTab[i].Id
		}
	}

	Owner := false
	for i := 0; i < len(ArrTagsBrut); i++ {

		row := db.QueryRow("SELECT * FROM posts WHERE id = ?;", ArrTagsBrut[i].Id)
		var p Post
		err := row.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)
		if !(strings.Contains(string(ArrTagsBrut[i].ViewList), strconv.Itoa(Connected_ID))) {

			if err != nil {
				fmt.Println(err)
			}

			upStmt := "UPDATE `posts` SET `views` = ? WHERE ( `id` = ?);"
			stmt, err := db.Prepare(upStmt)

			if err != nil {
				panic(err)
			}

			defer stmt.Close()
			var res sql.Result
			res, err = stmt.Exec(p.Views+1, name[0])
			rowsAff, _ := res.RowsAffected()
			if err != nil || rowsAff != 1 {
				panic(err)
			}

			upStmt2 := "UPDATE `posts` SET `viewlist` = ? WHERE ( `id` = ?);"
			stmt2, err2 := db.Prepare(upStmt2)

			if err2 != nil {
				panic(err2)
			}
			defer stmt2.Close()
			var res2 sql.Result
			res2, err2 = stmt2.Exec(strconv.Itoa(Connected_ID), p.Id)
			rowsAff2, _ := res2.RowsAffected()
			if err2 != nil || rowsAff2 != 1 {
				panic(err2)
			}
		}

		if strconv.Itoa(ArrTagsBrut[i].Id) == name[0] {
			if Connected_ID == ArrTagsBrut[i].User_id {
				Owner = true
			}
			Finals := Final{Post_info: ArrTagsBrut[i], Tag_info: strings.Split(ArrTagsBrut[i].Tags, "$")[1:], Owner: Owner}
			Tpl.ExecuteTemplate(response, "Post-view.html", Finals)
		}

	}

}
