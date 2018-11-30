package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/JuanTorr/basic-go-web/db"
	"github.com/JuanTorr/basic-go-web/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := db.GetOrCreateDB()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routes.Analitycs(r, db)
	routes.RedirectHTTPS(r)

	api := r.Group("/api")
	routes.API(api, db)

	//http://localhost:8080/api/session
	go r.Run(":8080")
	r.RunTLS(":3000", "./assets/certs/server.pem", "./assets/certs/server.key")
}
