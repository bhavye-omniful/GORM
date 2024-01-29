package create

import (
	"encoding/json"
	"github.com/bhavye-omniful/GORM/sqs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SendMessage(ctx *gin.Context) {
	var body interface{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println("Error binding json : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	msg, err := json.Marshal(body)
	if err != nil {
		log.Println("Error marshalling to json : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = sqs.SendMessageToSqs(msg)
	if err != nil {
		log.Println("Error sending msg to sqs : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"Message":      "Successfully sent message to sqs",
		"Message body": msg,
	})

}
