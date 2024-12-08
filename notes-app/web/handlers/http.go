package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"notes-app/internal/auth"
	"notes-app/internal/db"
	"notes-app/internal/notes"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if !auth.IsAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	username := auth.GetCurrentUser(r)

	userNotes, err := notes.GetNotesByUser(username)
	fmt.Print(len(userNotes))
	if err != nil {
		http.Error(w, "Ошибка загрузки заметок", http.StatusInternalServerError)
		return
	}
	fmt.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	tmpl := template.Must(template.ParseFiles("web/notes.html"))

	tmpl.Execute(w, userNotes)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("web/login.html"))
		tmpl.Execute(w, nil)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := auth.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
		return
	}

	err = auth.ValidatePassword(user, password)
	if err != nil {
		http.Error(w, "Неверный пароль", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   username,
		Expires: time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "/notes", http.StatusFound)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("web/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	exists, err := db.UserExists(username)
	if err != nil {
		http.Error(w, "Ошибка проверки пользователя", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Такой пользователь уже существует", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Print("ghhghgh")
	if err != nil {
		http.Error(w, "Ошибка регистрации", http.StatusInternalServerError)
		return
	}

	err = db.CreateUser(username, string(hashedPassword))
	fmt.Print("f")
	if err != nil {
		http.Error(w, "Ошибка регистрации", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
	fmt.Print("f")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})
	http.Redirect(w, r, "/login", http.StatusFound)
}
