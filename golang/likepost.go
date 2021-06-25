package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Likepost struct {
	Id int
}

func LikePost(response http.ResponseWriter, request *http.Request) {
	var X Likepost
	json.NewDecoder(request.Body).Decode(&X)

	session, _ := store.Get(request, "Logged...")

	row := db.QueryRow("SELECT * FROM posts WHERE id = ?;", X.Id)
	var p Post
	err := row.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)
	if err != nil {
		fmt.Println(err)
	}

	Verification := strings.Contains(p.LikeList, session.Values["Name "].(string))

	if !Verification {

		upStmt := "UPDATE `posts` SET `likelist` = ? WHERE ( `id` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(p.LikeList+session.Values["Name "].(string), X.Id)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		upStmt2 := "UPDATE `posts` SET `like` = ? WHERE ( `id` = ?);"
		stmt2, err2 := db.Prepare(upStmt2)

		if err2 != nil {
			panic(err2)
		}

		defer stmt2.Close()
		var res2 sql.Result
		res2, err2 = stmt2.Exec(p.Like+1, X.Id)
		rowsAff2, _ := res2.RowsAffected()
		if err2 != nil || rowsAff2 != 1 {
			panic(err2)
		}

	} else {
		upStmt := "UPDATE `posts` SET `likelist` = ? WHERE ( `id` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		Str := strings.ReplaceAll(p.LikeList, session.Values["Name "].(string), "")

		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(Str, X.Id)
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		upStmt2 := "UPDATE `posts` SET `like` = ? WHERE ( `id` = ?);"
		stmt2, err2 := db.Prepare(upStmt2)

		if err2 != nil {
			panic(err2)
		}

		defer stmt2.Close()
		var res2 sql.Result
		res2, err2 = stmt2.Exec(p.Like-1, X.Id)
		rowsAff2, _ := res2.RowsAffected()
		if err2 != nil || rowsAff2 != 1 {
			panic(err2)
		}
	}

}
