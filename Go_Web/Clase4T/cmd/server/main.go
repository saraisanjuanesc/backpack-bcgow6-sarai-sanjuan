package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase4T/cmd/server/handler"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase4T/docs"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase4T/internal/users"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase4T/pkg/store"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI BOOTCAMP API
// @version 1.0
// @description This API Handle MELI Users
// @termsOfservices https://developers.mercadolibre.com.mx/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.mx/support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	_ = godotenv.Load()

	db := store.NewStore(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewServices(repo)

	us := handler.NewUser(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ur := router.Group("/users")
	ur.POST("/", us.Store())
	ur.GET("/", us.GetAll())
	ur.PUT("/:id", us.Update())
	ur.DELETE("/:id", us.Delete())
	ur.PATCH("/:id", us.UpdateNameLastName())
	router.Run()

}
