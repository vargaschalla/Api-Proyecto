package app

import (
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SesionIndex(c *gin.Context) {
	var lis []models.Sesiones

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func SesionCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	var d models.Sesiones
	//d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}

func SesionGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Sesiones
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}

func SesionUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Sesiones
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}

func SesionDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Sesiones

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
