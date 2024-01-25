package retrieve

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUserByStringCond(ctx *gin.Context) {
	var res map[string]models.User
	db := database.Db

	//res[]

	// Get first matched record
	db.Where("name = ?", "ashish").First(res["First record with name ashish"])
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	log.Println(res)

	// Get all matched records
	//db.Where("name <> ?", "jinzhu").Find(&users)
	//// SELECT * FROM users WHERE name <> 'jinzhu';
	//
	//// IN
	//db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	//
	//// LIKE
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//// SELECT * FROM users WHERE name LIKE '%jin%';
	//
	//// AND
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	//
	//// Time
	//db.Where("updated_at > ?", lastWeek).Find(&users)
	//// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	//
	//// BETWEEN
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

	ctx.JSON(200, res)

}
