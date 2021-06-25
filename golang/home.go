package sql

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func Home(response http.ResponseWriter, request *http.Request) {

	ALLTABLE := SelectAllFromTable(db, "posts")
	ArrTagsBrut := []Post{}
	ThemePost := ""
	for ALLTABLE.Next() {
		var p Post
		err := ALLTABLE.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

		if err != nil {
			fmt.Println("->", err)
		}
		ThemePost += p.Tags
		ArrTagsBrut = append(ArrTagsBrut, p)
	}

	IntArr := []int{}
	for i := len(ArrTagsBrut) - 1; i >= 0; i-- {
		IntArr = append(IntArr, ArrTagsBrut[i].Like)
	}
	sort.Ints(IntArr)
	type Final struct {
		Post_info Post
		Tag_info  []string
	}

	ALLTABLE2 := SelectAllFromTable(db, "posts")
	FinalPost := []Post{}
	verif := false
	Likes := []int{}
	if len(IntArr) > 4 {
		for i := len(IntArr) - 1; i >= len(IntArr)-4; i-- {
			Likes = append(Likes, IntArr[i])
		}
		for ALLTABLE2.Next() {
			var p Post
			err := ALLTABLE2.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

			if err != nil {
				fmt.Println("->", err)
			}
			for i := 0; i < len(Likes); i++ {

				if len(FinalPost) < 4 {
					if p.Like == Likes[i] {
						for z := 0; z < len(FinalPost); z++ {
							if FinalPost[z] == p {
								verif = true
							}
						}
						if verif == true {
							verif = false
						} else {
							FinalPost = append(FinalPost, p)
						}

					}
				}
			}
			if len(IntArr) < 4 {

				FinalPost = append(FinalPost, p)

			}
		}
	} else {
		for ALLTABLE2.Next() {
			var p Post
			err := ALLTABLE2.Scan(&p.Id, &p.Like, &p.Views, &p.Content, &p.Name, &p.Tags, &p.User_id, &p.ViewList, &p.LikeList)

			if err != nil {
				fmt.Println("->", err)
			}
			for i := 0; i < len(IntArr); i++ {
				if p.Like == IntArr[i] {
					for z := 0; z < len(FinalPost); z++ {
						if FinalPost[z] == p {
							verif = true
						}
					}
					if verif == true {
						verif = false
					} else {
						FinalPost = append(FinalPost, p)
					}

				}
			}
			FinalPost = append(FinalPost)
		}
	}

	FinalArr := []Final{}

	for i := 0; i < len(FinalPost); i++ {
		Posts := Final{Post_info: FinalPost[i], Tag_info: strings.Split(FinalPost[i].Tags, "$")[1:]}
		FinalArr = append(FinalArr, Posts)
	}

	ALLTABLE3 := SelectAllFromTable(db, "posts")
	ArrStr := []string{}
	for ALLTABLE3.Next() {
		var x Post
		err := ALLTABLE3.Scan(&x.Id, &x.Like, &x.Views, &x.Content, &x.Name, &x.Tags, &x.User_id, &x.ViewList, &x.LikeList)

		if err != nil {
			fmt.Println("->", err)
		}
		Z := strings.Split(x.Tags, "$")
		for i := 0; i < len(Z); i++ {
			ArrStr = append(ArrStr, Z[i])
		}
	}
	FFC, PC, BC, DC, AC, IC, MC, INC, JC, FC, AFC, BBQC, KC, VC, ALC, BAC := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	for i := 0; i < len(ArrStr); i++ {
		if ArrStr[i] == "Fast_Food" {
			FFC += 1
		} else if ArrStr[i] == "Pizza" {
			PC += 1
		} else if ArrStr[i] == "Burger" {
			BC += 1
		} else if ArrStr[i] == "Dessert" {
			DC += 1
		} else if ArrStr[i] == "American" {
			AC += 1
		} else if ArrStr[i] == "Italia" {
			IC += 1
		} else if ArrStr[i] == "Mexican" {
			MC += 1
		} else if ArrStr[i] == "India" {
			INC += 1
		} else if ArrStr[i] == "Japan" {
			JC += 1
		} else if ArrStr[i] == "French" {
			FC += 1
		} else if ArrStr[i] == "Africa" {
			AFC += 1
		} else if ArrStr[i] == "BBQ" {
			BBQC += 1
		} else if ArrStr[i] == "Korea" {
			KC += 1
		} else if ArrStr[i] == "Vegan" {
			VC += 1
		} else if ArrStr[i] == "America_Latina" {
			ALC += 1
		} else if ArrStr[i] == "Bakery" {
			BAC += 1
		}
	}

	theme := make(map[string]int)
	theme["Fast_Food"] = FFC
	theme["Pizza"] = PC
	theme["Burger"] = BC
	theme["Dessert"] = DC
	theme["American"] = AC
	theme["Italia"] = IC
	theme["Mexican"] = MC
	theme["India"] = INC
	theme["Japan"] = JC
	theme["French"] = FC
	theme["Africa"] = AFC
	theme["BBQ"] = BBQC
	theme["Korea"] = KC
	theme["Vegan"] = VC
	theme["America_Latina"] = ALC
	theme["Bakery"] = BAC

	pre_order := [5]int{}
	order_theme := [5]string{}

	for i := 0; i < 5; i++ {
		last_add := ""
		for name, v := range theme {

			if i == 0 {
				if v >= pre_order[i] {
					pre_order[i] = v
					order_theme[i] = name
					last_add = name
				}
			} else {
				if v <= pre_order[i-1] && v >= pre_order[i] {
					pre_order[i] = v
					order_theme[i] = name
					last_add = name
				}
			}

		}
		theme[last_add] = -1
	}
	type PopularPost struct {
		FinalArr     []Final
		PopularTheme []string
	}

	ToSend := PopularPost{FinalArr: FinalArr, PopularTheme: order_theme[:]}
	Tpl.ExecuteTemplate(response, "home.html", ToSend)
}
