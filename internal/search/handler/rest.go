package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sog01/solid-go/internal/search/model"
	"github.com/sog01/solid-go/internal/search/service"
)

type Rest struct {
	SearchService service.Search
	SyncService   service.Sync
}

func NewRest(searchService service.Search,
	syncService service.Sync) *Rest {
	return &Rest{
		SearchService: searchService,
		SyncService:   syncService,
	}
}

func (rest *Rest) Router(mux *http.ServeMux) {
	mux.HandleFunc("/search", rest.SearchEmployeesHandler)
	mux.HandleFunc("/insert", rest.InsertEmployeeHandler)
	mux.HandleFunc("/update", rest.UpdateEmployeeHandler)
	mux.HandleFunc("/delete", rest.DeleteEmployeeHandler)
	mux.HandleFunc("/health", rest.HealthCheckHandler)
}

func (rest *Rest) ListenAndServe() {
	mux := &http.ServeMux{}
	log.Println("listening server on port 8080")
	rest.Router(mux)
	http.ListenAndServe(":8080", mux)
}

func (rest *Rest) InsertEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var employee *model.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	if err := rest.SyncService.InsertEmployee(employee); err != nil {
		writeResponseInternalError(w, err)
		return
	}
	writeResponseOK(w, employee)
}

func (rest *Rest) UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var employee *model.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	if err := rest.SyncService.UpdateEmployee(employee); err != nil {
		writeResponseInternalError(w, err)
		return
	}
	writeResponseOK(w, employee)
}

func (rest *Rest) DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	if err := rest.SyncService.DeleteEmployee(id); err != nil {
		writeResponseInternalError(w, err)
		return
	}
	writeResponseOK(w, model.Employee{Id: id})
}

func (rest *Rest) SearchEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")
	employees, err := rest.SearchService.SearchEmployees(keyword)
	if err != nil {
		writeResponseInternalError(w, err)
		return
	}
	writeResponseOK(w, employees)
}

func (rest *Rest) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := rest.SyncService.CheckHealth(); err != nil {
		writeResponseInternalError(w, err)
		return
	}
	writeResponseOK(w, map[string]string{
		"status": "OK",
	})
}

func writeResponseOK(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, response)
}

func writeResponseInternalError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	writeResponse(w, map[string]interface{}{
		"error": err,
	})
}

func writeResponse(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}
