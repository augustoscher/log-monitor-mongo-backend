package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/augustoscher/logs-monitor-api/dao"
	"bitbucket.org/augustoscher/logs-monitor-api/routes"
)

func main() {
	fmt.Println("Inciando server na porta 3000...")
	dao.Connect()
	fmt.Println("Conexão com mongo realizada com sucesso...")
	router := routes.InitRoutes()
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
