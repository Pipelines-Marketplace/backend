package main

import (
	"log"
	"net/http"

	"github.com/Pipelines-Marketplace/backend/pkg/api"
	"github.com/Pipelines-Marketplace/backend/pkg/models"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Add new routers
	if err := models.StartConnection(); err != nil {
		log.Fatalln(err)
	}
	// models.CreateDatabase()
	// models.AddContentsToDB()
	router.HandleFunc("/task/{name}", api.GetTaskFiles).Methods("GET")
	router.HandleFunc("/tags", api.GetAllTags).Methods("GET")
	// router.HandleFunc("/tasks?{tags:list}", api.GetAllFilteredTasks).Methods("GET")
	router.Path("/tasks").Queries("tags", "{tags}").HandlerFunc(api.GetAllFilteredTasks).Methods("GET")
	router.HandleFunc("/tasks", api.GetAllTasks).Methods("GET")
	http.ListenAndServe(":5000", router)
}
