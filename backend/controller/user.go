package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/config"
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
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	var data model.UserUpdate
	cookid := c.GetHeader("id")
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
	var id = c.GetHeader("id")
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
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	var id = c.GetHeader("id")
	memberid, _ := strconv.Atoi(id)
	mentorid := c.Param("id")
	newmentorid, _ := strconv.Atoi(mentorid)
	available := a.userService.CheckMentorAvailable(newmentorid)
	if available {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Berhasil Mengirim Request",
		})
		go func() {
			a.bookService.CreateRequest(memberid, newmentorid)
		}()
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "Mentor not found",
	})

}

func (a *AuthHandler) GetAllBookStatusMember(c *gin.Context) {
	var id = c.GetHeader("id")
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

func (a *AuthHandler) GetAllArtikel(c *gin.Context) {
	data, err := a.userService.ArtikelList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (a *AuthHandler) GetArtikelById(c *gin.Context) {
	var id = c.Param("id")
	newid, _ := strconv.Atoi(id)
	data, err := a.userService.ArtikelDetail(newid)
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

func (a *AuthHandler) GetDataMentor(c *gin.Context) {
	var nobookid = c.Param("bookid")
	data, err := a.userService.GetDataMentor(nobookid)
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
