package core

type User struct {
	UserId    string             `json:"userid"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `Json:"password" bson:"password"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	DoB       string             `json:"dob" bson:"dob"`
	Avatar    string             `Json:"avatar" bson:"avatar"`
	Address   string             `jsson:"address" address:"address"`
}
