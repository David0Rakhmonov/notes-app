package main

import (
	"log"
	"net/http"
	"notes-app/web/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RoflHandler).Methods("GET")
	r.HandleFunc("/notes", handlers.IndexHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler)
	r.HandleFunc("/register", handlers.RegisterHandler)
	r.HandleFunc("/create-note", handlers.CreateNoteHandler).Methods("POST", "GET")
	r.HandleFunc("/edit-note/{id}", handlers.EditNoteHandler).Methods("POST", "GET")
	r.HandleFunc("/delete-note/{id}", handlers.DeleteNoteHandler).Methods("POST", "GET")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST", "GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	log.Println("!")
	http.ListenAndServe(":8082", r)
}
