package services

import (
	. "companyservice/models"
	"companyservice/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService struct{}

func (service CompanyService) Create(company Company) (*Company, error) {
	repository := repositories.CompanyRepository{}
	returnValue, err := repository.Create(&company)
	return returnValue, err

}

func (service CompanyService) FindAll() ([]*Company, error) {
	repository := repositories.CompanyRepository{}
	companies, err := repository.FindAll()
	return companies, err
}

func (service CompanyService) FindOne(id primitive.ObjectID) (*Company, error) {
	repository := repositories.CompanyRepository{}
	company, err := repository.FindOne(id)
	return company, err
}

func (service CompanyService) Delete(id primitive.ObjectID) error {
	repository := repositories.CompanyRepository{}
	err := repository.Delete(id)
	return err
}

func (service CompanyService) Update(id primitive.ObjectID, company Company) (*Company, error) {
	repository := repositories.CompanyRepository{}
	returnValue, err := repository.Update(id, company)
	return returnValue, err

}

func (service CompanyService) FindAllUser(id string) ([]*Company, error) {
	repository := repositories.CompanyRepository{}
	returnValue, err := repository.FindAllUser(id)
	return returnValue, err
}

func (service CompanyService) AddUser(id primitive.ObjectID, user User) error {
	repository := repositories.CompanyRepository{}
	err := repository.AddUser(id, user)
	return err
}

func (service CompanyService) UpdateUser(user User) error {
	repository := repositories.CompanyRepository{}
	err := repository.UpdateUser(user)
	return err
}

func (service CompanyService) DeleteUser(id primitive.ObjectID, userId string) error {
	repository := repositories.CompanyRepository{}
	err := repository.DeleteUser(id, userId)
	return err
}
