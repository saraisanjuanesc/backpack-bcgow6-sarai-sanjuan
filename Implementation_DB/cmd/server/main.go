package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/cmd/server/handler"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/products"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/pkg/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	_, db := db.ConnectDatabase()
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/:name", p.GetByName())

	r.Run()

}
