package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/JuanTorr/basic-go-web/model"
	"github.com/JuanTorr/basic-go-web/repository"
	"github.com/JuanTorr/basic-go-web/service"
	"github.com/JuanTorr/basic-go-web/utils/bgwerrors"
)

//Session handler with a simple session usage example
func Session(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

//Signin signin handler
func Signin(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		userDao := repository.UserDao{DB: db}
		u, err := service.AuthenticateUser(userDao, email, password)
		if err != nil {
			switch err.(type) {
			case bgwerrors.ErrAuth, bgwerrors.ErrRegNotFound:
				c.JSON(http.StatusUnauthorized, err.Error())
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}
		c.JSON(http.StatusOK, u)
	}
}

//Signup signup handler
func Signup(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		u := model.User{
			Email:    c.PostForm("email"),
			Password: c.PostForm("password"),
			Name:     c.PostForm("name"),
			Surname:  c.PostForm("surname"),
		}
		u.Role.ID, err = strconv.Atoi(c.PostForm("roleId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		userDao := repository.UserDao{DB: db}
		err = service.RegisterUser(userDao, u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	}
}
