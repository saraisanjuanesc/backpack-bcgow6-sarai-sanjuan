package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase2T/cmd/server/handler"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase2T/internal/users"
)

func main() {
	repo := users.NewRepository()
	service := users.NewServices(repo)

	us := handler.NewUser(service)

	router := gin.Default()

	ur := router.Group("/users")
	ur.POST("/", us.Store())
	ur.GET("/", us.GetAll())
	ur.PUT("/:id", us.Update())
	router.Run()

}
