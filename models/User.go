package models

type User struct {
	UserId string `json:"userId" bson:"userId"`
	Name   string `json:"name" bson:"name"`
}

type UserDTO struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}
