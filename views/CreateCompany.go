package views

import . "companyservice/models"

type CreateCompany struct {
	Name 	string 	`json:"name"`
	User 	User   	`json:"owner"`
}
