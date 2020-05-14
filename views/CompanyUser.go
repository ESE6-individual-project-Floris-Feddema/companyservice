package views

import . "companyservice/models"

type CompanyUser struct {
	User    	User  	`json:"user"`
	CompanyId 	string 	`json:"companyId"`
}