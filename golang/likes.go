package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Likes(response http.ResponseWriter, request *http.Request) {

	var Like Like
	Like.Like = ""
	session, _ := store.Get(request, "Logged...")
	NameUser := session.Values["Name "]
	json.NewDecoder(request.Body).Decode(&Like)
	URL := Like.Url
	ID := ""
	for i := 0; i < len(URL); i++ {
		if i+1 != len(URL) {
			if string(URL[i]) == "i" && string(URL[i+1]) == "d" {
				ID = string(URL[i+3])
			}
		}
	}

	row := db.QueryRow("SELECT * FROM posts WHERE id = ?;", ID)
	var p Post
	err3 := row.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

	if err3 != nil {
		fmt.Println(err3)
	}

	if !(strings.Contains(p.LikeList, NameUser.(string))) {

		upStmt := "UPDATE `posts` SET `like` = ? WHERE ( `id` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(p.Like+1, ID)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		upStmt2 := "UPDATE `posts` SET `likelist` = ? WHERE ( `id` = ?);"
		stmt2, err2 := db.Prepare(upStmt2)

		if err2 != nil {
			panic(err2)
		}

		defer stmt2.Close()
		var res2 sql.Result
		res2, err2 = stmt2.Exec(NameUser, ID)
		rowsAff2, _ := res2.RowsAffected()
		if err2 != nil || rowsAff2 != 1 {
			panic(err2)
		}
	} else if strings.Contains(p.LikeList, NameUser.(string)) {
		upStmt := "UPDATE `posts` SET `like` = ? WHERE ( `id` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(p.Like-1, ID)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		upStmt2 := "UPDATE `posts` SET `likelist` = ? WHERE ( `id` = ?);"
		stmt2, err2 := db.Prepare(upStmt2)

		if err2 != nil {
			panic(err2)
		}

		Ress := strings.ReplaceAll(p.LikeList, NameUser.(string), "")
		defer stmt2.Close()
		var res2 sql.Result
		res2, err2 = stmt2.Exec(Ress, ID)
		rowsAff2, _ := res2.RowsAffected()
		if err2 != nil || rowsAff2 != 1 {
			panic(err2)
		}
	}

	NewString := ""
	for i := 0; i < len(URL); i++ {
		if i+1 != len(URL) {
			if string(URL[i]) == "/" && string(URL[i+1]) == "v" {
				for z := i; z < len(URL); z++ {
					NewString += string(URL[z])
				}

			}
		}
	}

	http.Redirect(response, request, NewString, http.StatusSeeOther)
}
