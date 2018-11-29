package main

import (
	"database/sql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memcached"
	"github.com/gin-gonic/gin"
	"github.com/memcachier/mc"

	"github.com/JuanTorr/project/routes"

	_ "github.com/lib/pq"
)

func main() {
	db := getDB()
	defer db.Close()

	r := gin.Default()
	client := mc.NewMC("localhost:11211", "", "")
	store := memcached.NewMemcacheStore(client, "", []byte("secret"))
	r.Use(sessions.Sessions("sessions", store))

	routes.Analitycs(r, db)
	routes.RedirectHTTPS(r)

	api := r.Group("/api")
	routes.API(api, db)

	//http://localhost:8080/api/session
	go r.Run(":8080")
	r.RunTLS(":3000", "assets/certs/server.pem", "assets/certs/server.key")
}

func getDB() *sql.DB {
	conData := "host=localhost port=5432 user=postgres password=XXX dbname=sistema sslmode=disable"
	db, err := sql.Open("postgres", conData)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
