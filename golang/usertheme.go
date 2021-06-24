package sql

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Usertheme(response http.ResponseWriter, request *http.Request) {
	URL := request.URL
	name, ok := URL.Query()["name"]
	if !ok || len(name[0]) < 1 {
		log.Println("Url Param 'name' is missing")
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
		ArrTagsBrut = append(ArrTagsBrut, p)
	}

	type Final struct {
		Post_info Post
		Tag_info  []string
	}
	type PopularPost struct {
		FinalArr  []Final
		NameTheme string
	}
	SelectedPosts := []Final{}
	NewArr2 := []string{}
	for i := 0; i < len(ArrTagsBrut); i++ {
		if strings.Contains(string(ArrTagsBrut[i].Tags), name[0]) {
			Final := Final{Post_info: ArrTagsBrut[i], Tag_info: strings.Split(ArrTagsBrut[i].Tags, "$")[1:]}
			SelectedPosts = append(SelectedPosts, Final)
		}
		NewArr2 = append(NewArr2, ArrTagsBrut[i].Tags)
		NewArr2 = append(NewArr2, "||")
	}
	row := db.QueryRow("SELECT `tags` FROM `posts`;")
	var p Post
	err1 := row.Scan(&p.Tags)
	if err1 != nil {
		fmt.Println(err1)
	}
	newstring := ""
	for i := 0; i < len(name[0]); i++ {

		if string(name[0][i]) == "_" {
			newstring += " "
		} else {
			newstring += string(name[0][i])
		}
	}
	SEND := PopularPost{FinalArr: SelectedPosts, NameTheme: newstring}
	Tpl.ExecuteTemplate(response, "view-theme.html", SEND)
}
