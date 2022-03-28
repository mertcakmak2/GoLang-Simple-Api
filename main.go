package main

import (
	"example/gin/config"
	"example/gin/controller"
	"example/gin/model"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=root dbname=db password=root sslmode=disable")
	defer db.Close()
	config.DB = db
	if err != nil {
		log.Fatalln(err)
	}

	config.DB.AutoMigrate(&model.Faculty{})
	config.DB.AutoMigrate(&model.Student{})
	config.DB.AutoMigrate(&model.Manager{})

	router := gin.Default()

	// Faculty
	router.GET("/faculties", controller.FindAllFaculties)
	router.GET("/faculties/:id", controller.FindFacultyById)
	router.POST("/faculties", controller.SaveFaculty)
	router.DELETE("/faculties/:id", controller.DeleteFaculty)

	router.Run(":8091")
}
