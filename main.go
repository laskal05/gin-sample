package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   60,
		Secure:   true,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("SESSION_ID", store))

	router.GET("/set", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("id", uuid.New().String())
		session.Save()
		c.JSON(http.StatusOK, gin.H{})
		return
	})

	router.GET("/get", func(c *gin.Context) {
		session := sessions.Default(c)
		c.JSON(http.StatusOK, gin.H{
			"id": session.Get("id"),
		})
		return
	})

	router.Run(":12345")
}
