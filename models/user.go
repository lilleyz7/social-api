package models

type User struct {
	Username   string   `json:"username" bson:"username"`
	Email      string   `json:"email" bson:"email"`
	Password   string   `json:"password" bson:"password"`
	Picture    string   `json:"picture" bson:"picture"`
	Bio        string   `json:"bio" bson:"bio"`
	Name       string   `json:"name" bson:"name"`
	Age        int16    `json:"age" bson:"age"`
	Friends    []string `json:"friends" bson:"friends"`
	UserPosts  []Post   `json:"userposts" bson:"userposts"`
	SavedPosts []Post   `json:"savedposts" json:"savedposts"`
}
