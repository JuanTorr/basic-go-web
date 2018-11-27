package api

import (
	"fmt"
	"net/http"

	"github.com/JuanTorr/project/bberrors"
	"github.com/JuanTorr/project/model"
	"github.com/JuanTorr/project/repository"
	"github.com/JuanTorr/project/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Session handler with a simple session usage example
func Session(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()
	c.JSON(http.StatusOK, fmt.Sprintf("count: %d", count))
}

//Signin signin handler
func Signin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	userDao := repository.UserDao{}
	u, err := service.AuthenticateUser(userDao, email, password)
	if err != nil {
		if err == bberrors.ErrAuth || err == bberrors.ErrRegNotFound {
			c.JSON(http.StatusUnauthorized, "Ivalid auth")
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, u)
}

//Signup signup handler
func Signup(c *gin.Context) {
	u := model.User{
		Email:         c.PostForm("email"),
		Password:      c.PostForm("password"),
		Name:          c.PostForm("name"),
		FirstSurname:  c.PostForm("firstSurname"),
		SecondSurname: c.PostForm("secondSurname"),
	}

	userDao := repository.UserDao{}
	err := service.RegisterUser(userDao, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
