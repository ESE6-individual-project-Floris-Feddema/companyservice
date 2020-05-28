package handlers

import (
	"companyservice/models"
	"companyservice/services"
	"encoding/json"
	"log"
)

type UserChangedHandler struct {}

func (handler UserChangedHandler) HandleMessageAsync(body []byte) {
	data := models.User{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	service :=  services.CompanyService{}
	_ = service.UpdateUser(data)
}


