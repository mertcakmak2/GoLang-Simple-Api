package controller

import (
	"example/gin/config"
	"example/gin/model"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllFaculties(c *gin.Context) {
	var faculties []model.Faculty
	result := config.DB.Preload("Students").Preload("Manager").Find(&faculties)
	if result.Error != nil {
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusOK, faculties)
}

func FindFacultyById(c *gin.Context) {
	id := c.Param("id")
	var faculty model.Faculty
	result := config.DB.Preload("Students").Preload("Manager").Find(&faculty, id)

	if result.Error != nil {
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusOK, faculty)
}

func SaveFaculty(c *gin.Context) {
	var faculty model.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&faculty)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, faculty)
}

func DeleteFaculty(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&model.Faculty{}, id)

	if result.Error != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusNoContent, "deleted")
}
