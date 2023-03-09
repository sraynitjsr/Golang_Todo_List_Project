package router

import (
	"github.com/gorilla/mux"
	"github.com/sraynitjsr/golang-react-todo/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/task", middleware.GetAllTasks.Methods("GET", "OPTIONS"))
	return r
}
