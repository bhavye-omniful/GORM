package create

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateMultipleUser(ctx *gin.Context) {
	var user []models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal("Error binding json to user : ", err.Error())
	}

	log.Println(user)

	result := database.Db.Create(user)
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	ctx.JSON(200, user)
}
