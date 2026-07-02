package main

import (
	"log"
	"net/http"

	"github.com/Sidharth461/crud-mysql-go/database"
	"github.com/Sidharth461/crud-mysql-go/handlers"
	"github.com/Sidharth461/crud-mysql-go/middlewares"
	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()
	r := chi.NewRouter()
	r.Use(middlewares.Logging)
	//~~Routes -------------
	// http.HandleFunc("/users", handlers.UsersHandler)
	r.Get("/users", handlers.GetUsers)
	r.Post("/users", handlers.PostUsers)
	r.Put("/users/{id}", handlers.UpdateUsers)
	r.Delete("/users/{id}", handlers.DeleteUser)
	log.Fatal(http.ListenAndServe(":9090", r))

}
