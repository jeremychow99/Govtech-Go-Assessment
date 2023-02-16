package main
// ** UNIT TESTS NOT WORKING, PLEASE IGNORE **
// ** UNIT TESTS NOT WORKING, PLEASE IGNORE **
// ** UNIT TESTS NOT WORKING, PLEASE IGNORE **
import (
	"bytes"
	"encoding/json"
	"example/govtech-test/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	server := gin.Default()

	input := struct {
		Teacher string `json:"teacher"`
		Students []string `json:"students"`
	}{
		Teacher: "teacher4234242@gmail.com",
		Students: []string{
			"student1@gmail.com",
			"student2@gmail.com",
		},
	}
	bodyReq, _ := json.Marshal(input)

	// handler
	server.POST("/api/register", controllers.Register)

	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(bodyReq))
	resp := httptest.NewRecorder()

	// serve http
	server.ServeHTTP(resp, req)

	// logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)

	assert.Equal(t, 409, resp.Result().StatusCode)
}

func TestSuspend(t *testing.T){
	server := gin.Default()

	input := struct {
		Student string `json:"student"`
	}{
		Student: "student69@gmail.com",
	}

	// handler
	server.POST("/api/suspend", controllers.Suspend)
	bodyReq, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/api/suspend", bytes.NewBuffer(bodyReq))
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	// logs + assertion
	t.Log(resp.Result().StatusCode)
	t.Log(resp.Body)

	assert.Equal(t, 409, resp.Result().StatusCode)
}