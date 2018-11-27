package routes

import (
	"github.com/JuanTorr/project/controller/api"
	"github.com/gin-gonic/gin"
)

//API api
func API(r *gin.RouterGroup) {
	r.GET("/session", api.Session)
	r.POST("/signin", api.Signin)
	r.POST("/signup", api.Signup)
}
