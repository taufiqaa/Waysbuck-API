package routes

import (
	"waysbuck-API/handlers"
	"waysbuck-API/pkg/middleware"
	"waysbuck-API/pkg/mysql"
	"waysbuck-API/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/toppings", middleware.Auth(h.FindProducts)).Methods("GET")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.GetProduct)).Methods("GET")
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
}
