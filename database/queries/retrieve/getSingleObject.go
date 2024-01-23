package retrieve

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetSingleObject(ctx *gin.Context) {
	var firstUser, lastUser models.User

	result := database.Db.First(&firstUser)
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	result = database.Db.Last(&lastUser)
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"First Record": firstUser,
		"Last Record":  lastUser,
	})
}
