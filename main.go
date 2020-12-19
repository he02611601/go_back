package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"go_back/util"
)

type IndexData struct {
	Title   string
	Content string
}

func loginPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginAuth(c *gin.Context)  {
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	if userName == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱以及密碼！"),
		})
		return
	}

	if err := auth.Auth(userName, password); err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"success": "登入成功！",
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/html/*")
	r.Static("/assets", "template/assets")
	r.GET("/login", loginPage)
	r.POST("login", loginAuth)
	r.Run(":8888")
}