package router

import (
	"go-api-student-mysql-gorm/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/create", middleware.Create).Methods("POST")
	router.HandleFunc("/api/readAll", middleware.ReadAll).Methods("GET")
	router.HandleFunc("/api/update", middleware.Update).Methods("PUT")
	router.HandleFunc("/api/delete", middleware.Delete).Methods("DELETE")
	router.HandleFunc("/api/read", middleware.ReadById).Methods("GET")
	return router
}