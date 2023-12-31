package controllers

import (
	"e-combomb/models"
	"e-combomb/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(AuthService *services.AuthService) *AuthController {
	return &AuthController{
		authService: AuthService,
	}
}

func (ac *AuthController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body"})
			return
		}

		if err := ac.authService.CreateUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "there's an error on registration"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "registration successful",
		})
	}
}

func (ac *AuthController) Login(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestUser models.User

		if err := c.ShouldBindJSON(&requestUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		valid, user, err := ac.authService.ValidateUserCredentials(requestUser.Username, requestUser.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "there's an error on login"})
			fmt.Printf("login error: %s\n", err)
			return
		}
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"message": "there's an error on login"})
			fmt.Printf("login invalid")
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
		session.Values["user_id"] = user.ID

		// // Set session expiration
		// session.Options = &sessions.Options{
		// 	Path:     "/",
		// 	MaxAge:   3600,
		// 	HttpOnly: true,
		// 	Secure:   true,
		// 	SameSite: http.SameSiteLaxMode,
		// }

		err = session.Save(c.Request, c.Writer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot login"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
	}
}
