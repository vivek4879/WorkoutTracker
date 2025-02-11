package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"runtime"
	"time"
)

func hashing(password string) string {
	params := &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}
	hash, erro := argon2id.CreateHash(password, params)
	if erro != nil {
		log.Fatal(erro)
	}
	return hash
}
func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No session token cookie", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	s, err2 := app.Models.UserModel.QUERYSESSION(sessionToken)
	if err2 != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	err3 := app.Models.UserModel.DeleteSession(*s)
	if err3 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	//http.Error(w, "Unauthorized:Session expired", http.StatusUnauthorized)

	fmt.Println("User successfully logged out")
	return
}

func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	dec := json.NewDecoder(r.Body)
	dec.Decode(&input)
	_, err1 := app.Models.UserModel.Query(input.Email)
	if err1 == nil {
		fmt.Printf("%s already exists\n", input.Email)
		return
	}

	hashedPassword := hashing(input.Password)
	err := app.Models.UserModel.Insert(input.FirstName, input.LastName, input.Email, hashedPassword)
	fmt.Println("User created")
	if err != nil {
		fmt.Println("here")
		fmt.Println(err)
	}
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := app.Models.UserModel.Query(input.Email)
	if err != nil {
		fmt.Println(err)
		fmt.Println("User not registered, please Sign up and then login")
		return
	}
	match, err := argon2id.ComparePasswordAndHash(input.Password, user.Password)
	if err != nil {
		log.Printf("Incorrect Password for user: %s ", user.Email)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Printf("User %s matches %t ", user.Email, match)
	//Create new random session token
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(48 * time.Hour)

	err1 := app.Models.UserModel.InsertSession(user.ID, sessionToken, expiresAt)
	if err1 != nil {
		fmt.Println(err1)
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

}
func (app *application) CheckCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No session token cookie", http.StatusUnauthorized)
			return "", errors.New("No session token cookie")
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return "", err
	}
	return cookie.Value, nil
}

func (app *application) CheckToken(w http.ResponseWriter, sessionToken string) (*string, error) {
	s, err := app.Models.UserModel.QUERYSESSION(sessionToken)
	if err != nil {
		http.Error(w, "Unauthorized: Session not found", http.StatusUnauthorized)
		return nil, err
	}
	if s.Expiry.Before(time.Now()) {
		err1 := app.Models.UserModel.DeleteSession(*s)
		if err1 != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil, err1
		}
		http.Error(w, "Unauthorized: Session expired", http.StatusUnauthorized)
		return nil, errors.New("session expired")
	}
	return &s.Token, nil
}

func (app *application) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := app.CheckCookie(w, r)
	if err != nil {
		return // Exit early if error occurs
	}

	_, err3 := app.CheckToken(w, sessionToken)
	if err3 != nil {
		return // Exit early if token is invalid
	}

	fmt.Println("Welcome to the home page")
	_, writeErr := w.Write([]byte("Welcome to the home page!\n"))
	if writeErr != nil {
		fmt.Println("Error writing response:", writeErr)
	}
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := app.CheckCookie(w, r)
	if err != nil {
		return
	}
	session, err2 := app.Models.UserModel.QUERYSESSION(sessionToken)
	if err2 != nil {
		return
	}
	user, err3 := app.Models.UserModel.QueryUserId(session.UserID)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	err4 := app.Models.UserModel.DeleteUser(*user)
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	err5 := app.Models.UserModel.DeleteSession(*session)
	if err5 != nil {
		fmt.Println(err5)
		return
	}
	fmt.Println("User successfully deleted")
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/signup", app.signupHandler)
	router.HandlerFunc(http.MethodPost, "/login", app.loginHandler)
	router.HandlerFunc(http.MethodGet, "/Dashboard", app.DashboardHandler)
	router.HandlerFunc(http.MethodPost, "/logout", app.logoutHandler)
	router.HandlerFunc(http.MethodDelete, "/deleteAccount", app.deleteHandler)
	return router
}
