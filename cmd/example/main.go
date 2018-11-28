package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memcached"
	"github.com/gin-gonic/gin"
	"github.com/memcachier/mc"
	"github.com/unrolled/secure"

	"github.com/JuanTorr/project/routes"
)

func main() {
	r := gin.Default()
	r.Use(redirectHTTPS)
	client := mc.NewMC("localhost:11211", "", "")
	store := memcached.NewMemcacheStore(client, "", []byte("secret"))
	r.Use(sessions.Sessions("sessions", store))

	api := r.Group("/api")
	routes.API(api)
	// HTTP
	//http://localhost:8080/api/session
	go r.Run(":8080")
	r.RunTLS(":3000", "assets/certs/server.pem", "assets/certs/server.key")
}

func redirectHTTPS(c *gin.Context) {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		SSLHost:     "localhost:3000",
	})
	err := secureMiddleware.Process(c.Writer, c.Request)

	// If there was an error, do not continue.
	if err != nil {
		return
	}
	c.Next()
}
