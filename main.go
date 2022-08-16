package main

import (
	"fmt"
	"net/http"
	"waysbuck-API/database"
	"waysbuck-API/pkg/mysql"
	"waysbuck-API/routes"

	"github.com/gorilla/mux"
)

func main() {

	//initial DB
	mysql.DatabaseInit()

	//run imigration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("New Server is running on localhost:2500")
	http.ListenAndServe("localhost:2500", r)
}
