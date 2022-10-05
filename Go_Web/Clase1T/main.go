package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

var listUsers []user

func readFile() []user {
	var listUsers []user
	data, err := os.ReadFile("../users.json")
	if err != nil {
		panic("File not found")
	}
	errUn := json.Unmarshal(data, &listUsers)
	if errUn != nil {
		panic("Error: Unmarshal cannot be performed")
	}
	return listUsers
}

func getAll(ctx *gin.Context) {
	listUsers := readFile()
	var listfilter []user
	queris := ctx.Request.URL.Query()
	for key, value := range queris {
		if key == "name" {
			for _, us := range listUsers {
				if value[0] == us.Name {
					listfilter = append(listfilter, us)
				}

			}
		}
		if key == "last_name" {
			for _, us := range listUsers {
				if value[0] == us.Last_name {
					listfilter = append(listfilter, us)
				}

			}
		}
		if key == "email" {
			for _, us := range listUsers {
				if value[0] == us.Email {
					listfilter = append(listfilter, us)
				}

			}
		}
		if key == "age" {
			qage, err := strconv.Atoi(value[0])
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			for _, us := range listUsers {
				if qage == us.Age {
					listfilter = append(listfilter, us)
				}

			}
		}
		if key == "height" {
			qheight, err := strconv.Atoi(value[0])
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			for _, us := range listUsers {
				if qheight == us.Height {
					listfilter = append(listfilter, us)
				}

			}
		}
		if key == "status" {
			qstatus, err := strconv.ParseBool(value[0])
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			for _, us := range listUsers {
				if qstatus == us.Status {
					listfilter = append(listfilter, us)
				}

			}
		}
		if key == "creation_date" {
			for _, us := range listUsers {
				if value[0] == us.Creation_date {
					listfilter = append(listfilter, us)
				}

			}
		}
	}
	ctx.JSON(http.StatusOK, listfilter)
}

func getOne(ctx *gin.Context) {
	listUsers := readFile()
	idUser, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for _, value := range listUsers {
		if idUser == value.ID {
			ctx.JSON(http.StatusOK, gin.H{"Datos": value})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el ID proporcionado"})
	return

}
func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Sarai",
		})
	})
	router.GET("/users", getAll)
	router.GET("/users/:id", getOne)
	router.Run()
}
