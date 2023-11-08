package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eaigner/jet"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
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
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_UR¨"))
	logFatal(err)
	db, err := jet.Open("postgres", pgUrl)
	logFatal(err)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/signup", signup).Methods("POST")

	fmt.Println("Server started running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Singup invoked")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
