package routes

import (
	"github.com/bhavye-omniful/GORM/database/queries/create"
	"github.com/gin-gonic/gin"
	"log"
)

func Routes() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/createUser", create.CreateUser)
	r.POST("/createMultipleUser", create.CreateMultipleUser)

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
