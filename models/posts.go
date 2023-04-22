package models

type Post struct {
	Username    string    `json:"username" bson:"username"`
	Picture_url string    `json:"picture_url" bson:"picture_url"`
	Title       string    `json:"title" bson:"title"`
	Likes       int       `json:"likes" bson:"likes"`
	Comments    []Comment `json:"comments" bson:"comments"`
}

type Comment struct {
	Username string `json:"username" bson:"username"`
	Content  string `json:"content" bson:"content"`
}
