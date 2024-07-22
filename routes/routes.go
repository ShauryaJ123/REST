package routes

import (
	"abc.com/calc/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // /events/place any number here
	//called a dynamic path handler
	// authenticated.POST("/events", middlewares.Authenticate, createEvent)
	
	authenticated:=server.Group("/")
	
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id",updateEvent)//put is used for updation
	authenticated.DELETE("/events/:id",deleteEvent)
	authenticated.POST("/events/:id/register",register)
	authenticated.DELETE("/events/:id/register",cancelRegistration)
	
	server.POST("/signup",signup)
	server.POST("/login",login)

}