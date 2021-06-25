package sql

import (
	"net/http"
)

func DeletePostModif(response http.ResponseWriter, request *http.Request) {

	idvalue := (request.FormValue("DeletePost"))
	if idvalue != "" {
		upStmt := "DELETE FROM posts WHERE ( `id` = ?);"
		stmt, err := db.Prepare(upStmt)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()

		_, err = stmt.Exec(idvalue)
		http.Redirect(response, request, "/", http.StatusSeeOther)
	}
}
