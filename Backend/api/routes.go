package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"runtime"
)

func (app *application) registerHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	dec := json.NewDecoder(r.Body)
	dec.Decode(&input)
	params := &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}
	hash, erro := argon2id.CreateHash("input.password", params)
	if erro != nil {
		log.Fatal(erro)
	}
	err := app.Models.UserModel.Insert(input.FirstName, input.LastName, input.Email, hash)
	if err != nil {
		fmt.Println(err)
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/home", app.registerHandler)
	return router
}
