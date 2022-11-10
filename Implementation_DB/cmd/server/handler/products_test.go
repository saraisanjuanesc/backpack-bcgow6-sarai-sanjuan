package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/products"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/pkg/db"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASSWORD", "")
	os.Setenv("DBNAME", "storage")

	gin.SetMode(gin.ReleaseMode)

	_, db := db.ConnectDatabase()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.GET("/:name", p.GetByName())
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.DELETE("/:id", p.Delete())

	return r
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestStore(t *testing.T) {
	r := createServer()
	req, rr := createRequest(http.MethodPost, "/api/v1/products/", `
	{
		"nombre": "nuevo producto",
		"tipo": "celulares",
		"cantidad": 7,
		"precio": 934.5
	 }`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

}

func TestGetByName(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodGet, "/api/v1/products/Producto1", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAll(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodGet, "/api/v1/products/", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDelete(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodDelete, "/api/v1/products/12", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
