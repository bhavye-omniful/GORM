package retrieve

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserByID(ctx *gin.Context) {
	var user models.User

	id := ctx.Query("ID")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Param with name {ID} not found",
		})
		return
	}

	result := database.Db.Find(&user, id)
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	ctx.JSON(200, user)
}
