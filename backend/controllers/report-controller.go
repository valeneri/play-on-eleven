package controllers

import (
	"encoding/json"
	"fmt"
	"play-on-eleven/backend/daos"
	"play-on-eleven/backend/models"
	"play-on-eleven/backend/router"
	"play-on-eleven/backend/server"

	"net/http"
)

const prefix = "/reports"

type ReportController struct {
	ReportDao *daos.ReportDao
	Routes    []router.Route
	Prefix    string
}

func NewReportController(db *server.Database) *ReportController {
	// build routes
	routes := []router.Route{}

	controller := ReportController{
		ReportDao: daos.NewReportDao(db),
		// Routes:    routes,
		Prefix: prefix,
	}
	// fetchAllReports
	routes = append(routes, router.Route{
		Name:        "Get all reports",
		Method:      http.MethodGet,
		Pattern:     "",
		HandlerFunc: controller.fetchAllReports,
	})

	routes = append(routes, router.Route{
		Name:        "Create new report",
		Method:      http.MethodPost,
		Pattern:     "",
		HandlerFunc: controller.createReport,
	})

	controller.Routes = routes

	return &controller
}

func (c *ReportController) fetchAllReports(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetch all reports requested")
	reports, _ := c.ReportDao.FetchAllReports()
	err := json.NewEncoder(w).Encode(reports)
	if err != nil {
		panic(err)
	}
}

func (c *ReportController) createReport(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	report := &models.Report{}

	err := json.NewDecoder(r.Body).Decode(report)
	if err != nil {
		json.NewEncoder(w).Encode("error")
	}
	err = c.ReportDao.CreateReport(report)
	if err != nil {
		json.NewEncoder(w).Encode("error")
	}
	fmt.Print(json.NewEncoder(w).Encode(report))
	json.NewEncoder(w).Encode(report)
}
