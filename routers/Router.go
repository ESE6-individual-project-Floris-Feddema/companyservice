package routers

import (
	"companyservice/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"PUT", "PATCH", "POST", "GET", "DELETE"}
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setRoutes(router)
	return router
}

func setRoutes(router *gin.Engine) {
	companyController := new(controllers.CompanyController)
	router.GET("/company", companyController.GetAll)
	router.POST("/company", companyController.Create)
	router.GET("/company/:id", companyController.GetOne)
	router.PUT("/company/:id", companyController.Update)
	router.DELETE("/company/:id", companyController.Delete)
	router.GET("/user/:id", companyController.GetALlUser)
	router.POST("/user/:id", companyController.AddUser)
	router.DELETE("/user/:id", companyController.DeleteUser)
}
