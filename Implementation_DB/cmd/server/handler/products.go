package handler

import (
	"net/http"

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
		//ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, newproduct, ""))
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
