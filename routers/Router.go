package routers

import (
	"companyservice/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setRoutes(router)
	return router
}

func setRoutes(router *gin.Engine) {
	healthController := new(controllers.HealthController)
	router.GET("/health", healthController.GetHealth)

	companyController := new(controllers.CompanyController)
	router.GET("/company", companyController.GetAll)
	router.POST("/company", companyController.Create)
	router.GET("/company/:id", companyController.GetOne)
	router.PUT("/company/:id", companyController.Update)
	router.DELETE("/company/:id", companyController.Delete)
}