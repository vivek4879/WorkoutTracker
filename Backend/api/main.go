package main

import (
	internal "WorkoutTracker/internal/database"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type application struct {
	Models         internal.Models
	PasswordHasher PasswordHasher
}

func main() {

	dsn := "postgres://vivekaher:@localhost/workoutusers?sslmode=disable"
	conn := Connect(dsn)
	if conn == nil {
		log.Fatal("Failed to connect to database")
	}
	app := application{
		Models: internal.Models{
			UserModel: internal.NewMyModel(conn), // Ensure this is not nil
		},
		PasswordHasher: Argon2Hasher{}, // Ensure this is not nil
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

func (app *application) AuthenticationOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.200:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.WriteHeader(http.StatusOK)

}
func (app *application) AddWorkoutOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.200:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}

func (app *application) GetAllExercisesOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.200:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/signup", app.signupHandler)
	router.HandlerFunc(http.MethodPost, "/authenticate", app.AuthenticationHandler)
	router.HandlerFunc(http.MethodOptions, "/authenticate", app.AuthenticationOptionsHandler)
	router.HandlerFunc(http.MethodGet, "/dashboard", app.DashboardHandler)
	router.HandlerFunc(http.MethodPost, "/logout", app.logoutHandler)
	router.HandlerFunc(http.MethodDelete, "/delete-account", app.deleteHandler)
	router.HandlerFunc(http.MethodPost, "/add-workout", app.AddWorkoutHandler)
	router.HandlerFunc(http.MethodOptions, "/add-workout", app.AddWorkoutOptionsHandler)
	router.HandlerFunc(http.MethodGet, "/user-best", app.AddHandler)
	router.HandlerFunc(http.MethodGet, "/exercises", app.GetAllExercisesHandler)
	router.HandlerFunc(http.MethodOptions, "/exercises", app.GetAllExercisesOptionsHandler)
	router.HandlerFunc(http.MethodGet, "/measurements", app.GetMeasurementsHandler)
	router.HandlerFunc(http.MethodPut, "/update-measurements", app.UpdateMeasurementsHandler)
	router.HandlerFunc(http.MethodGet, "/get-streak-data", app.GetStreakDataHandler)

	return router
}
