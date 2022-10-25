package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase3T/cmd/server/handler"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase3T/internal/users"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go-Testing/Clase3T/pkg/store"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "HolaToken")

	gin.SetMode(gin.ReleaseMode)
	db := store.NewStore(store.FileType, "userstest.json")
	repo := users.NewRepository(db)
	service := users.NewServices(repo)
	us := handler.NewUser(service)

	router := gin.Default()

	ur := router.Group("/users")
	ur.GET("/", us.GetAll())
	ur.PUT("/:id", us.Update())
	ur.DELETE("/:id", us.Delete())
	return router
}

func createRequeatTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application-json")
	req.Header.Add("token", "HolaToken")

	return req, httptest.NewRecorder()
}

func TestUpdateUser(t *testing.T) {

	r := createServer()

	req, rr := createRequeatTest(http.MethodPut, "/users/1",
		`{"name": "Jose Martin ",
		"last_name": "Juarez López",
		"email": "jose.jualez@mercadolibre.com.mx",
		"age": 32,
		"height": 179,
		"status": true,
		"creation_date": "2012-09-17T10:12:11+06:00"
	}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestDeleteUser(t *testing.T) {
	r := createServer()

	req, rr := createRequeatTest(http.MethodDelete, "/users/3", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	req, rr = createRequeatTest(http.MethodDelete, "/users/3", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
