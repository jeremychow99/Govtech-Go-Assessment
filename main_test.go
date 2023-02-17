package main

// ** UNIT TESTS NOT WORKING, PLEASE IGNORE **
// ** UNIT TESTS NOT WORKING, PLEASE IGNORE **
// ** UNIT TESTS NOT WORKING, PLEASE IGNORE **
import (
	"bytes"
	"encoding/json"
	"example/govtech-test/controllers"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	var DB *gorm.DB
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	fmt.Println(DB)
	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	server := gin.Default()

	input := struct {
		Teacher  string   `json:"teacher"`
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

	// if one or more alr registered, should be 409, if success then 204
	assert.Equal(t, 409, resp.Result().StatusCode)
}

func TestSuspend(t *testing.T) {
	server := gin.Default()

	input := struct {
		Student string `json:"student"`
	}{
		Student: "student70@gmail.com",
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

	// if student already suspended code is 409, if success then 204
	assert.Equal(t, 409, resp.Result().StatusCode)
}

// func TestRetrieve(){
// 	// if success code 200, on error should be 400
// 	assert.Equal(t, 200, resp.Result().StatusCode)
// }