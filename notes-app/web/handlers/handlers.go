package handlers

import (
	"html/template"
	"net/http"
	"notes-app/internal/auth"
	"notes-app/internal/notes"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.IsAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("web/create_note.html"))
		tmpl.Execute(w, nil)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	username := auth.GetCurrentUser(r)

	err := notes.CreateNote(username, title, content)
	if err != nil {
		http.Error(w, "Ошибка создания заметки", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/notes", http.StatusFound)
}

func EditNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(id)
		note, err := notes.GetNoteByID(id)
		if err != nil {
			http.Error(w, "Ошибка загрузки заметки", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("web/edit_note.html"))
		tmpl.Execute(w, note)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	err := notes.EditNoteByID(id, title, content)
	if err != nil {
		http.Error(w, "Ошибка редактирования заметки", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/notes", http.StatusFound)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := notes.DeleteNoteByID(id)
	if err != nil {
		http.Error(w, "Ошибка удаления заметки", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/notes", http.StatusFound)
}

func RoflHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/notes", http.StatusFound)
}
