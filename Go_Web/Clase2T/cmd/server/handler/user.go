package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase2T/internal/users"
)

type request struct {
	Name          string `json:"name" binding:"required"`
	Last_name     string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Age           int    `json:"age" binding:"required"`
	Height        int    `json:"height" binding:"required"`
	Status        bool   `json:"status" binding:"required"`
	Creation_date string `json:"creation_date" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(us users.Service) *User {
	return &User{
		service: us,
	}

}

func (u User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "holatoken" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		us, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
		}
		ctx.JSON(http.StatusOK, us)
	}
}

func (u User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "holatoken" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		us, err := u.service.Store(req.Name, req.Last_name, req.Email, req.Age, req.Height, req.Status, req.Creation_date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, us)
	}

}
func (u User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "holatoken" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "ID invalido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del usuario es requerido"})
			return
		}
		if req.Last_name == "" {
			ctx.JSON(400, gin.H{"error": "El apellido del usuario es requerido"})
			return
		}
		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "El email del usuario es requerido"})
			return
		}
		if req.Age == 0 {
			ctx.JSON(400, gin.H{"error": "La edad del usuario es requerido"})
			return
		}
		if req.Height == 0 {
			ctx.JSON(400, gin.H{"error": "La altura del usuario es requerido"})
			return
		}
		if req.Creation_date == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de creación del usuario es requerido"})
			return
		}
		us, err := u.service.Update(int(id), req.Name, req.Last_name, req.Email, req.Age, req.Height, req.Status, req.Creation_date)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, us)
	}
}
