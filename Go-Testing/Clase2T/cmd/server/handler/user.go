package handler

import (
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase2T/internal/users"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase2T/pkg/web"
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
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		us, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		}
		if len(us) == 0 {
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, "Aun no hay datos para mostrar", ""))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, us, ""))
	}
}

func (u User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		us, err := u.service.Store(req.Name, req.Last_name, req.Email, req.Age, req.Height, req.Status, req.Creation_date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, us, ""))
	}

}
func (u User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		var emptyFiled []string
		values := reflect.ValueOf(req)
		for i := 0; i < values.NumField(); i++ {
			if values.Field(i).Interface() == reflect.Zero(values.Field(i).Type()).Interface() {
				emptyFiled = append(emptyFiled, values.Type().Field(i).Name)
			}
		}
		if len(emptyFiled) > 0 {
			for _, v := range emptyFiled {
				if v != "ID" {
					ctx.JSON(400, web.NewResponse(400, nil, "El campo: "+v+" es requerido"))
				}
			}
			return
		}
		us, err := u.service.Update(int(id), req.Name, req.Last_name, req.Email, req.Age, req.Height, req.Status, req.Creation_date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, us, ""))
	}
}

func (u User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, "Usuario eliminado", ""))
	}
}

func (u User) UpdateNameLastName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del usuario es requerido"))
			return
		}
		if req.Last_name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido del usuario es requerido"))
			return
		}

		us, err := u.service.UpdateNameLastName(int(id), req.Name, req.Last_name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, us, ""))
	}
}
