package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SessionAuth(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// session, err := store.Get(c.Request, "session-name")
	}
}
