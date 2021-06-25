package sql

import (
	"database/sql"
	"net/http"
)

func Modifpost(response http.ResponseWriter, request *http.Request, Namestring string, Modif string) {
	if len(Namestring) > 0 && len(Modif) > 0 {
		URL := request.URL
		name, _ := URL.Query()["id"]
		upStmt := "UPDATE `posts` SET `name` = ? WHERE ( `id` = ?);"
		stmt, err := db.Prepare(upStmt)
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		var res sql.Result
		res, err = stmt.Exec(Namestring, name[0])
		rowsAff, _ := res.RowsAffected()
		if err != nil || rowsAff != 1 {
			panic(err)
		}

		// ====================================

		upStmt2 := "UPDATE `posts` SET `content` = ? WHERE ( `id` = ?);"
		stmt2, err := db.Prepare(upStmt2)
		if err != nil {
			panic(err)
		}
		defer stmt2.Close()
		var res2 sql.Result
		res2, err = stmt2.Exec(Modif, name[0])
		rowsAff2, _ := res2.RowsAffected()
		if err != nil || rowsAff2 != 1 {
			panic(err)
		}
	}
}
