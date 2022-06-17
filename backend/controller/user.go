package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (a *AuthHandler) GetMentor(c *gin.Context) {
	skill := c.Query("skill")
	if skill == "" {
		data, err := a.userService.GetAllMentor()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   data,
		})
	} else {
		data, err := a.userService.GetMentorBySkill(skill)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 404,
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   data,
		})
	}

}

func (a *AuthHandler) DetailMentor(c *gin.Context) {
	var id = c.Param("id")
	newid, _ := strconv.Atoi(id)
	data, err := a.userService.GetMentorById(newid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   data,
	})
}

func (a *AuthHandler) UpdateUserById(c *gin.Context) {
	var data model.UserUpdate
	cookid, _ := c.Cookie("id")
	newcookid, _ := strconv.Atoi(cookid)
	var id = c.Param("id")
	newid, _ := strconv.Atoi(id)
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}

	err := a.userService.UpdateUserById(&data, newid, newcookid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Update Success",
	})
}

func (a *AuthHandler) UserProfile(c *gin.Context) {
	var id, _ = c.Cookie("id")
	newid, _ := strconv.Atoi(id)
	data, err := a.userService.GetUserDataById(newid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   data,
	})

}

func (a *AuthHandler) GetRequestMentoring(c *gin.Context) {
	var id, _ = c.Cookie("id")
	memberid, _ := strconv.Atoi(id)
	mentorid := c.Param("id")

	newmentorid, _ := strconv.Atoi(mentorid)
	err := a.userService.CreateRequest(memberid, newmentorid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil Mengirim Request",
	})

}

func (a *AuthHandler) GetAllBookStatusMember(c *gin.Context) {
	var id, _ = c.Cookie("id")
	newid, _ := strconv.Atoi(id)
	data, err := a.userService.GetAllBookStatusMember(newid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   data,
	})
}

func (a *AuthHandler) UpdateSatusBooking(c *gin.Context) {
	bookid := c.Param("bookid")
	isfound := a.bookService.CheckFound(bookid)
	if !isfound {
		c.HTML(http.StatusNotFound, "errro.tmpl", gin.H{
			"title": "Not Found",
		})
		return
	}
	err := a.bookService.UpdateStatusBooking(bookid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "accepted.tmpl", gin.H{
		"title": "Page Accepted",
	})

}
