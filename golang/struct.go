package sql

import "time"

type Test struct {
	Tableau []string
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Like     int
	Post     int
}

type Post struct {
	Id       int
	Like     int
	Views    int
	Content  string
	Name     string
	Tags     string
	User_id  int
	ViewList string
	LikeList string
}

type Cookie struct {
	User       string
	Mail       string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

type Like struct {
	Like string
	Url  string
}
