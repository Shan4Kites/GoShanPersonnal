package handlers

import (
	"github.com/gorilla/mux"
)

func Router() (router *mux.Router) {
	router = mux.NewRouter()
	buildDockTypeRouters(router)
	return
}

func buildDockTypeRouters(router *mux.Router) {
	prefix := "/docktype"
	router.HandleFunc(prefix, CreateDockTypeHandler).Methods("POST")
}