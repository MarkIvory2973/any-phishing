package routes

import (
	"backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRouteClients(router *gin.Engine) {
	router.GET("/clients", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": db.DBClients,
		})
	})

	router.PUT("/clients/*url", func(context *gin.Context) {
		filterURL := context.Param("url")[1:]

		var dbClient db.DBClient
		err := context.ShouldBindJSON(&dbClient)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "failure",
				"message": "Bad request body",
			})
			return
		}

		_, _, ok := db.QueryDBClients(filterURL)
		if ok {
			context.JSON(http.StatusConflict, gin.H{
				"status": "failure",
				"message": "Client already exists",
			})
			return
		}

		db.DBClients = append(db.DBClients, dbClient)

		err = db.UpdateDBClients()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": "failure",
				"message": "Can't update DB",
			})
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"status": "success",
		})
	})

	router.PATCH("/clients/*url", func(context *gin.Context) {
		filterURL := context.Param("url")[1:]

		var dbClient db.DBClient
		err := context.ShouldBindJSON(&dbClient)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "failure",
				"message": "Bad request body",
			})
			return
		}

		indexDBClient, _, ok := db.QueryDBClients(filterURL)
		if !ok {
			context.JSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"message": "No such client",
			})
		}

		dbClient.Sessions = db.DBClients[indexDBClient].Sessions
		db.DBClients[indexDBClient] = dbClient

		err = db.UpdateDBClients()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": "failure",
				"message": "Can't update DB",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	router.DELETE("/clients/*url", func(context *gin.Context) {
		filterURL := context.Param("url")[1:]

		indexDBClient, _, ok := db.QueryDBClients(filterURL)
		if !ok {
			context.JSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"message": "No such client",
			})
			return
		}

		db.DBClients = append(db.DBClients[:indexDBClient], db.DBClients[indexDBClient+1:]...)

		err := db.UpdateDBClients()
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"status": "failure",
				"message": "Can't update DB",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})
}