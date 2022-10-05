package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var listUsers []user

type user struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Last_name     string `json:"last_name"`
	Email         string `json:"email"`
	Age           int    `json:"age"`
	Height        int    `json:"height"`
	Status        bool   `json:"status"`
	Creation_date string `json:"creation_date"`
}

func receivesEntity(ctx *gin.Context) {
	var us user
	if err := ctx.ShouldBindJSON(&us); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	us.ID = len(listUsers) + 1
	listUsers = append(listUsers, us)
	ctx.JSON(http.StatusOK, us)
}
func main() {
	router := gin.Default()

	router.POST("", receivesEntity)

	router.Run()
}
