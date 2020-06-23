package routers

import (
	"companyservice/controllers"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://0542f3d37a7a44698d586ef10f6fb309@o395097.ingest.sentry.io/5273071",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	router.Use(sentrygin.New(sentrygin.Options{}))

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
	router.DELETE("/user/:id/:userid", companyController.DeleteUser)
}
