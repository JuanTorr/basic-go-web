package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"github.com/JuanTorr/basic-go-web/controller/analytics"
)

//Analitycs monitoring module
func Analitycs(r *gin.Engine, db *sql.DB) {
	r.Use(analytics.AnalizeRequest)
}

//RedirectHTTPS redirects from http to https protocol
func RedirectHTTPS(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
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
	})
}
