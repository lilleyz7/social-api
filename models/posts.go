package models

type Post struct {
	Username    string    `json:"username"`
	Picture_url string    `json:"picture_url"`
	Caption     string    `json:"caption"`
	Likes       int       `json:"likes"`
	Comments    []Comment `json:"comments"`
}

type Comment struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
