package main

import (
	"backend/db"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initDBClients() {
	err := db.LoadDBClients()
	if err != nil {
		panic(err)
	}
}

func initRouter() {
	router := gin.Default()
	configCORS := cors.DefaultConfig()
	configCORS.AllowAllOrigins = true
	configCORS.AllowHeaders = append(configCORS.AllowHeaders, "X-Client")
	router.Use(cors.New(configCORS))
	
	routes.LoadRouteClients(router)
	routes.LoadRouteSessions(router)
	
	gin.SetMode(gin.ReleaseMode)
	router.Run(":8080")
}

func main() {
	initDBClients()
	initRouter()
}