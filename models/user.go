package models

type User struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Picture  string   `json:"picture"`
	Bio      string   `json:"bio"`
	Name     string   `json:"name"`
	Age      int16    `json:"age"`
	Friends  []string `json:"friends"`
	Posts    []Post   `json:"posts"`
}
