package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Message string
}

func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Healthy",
	})
}

func Flip(c *gin.Context) {
	//tmpl := template.Must(template.ParseFiles("templates/index.html"))

	msg, err := c.Cookie("flip")
	if err != nil {
		c.SetCookie("flip", "Hello", 0, "", "", false, false)
		msg = "Hello"
	} else {
		if msg == "Hello" {
			c.SetCookie("flip", "World", 0, "", "", false, false)
			msg = "World"
		} else {
			c.SetCookie("flip", "Hello", 0, "", "", false, false)
			msg = "Hello"
		}
	}
	data := PageData{Message: msg}
	//err = tmpl.Execute(c.Copy().Writer, data)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "Error loading template.",
	//	})
	//}

	c.HTML(http.StatusOK, "index.html", data)
	//c.HTML(http.StatusOK, "index.html", data)
}
