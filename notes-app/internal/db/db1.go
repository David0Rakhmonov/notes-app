package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID           int
	Username     string
	PasswordHash string
}

type Note struct {
	ID      int
	UserID  int
	Title   string
	Content string
}

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/notes_app")
	fmt.Print(err)
	if err != nil {
		return fmt.Errorf("ошибка при подключении к базе данных: %v", err)
	}
	fmt.Print(DB)
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password_hash TEXT NOT NULL
	);`)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблиц: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT,
		content TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблицы заметок: %v", err)
	}

	return nil
}

func CreateUser(username, passwordHash string) error {
	InitDB()
	_, err := DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, passwordHash)
	if err != nil {
		return fmt.Errorf("ошибка при создании пользователя: %v", err)
	}
	return nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("пользователь не найден")
		}
		return nil, fmt.Errorf("ошибка при получении пользователя: %v", err)
	}
	return &user, nil
}

func GetNoteByID(id int) (*Note, error) {
	InitDB()
	var note Note
	err := DB.QueryRow("SELECT id, user_id, title, content FROM notes WHERE id = ?", id).
		Scan(&note.ID, &note.UserID, &note.Title, &note.Content)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("заметка с id %d не найдена", id)
		}
		return nil, fmt.Errorf("ошибка при получении заметки: %v", err)
	}
	return &note, nil
}

func UserExists(username string) (bool, error) {
	InitDB()
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	var count int
	err := DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
