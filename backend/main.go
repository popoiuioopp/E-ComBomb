package main

import (
	"e-combomb/routes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte("พี่สาวเชฟนิ้วก้อยแห่ง PikkuuStudio"))
}

func main() {
	router := gin.New()

	routes.SetupRoutes(router, store)

	router.Run()
}
