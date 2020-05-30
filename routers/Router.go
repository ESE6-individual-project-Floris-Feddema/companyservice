package routers

import (
	"companyservice/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func InitRoute() *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()
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
}
