package routes

import (
	"backend/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoadRouteSessions(router *gin.Engine) {
	router.GET("/sessions/*url", func(context *gin.Context) {
		filterURL := context.Param("url")[1:]

		_, dbClient, ok := db.QueryDBClients(filterURL)
		if !ok {
			context.JSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"message": "No such client",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": dbClient.Sessions,
		})
	})

	router.PUT("/sessions/:uuid", func(context *gin.Context) {
		filterURL := context.GetHeader("X-Client")
		UUID := context.Param("uuid")
		IP := context.RemoteIP()

		indexDBClient, _, ok := db.QueryDBClients(filterURL)
		if !ok {
			context.JSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"message": "No such client",
			})
			return
		}

		dbSession := db.DBSession{
			UUID: UUID,
			IP: IP,
			Time: int(time.Now().Unix()),
			Inputs: []string{},
		}
		db.DBClients[indexDBClient].Sessions = append(db.DBClients[indexDBClient].Sessions, dbSession)

		err := db.UpdateDBClients()
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

	router.PATCH("/sessions/:uuid", func(context *gin.Context) {
		filterURL := context.GetHeader("X-Client")
		filterUUID := context.Param("uuid")

		var Inputs []string
		err := context.ShouldBindJSON(&Inputs)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "failure",
				"message": "Bad request body",
			})
			return
		}

		indexDBClient, indexDBSession, _, ok := db.QueryDBSessions(filterURL, filterUUID)
		if !ok {
			context.JSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"message": "No such session",
			})
			return
		}

		if len(Inputs[0]) == 0 {
			context.JSON(http.StatusOK, gin.H{
				"status": "success",
			})
			return
		}
		
		db.DBClients[indexDBClient].Sessions[indexDBSession].Inputs = append(db.DBClients[indexDBClient].Sessions[indexDBSession].Inputs, Inputs[0])

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
}