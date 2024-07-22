package routes

import (
	"net/http"
	"strconv"

	"abc.com/calc/models"
	// "abc.com/calc/utils"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	//gives the path parameter  context.Param("id")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "lol"})
		return
	}
	event, err := models.GetById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "lol"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "lol"})
		return
	}
	context.JSON(http.StatusOK, events)
	//context.HTML()//also possible
}

func createEvent(context *gin.Context) {

	// token:=context.Request.Header.Get("Authorization")

	// if token==""{
	// 	context.JSON(http.StatusUnauthorized,gin.H{"message":"not authorized"})
	// 	return
	// }

	// userId,err:=utils.VerifyToken(token)
	// if err!=nil{
	// 	context.JSON(http.StatusUnauthorized,gin.H{"message":"not authorized"})
	// 	return
	// }
	


	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event) //wokrs like scan from fmt
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse it"})
		return
	}
	userId:=context.GetInt64("userId")

	event.UserId = int(userId)

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "lol"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

// func updateEvent(context *gin.Context){
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err!=nil{
// 		context.JSON(http.StatusBadRequest,gin.H{"message":"cannot delete what is not there"})
// 	}
// 	_,err=models.GetById(eventId)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "such an event does not exist"})
// 		return
// 	}

// 	var updatedEvent models.Event
// 	err = context.ShouldBindJSON(&updatedEvent)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "failure point"})
// 		return
// 	}

// 	updatedEvent.ID=eventId
// 	err=updatedEvent.Update()
// 	if err!=nil{
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "not able to update"})
// 		return 
// 	}
// 	context.JSON(http.StatusOK,gin.H{"message":"event updated sucessfully"})

// }


func updateEvent(context *gin.Context) {
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event ID"})
        return
    }
	userId:=context.GetInt64("userId")
    event, err := models.GetById(eventId)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "event does not exist"})
        return
    }

	if event.UserId!=int(userId){
		context.JSON(http.StatusUnauthorized,gin.H{"message":"not authorized"})
		return
	}

    var updatedEvent models.Event
    err = context.ShouldBindJSON(&updatedEvent)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "invalid JSON payload", "error": err.Error()})
        return
    }

    updatedEvent.ID = eventId
    err = updatedEvent.Update()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "unable to update event", "error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event ID"})
        return
    }

	userIdd:=context.GetInt64("userId")
    theEvent, err := models.GetById(eventId)

	if theEvent.UserId!=int(userIdd){
		context.JSON(http.StatusUnauthorized,gin.H{"message":"not authorized"})
		return
	}

	
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "event does not exist"})
        return
    }
	err=theEvent.DeleteIt()
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "event does not exist"})
        return
    }
	context.JSON(http.StatusOK,gin.H{"message":"deleted successfully"})
}