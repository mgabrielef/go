package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mgabrielef/GOLANG/controllers"
	"github.com/mgabrielef/GOLANG/database"
	"github.com/mgabrielef/GOLANG/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func RoutesTestSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Name Test", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGreetingsHandler(t *testing.T) {
	r := RoutesTestSetup()
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/Test", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "They should be equal")

	responseMock := `{"API says":"Hello, Test, how are you?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestReturnStudentsHandler(t *testing.T) {
	database.DbConnection()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesTestSetup()
	r.GET("/students", controllers.ReturnStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudentByCPFHandler(t *testing.T) {
	database.DbConnection()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesTestSetup()
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestReturnStudentByIDHandler(t *testing.T) {
	database.DbConnection()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesTestSetup()
	r.GET("/students/:id", controllers.ReturnStudentByID)
	returnPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", returnPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "Name Test", studentMock.Name)
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudentHandler(t *testing.T) {
	database.DbConnection()
	CreateStudentMock()
	r := RoutesTestSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	deletePath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditStudentHandler(t *testing.T) {
	database.DbConnection()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesTestSetup()
	r.PATCH("/students/:id", controllers.EditStudent)
	student := models.Student{Name: "Second Name Test", CPF: "71345678901", RG: "123456700"}
	jsonValue, _ := json.Marshal(student)
	editPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", editPath, bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var updatedStudentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &updatedStudentMock)

	assert.Equal(t, "Second Name Test", updatedStudentMock.Name)
	assert.Equal(t, "71345678901", updatedStudentMock.CPF)
	assert.Equal(t, "123456700", updatedStudentMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}
