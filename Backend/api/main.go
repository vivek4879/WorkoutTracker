package main

import (
	internal "WorkoutTracker/internal/database"
	"fmt"
	"net/http"
)

type application struct {
	Models internal.Models
}

func main() {
	app := application{}
	dsn := "postgres://postgres:@localhost/workoutusers?sslmode=disable"
	conn := Connect(dsn)
	app.Models = internal.NewModels(conn)
	fmt.Println("Connected to Database")

	srv := http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}
	fmt.Printf("Listening on port %s\n", srv.Addr)
	_ = srv.ListenAndServe()
}
