package routes

import (
	"github.com/bhavye-omniful/GORM/database/queries/create"
	"github.com/bhavye-omniful/GORM/database/queries/retrieve"
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
	r.POST("/createSelectUser", create.CreateSelectUser)
	r.POST("/createOmitUser", create.CreateOmitUser)
	r.POST("/createBatchUser", create.CreateBatchUser)
	r.POST("/createUserFromMap", create.CreateRecordFromMap)
	r.POST("/createWithConflict", create.CreateWithConflict)
	r.POST("/createDay", create.CreateDay)

	r.GET("/getSingleUser", retrieve.GetSingleObject)
	r.GET("/getUserFromID", retrieve.GetUserByID)
	r.GET("/getAllUsers", retrieve.GetAllUsers)
	r.GET("/getUserByStringCondition", retrieve.GetUserByStringCond)

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
