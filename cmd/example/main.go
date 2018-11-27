package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memcached"
	"github.com/gin-gonic/gin"
	"github.com/memcachier/mc"

	"github.com/JuanTorr/project/routes"
)

func main() {
	r := gin.Default()
	client := mc.NewMC("localhost:11211", "", "")
	store := memcached.NewMemcacheStore(client, "", []byte("secret"))
	r.Use(sessions.Sessions("sessions", store))

	api := r.Group("/api")

	routes.API(api)
	r.Run(":3000")
}
