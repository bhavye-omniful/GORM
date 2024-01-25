package create

import (
	"fmt"
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

func CreateWithConflict(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Println("Error binding json to user : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(user)

	//Do nothing on conflict
	result := database.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	if result.Error != nil {
		log.Println(result.Error.Error())
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		result = database.Db.Create(&models.Log{
			TableName: "users",
			Comment:   fmt.Sprintf("%s Had a conflict inserting user details", user.Name),
		})

		if result.Error != nil {
			log.Println(result.Error.Error())
			ctx.JSON(http.StatusInternalServerError, result.Error.Error())
			return
		}

		ctx.JSON(200, gin.H{"message": "conflict happened"})
		return
	}

	ctx.JSON(200, user)
}
