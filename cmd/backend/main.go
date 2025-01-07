package main

import (
    "log"
    "net/http"
    "go-http-backend-server/pkg/handlers"
    "go-http-backend-server/pkg/db"
)

func main() {
    db := db.NewDatabase() // Initialize db connection
    handler := handlers.NewUserHandler(db)

    http.HandleFunc("/register", handler.Register)
    http.HandleFunc("/login", handler.Login)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
