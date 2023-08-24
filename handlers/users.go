package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vickyshaw29/events/models"
)

func Signup(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func Login(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user data"})
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}
