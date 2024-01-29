package retrieve

import (
	"github.com/bhavye-omniful/GORM/sqs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RecieveMessage(ctx *gin.Context) {
	msgs, err := sqs.RecieveMessageFromSqs()

	if err != nil {
		log.Println("Error receiving messages from queue : ", err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(200, msgs)
}
