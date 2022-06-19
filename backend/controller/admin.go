package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (a *AuthHandler) RegisMentor(c *gin.Context) {
	var data model.MentorRegis
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}
	err := a.userService.RegisMentor(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Register Mentor Success",
	})

}

func (a *AuthHandler) DeleteById(c *gin.Context) {
	data := c.Param("id")
	id, _ := strconv.Atoi(data)

	err := a.adminService.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Delete Success",
	})
}

func (a *AuthHandler) CreateArtikel(c *gin.Context) {
	var data model.Artikel
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}
	err := a.adminService.CreateArtikel(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil Menambahkan Artikel",
	})

}
