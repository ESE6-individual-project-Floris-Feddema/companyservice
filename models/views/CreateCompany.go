package views

import (
	. "companyservice/models"
)

type CreateCompany struct {
	Name  string `json:"name"`
	Owner User   `json:"owner"`
}
