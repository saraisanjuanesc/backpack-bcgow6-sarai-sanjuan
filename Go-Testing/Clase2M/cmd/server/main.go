package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase2M/cmd/server/handler"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase2M/internal/users"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase2M/pkg/store"
)

func main() {

	_ = godotenv.Load()

	db := store.NewStore(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewServices(repo)

	us := handler.NewUser(service)

	router := gin.Default()

	ur := router.Group("/users")
	ur.POST("/", us.Store())
	ur.GET("/", us.GetAll())
	ur.PUT("/:id", us.Update())
	ur.DELETE("/:id", us.Delete())
	ur.PATCH("/:id", us.UpdateNameLastName())
	router.Run()

}
