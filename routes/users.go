package routes

import (
	"net/http"

	"abc.com/calc/models"
	"abc.com/calc/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user) //wokrs like scan from fmt
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse it"})
		return
	}
	err=user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse it"})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message":"user created successsfully"})

}

func login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user) //wokrs like scan from fmt
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse it"})
		return
	}
	err=user.CheckCredentials()

	if err!=nil{
		context.JSON(http.StatusUnauthorized,gin.H{"message":"could not authenticate user"})
		return
	}

	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"could not authenticate user"})
		return
	}


	context.JSON(http.StatusOK,gin.H{"message":"logged in","token":token})



}