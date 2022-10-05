package main

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

var listUsers []user

type user struct {
	ID            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	Last_name     string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Age           int    `json:"age" binding:"required"`
	Height        int    `json:"height" binding:"required"`
	Status        bool   `json:"status" binding:"required"`
	Creation_date string `json:"creation_date" binding:"required"`
}

func receivesEntity(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "holatoken" || token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No tiene permisos para realizar la petición solicitada"})
		return
	}
	var us user
	if err := ctx.ShouldBindJSON(&us); err != nil {
		var emptyFiled []string
		values := reflect.ValueOf(us)
		for i := 0; i < values.NumField(); i++ {
			if values.Field(i).Interface() == reflect.Zero(values.Field(i).Type()).Interface() {
				emptyFiled = append(emptyFiled, values.Type().Field(i).Name)
			}
		}
		if len(emptyFiled) > 0 {
			for _, v := range emptyFiled {
				if v != "ID" {
					ctx.JSON(400, gin.H{"error": "El campo: " + v + " es requerido"})
				}
			}
			return
		}
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
