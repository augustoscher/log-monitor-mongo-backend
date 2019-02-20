package routes

import (
	"bitbucket.org/augustoscher/logs-monitor-api/controller"
	"github.com/gorilla/mux"
)

//InitRoutes init routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = setRoutes(router)
	router = setRoutesAuthenticated(router)
	return router
}

func setRoutesAuthenticated(routes *mux.Router) *mux.Router {
	routes.HandleFunc("/logs", controller.AllLogsEndPoint).Methods("GET")
	routes.HandleFunc("/logs", controller.CreateLogEndPoint).Methods("POST")
	routes.HandleFunc("/logs", controller.DeleteLogEndPoint).Methods("DELETE")
	routes.HandleFunc("/logs/{id}", controller.FindLogEndpoint).Methods("GET")
	routes.HandleFunc("/groupedlogs", controller.AllLogsGroupEndPoint).Methods("GET")
	return routes
}

func setRoutes(routes *mux.Router) *mux.Router {
	routes.HandleFunc("/status", controller.FindStatusEndpoint).Methods("GET")
	return routes
}
