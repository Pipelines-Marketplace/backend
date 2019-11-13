/*
Package api handles the API requests related to tasks
*/
package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Pipelines-Marketplace/backend/pkg/models"
	"github.com/gorilla/mux"
)

// GetAllTasks writes json encoded tasks to ResponseWriter
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allTasks())
}

// GetTaskWithID writes json encoded task to ResponseWriter
func GetTaskWithID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getTask(mux.Vars(r)["name"]))
}

// GetTaskFiles returns a compressed zip with task files
func GetTaskFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	GetCompressedFiles(mux.Vars(r)["name"])
	// Serve the created zip file
	http.ServeFile(w, r, "finalZipFile.zip")
}

// GetAllTags writes json encoded list of tags to Responsewriter
func GetAllTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetAllTags())
}

// GetAllFilteredTasks writes json encoded list of filtered tasks to Responsewriter
func GetAllFilteredTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetAllTasksWithGivenTags(strings.Split(r.FormValue("tags"), "|")))
}
