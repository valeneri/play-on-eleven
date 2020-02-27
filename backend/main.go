package main

import (
	"log"
	"net/http"

	"github.com/valeneri/play-on-eleven/backend/controllers"
	"github.com/valeneri/play-on-eleven/backend/router"
	"github.com/valeneri/play-on-eleven/backend/server"
)

func main() {
	config := server.SetConfig()
	db, err := server.CreateDb(config)
	if err != nil {
		log.Fatal("error connecting db", err)
	}
	r := router.NewRouter()
	reports := controllers.NewReportController(db)
	r.AddRoute(reports.Routes, reports.Prefix)

	port := config.Server.Port
	log.Fatal(http.ListenAndServe(port, r))
}
