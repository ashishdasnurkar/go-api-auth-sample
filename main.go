package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/eaigner/jet"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_SQL"))
	logFatal(err)
	db, err := jet.Open("postgres", pgUrl)
	db.Ping()
	logFatal(err)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/signup", signup).Methods("POST")

	fmt.Println("Server started running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func sendError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error
	json.NewDecoder(r.Body).Decode(&user)
	spew.Dump(user)
	if strings.TrimSpace(user.Email) == "" {
		error.Message = "Email is missing."
		sendError(w, http.StatusBadRequest, error)
		return
	}
	if strings.TrimSpace(user.Password) == "" {
		error.Message = "Password is missing."
		sendError(w, http.StatusBadRequest, error)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		logFatal(err)
	}
	fmt.Printf("Pass plain text:%s\n", user.Password)
	fmt.Printf("Hash:%s", hash)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
