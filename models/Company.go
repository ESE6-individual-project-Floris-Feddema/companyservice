package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyModel struct {
	Name	string 				`bson:"name"`
	Owner	User				`bson:"owner"`
	Users   []User  			`bson:"users"`
}

type Company struct {
	ID		primitive.ObjectID	`bson:"_id"`
	Name	string 				`bson:"name"`
	Owner	User				`bson:"owner"`
	Users   []User  			`bson:"users"`
}

type CompanyDTO struct {
	Id    	string 				`json:"id"`
	Name  	string 				`json:"name"`
	Owner 	User 				`json:"owner"`
	Users 	[]User  			`json:"users"`
}

func (c Company) Model() CompanyModel {
	model := CompanyModel{
		Name:  c.Name,
		Owner: c.Owner,
		Users: c.Users,
	}
	return model
}

func (c Company) DTO() CompanyDTO {
	dto := CompanyDTO{
		Id:    c.ID.Hex(),
		Name: c.Name,
		Owner: c.Owner,
		Users: c.Users,
	}
	return dto
}

func (c Company) Empty() Company {
	company := Company{}
	return company
}