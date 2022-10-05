package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name          string `json:"name"`
	Last_name     string `json:"last_name"`
	Email         string `json:"email"`
	Age           int    `json:"age"`
	Height        int    `json:"height"`
	Status        bool   `json:"status"`
	Creation_date string `json:"creation_date"`
}

func getAll(ctx *gin.Context) {
	var listUsers []user
	data, err := os.ReadFile("../users.json")
	if err != nil {
		panic("File not found")
	}
	errUn := json.Unmarshal(data, &listUsers)
	if errUn != nil {
		panic("Error: Unmarshal cannot be performed")
	}
	ctx.JSON(http.StatusOK, listUsers)
}
func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Sarai",
		})
	})
	router.GET("/users", getAll)
	router.Run()
}
