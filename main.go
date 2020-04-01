package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   60,
		Secure:   true,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("SESSION_ID", store))

	router.GET("/set", func(c *gin.Context) {
		session := sessions.Default(c)
		r := uuid.New().String()
		fmt.Println(r)
		session.Set("id", r)
		session.Save()
		c.JSON(http.StatusOK, gin.H{})
		return
	})

	router.GET("/get", func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println(session.Get("id"))
		c.JSON(http.StatusOK, gin.H{})
		return
	})

	router.Run(":12345")
}
