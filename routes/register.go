package routes

import (
	"net/http"
	"strconv"

	"abc.com/calc/models"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	userId:=context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event ID"})
        return
    }

	event,err:=models.GetById(eventId)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"could not return "})
	}

	event.Register(userId)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"could not return "})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"registered successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId:=context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event ID"})
        return
    }
	var event models.Event
	event.ID=eventId
	err=event.Cancel(userId)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "could not cancel"})
        return
    }
	context.JSON(http.StatusOK,gin.H{"message":"cancelled successfully"})

	
}