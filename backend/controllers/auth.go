package controllers

import (
	"e-combomb/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body"})
			return
		}

		// Check user exists
		if _, exists := models.InMemoryUsers[newUser.Username]; exists {
			c.JSON(http.StatusBadGateway, gin.H{"message": "this username already exists"})
			return
		}

		// Hash the password
		if hashedPassword, err := models.HashPassword(newUser.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "there's an error on registration"})
			return
		} else {
			// Replace password with hashed password
			newUser.Password = string(hashedPassword)
		}

		models.InMemoryUsers[newUser.Username] = newUser

		c.JSON(http.StatusOK, gin.H{
			"message": "registration successful",
		})
	}
}

func (ac *AuthController) Login(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !models.CheckPassword(user.Username, user.Password) {
			c.JSON(http.StatusForbidden, gin.H{"message": "incorrect password"})
			return
		}

		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			fmt.Println("Error retrieving session:", err)
			c.JSON(http.StatusForbidden, gin.H{"message": "cannot get session from the store login"})
			return
		}

		session.Values["authenticated"] = true
		session.Values["username"] = user.Username
		err = session.Save(c.Request, c.Writer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot login"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
	}
}
