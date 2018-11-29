package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/JuanTorr/project/controller/api"
)

//API api
func API(r *gin.RouterGroup, db *sql.DB) {
	r.GET("/session", api.Session(db))
	r.POST("/signin", api.Signin(db))
	r.POST("/signup", api.Signup(db))
}
