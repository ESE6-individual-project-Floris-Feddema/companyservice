package main

import (
	"companyservice/routers"
	"companyservice/utils"
)

func main() {

	router := routers.InitRoute()
	port := utils.EnvVar("SERVER_PORT")
	router.Run(":" + port)
}
