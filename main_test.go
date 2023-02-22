package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/victorradael/rest_api_go_gin/controllers"
	"github.com/victorradael/rest_api_go_gin/database"
	"github.com/victorradael/rest_api_go_gin/models"
)

var ID int

func SetupTestsRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func createMockStudent() {
	student := models.Student{Name: "Student Test", CPF: "12345678901"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func deleteMockStudent() {
	// var student models.Student
	database.DB.Exec("DELETE FROM students WHERE id=" + strconv.Itoa(ID))
}

func TestGetAllStudents(t *testing.T) {
	database.ConnectWithDatabase()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestsRoutes()
	r.GET("/api/students", controllers.GetAllStudents)
	req, _ := http.NewRequest("GET", "/api/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")

}

func TestSearchStudentByCpf(t *testing.T) {
	database.ConnectWithDatabase()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestsRoutes()
	r.GET("/api/students/cpf/:cpf", controllers.SearchStudentsByCpf)
	req, _ := http.NewRequest("GET", "/api/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")

}

func TestGetOneStudentById(t *testing.T) {
	database.ConnectWithDatabase()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestsRoutes()
	r.GET("/students/:id", controllers.GetOneStudentById)
	pathDaBusca := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var studentMock models.Student
	json.Unmarshal(resposta.Body.Bytes(), &studentMock)
	assert.Equal(t, "Student Test", studentMock.Name, "the names should be equals")
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeleteOneStudentById(t *testing.T) {
	database.ConnectWithDatabase()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestsRoutes()
	r.DELETE("/students/:id", controllers.DeleteOneStudentById)
	pathDeBusca := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestUpdateOneStudentById(t *testing.T) {
	database.ConnectWithDatabase()
	createMockStudent()
	defer deleteMockStudent()
	r := SetupTestsRoutes()
	r.PATCH("/students/:id", controllers.UpdateOneStudentById)
	student := models.Student{Name: "Student Test", CPF: "47123456789"}
	valorJson, _ := json.Marshal(student)
	pathParaEditar := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var studentMockAtualizado models.Student
	json.Unmarshal(resposta.Body.Bytes(), &studentMockAtualizado)
	assert.Equal(t, "47123456789", studentMockAtualizado.CPF)
	assert.Equal(t, "Student Test", studentMockAtualizado.Name)
}
