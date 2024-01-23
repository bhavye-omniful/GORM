package create

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateRecordFromMap(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Println("Error binding json to user : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(user)

	result := database.Db.Model(&models.User{}).Create(map[string]interface{}{
		"Name":    user.Name,
		"Email":   user.Email,
		"HouseNo": user.Address.HouseNo,
		"City":    user.Address.City,
		"State":   user.Address.State,
		"Country": user.Address.Country,
	})
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	ctx.JSON(200, user)
}
