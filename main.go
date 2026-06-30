package main

import (
	"log"
	"net/http"

	"github.com/Sidharth461/crud-mysql-go/database"
	"github.com/Sidharth461/crud-mysql-go/handlers"
)

func main() {
	database.ConnectDB()

	//~~Routes -------------
	http.HandleFunc("/users", handlers.UsersHandler)

	log.Fatal(http.ListenAndServe(":9090", nil))

}
