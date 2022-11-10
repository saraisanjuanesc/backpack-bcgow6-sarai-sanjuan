package handler

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/domains"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/products"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/pkg/web"
)

type request_product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		product, err := p.service.GetByName(ctx, name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request_product

		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		newproduct := domains.Product(req)

		var emptyFiled []string
		values := reflect.ValueOf(req)
		for i := 0; i < values.NumField(); i++ {
			if values.Field(i).Interface() == reflect.Zero(values.Field(i).Type()).Interface() {
				emptyFiled = append(emptyFiled, values.Type().Field(i).Name)
			}
		}

		id, err := p.service.Store(ctx, newproduct)
		if err != nil {
			ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}
		newproduct.ID = id
		if newproduct.ID != 0 {
			ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, newproduct, ""))
			return
		}
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "No se insertó en la tabla"))
		return
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll(ctx)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}

		err = p.service.DeleteS(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}
