package models

type User struct {
	UserId		string		`bson:"userId"`
	Name		string		`bson:"name"`
}

type UserDTO struct {
	UserId		string		`json:"userId"`
	Name		string		`json:"name"`
}