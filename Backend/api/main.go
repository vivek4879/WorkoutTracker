package main

import (
	internal "WorkoutTracker/internal/database"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type application struct {
	Models internal.Models
}

func main() {
	app := application{}
	dsn := "postgres://vivekaher:@localhost/workoutusers?sslmode=disable"
	conn := Connect(dsn)
	if conn == nil {
		log.Fatal("Failed to connect to database")
	}
	app.Models = internal.Models{
		UserModel: internal.NewMyModel(conn),
	}
	fmt.Println("Connected to Database")

	MigrateDB(conn)

	srv := http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}
	fmt.Printf("Listening on port %s\n", srv.Addr)
	_ = srv.ListenAndServe()
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/signup", app.signupHandler)
	router.HandlerFunc(http.MethodPost, "/Authenticate", app.AuthenticationHandler)
	router.HandlerFunc(http.MethodGet, "/Dashboard", app.DashboardHandler)
	router.HandlerFunc(http.MethodPost, "/logout", app.logoutHandler)
	router.HandlerFunc(http.MethodDelete, "/deleteAccount", app.deleteHandler)
	router.HandlerFunc(http.MethodPost, "/addWorkout", app.AddWorkoutHandler)
	return router
}
