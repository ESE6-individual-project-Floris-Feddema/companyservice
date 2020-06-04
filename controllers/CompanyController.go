package controllers

import (
	. "companyservice/models"
	"companyservice/models/views"
	"companyservice/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CompanyController struct{}

func (controller CompanyController) GetAll(c *gin.Context) {
	service := services.CompanyService{}
	companies, err := service.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var returnValue []CompanyDTO
	for _, element := range companies {
		returnValue = append(returnValue, element.DTO())
	}

	c.JSON(http.StatusOK, returnValue)
}

func (controller CompanyController) GetOne(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	service := services.CompanyService{}
	company, err := service.FindOne(objectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, company.DTO())
}

func (controller CompanyController) Create(c *gin.Context) {

	var createCompany views.CreateCompany

	if err := c.ShouldBindJSON(&createCompany); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	insertValue := Company{
		Name:  createCompany.Name,
		Owner: createCompany.Owner,
		Users: []User{},
	}

	service := services.CompanyService{}
	company, err := service.Create(insertValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, company.DTO())
}

func (controller CompanyController) Update(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var updateCompany Company

	if err := c.ShouldBindJSON(&updateCompany); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	service := services.CompanyService{}
	company, err := service.Update(objectId, updateCompany)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, company.DTO())
}

func (controller CompanyController) Delete(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	service := services.CompanyService{}
	err = service.Delete(objectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (controller CompanyController) GetALlUser(c *gin.Context) {
	id := c.Param("id")
	service := services.CompanyService{}
	companies, err := service.FindAllUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	var returnValue []CompanyDTO
	for _, element := range companies {
		returnValue = append(returnValue, element.DTO())
	}

	c.JSON(http.StatusOK, returnValue)
}

func (controller CompanyController) AddUser(c *gin.Context){
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	service := services.CompanyService{}
	err = service.AddUser(objectId, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}