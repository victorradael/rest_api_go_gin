package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorradael/rest_api_go_gin/database"
	"github.com/victorradael/rest_api_go_gin/models"
)

func ShowHome(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func Show404(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"path": c.Request.URL,
	})
}

func JustReturn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetOneStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.Find(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteOneStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.Find(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted",
	})
}

func UpdateOneStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := database.DB.Create(&student).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, student)

}

func SearchStudentsByCpf(c *gin.Context) {
	cpf := c.Param("cpf")
	var student models.Student
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}
