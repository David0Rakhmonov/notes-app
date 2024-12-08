package auth

import (
	"fmt"
	"net/http"
	"notes-app/internal/db"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(user *db.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return fmt.Errorf("неверный пароль: %v", err)
	}
	return nil
}

func GetUserByUsername(username string) (*db.User, error) {
	db.InitDB()
	var user db.User
	fmt.Println(db.DB)
	err := db.DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден: %v", err)
	}
	return &user, nil
}

func IsAuthenticated(r *http.Request) bool {
	_, err := r.Cookie("username")
	return err == nil
}

func GetCurrentUser(r *http.Request) string {
	cookie, err := r.Cookie("username")
	if err != nil {
		return ""
	}
	return cookie.Value
}
