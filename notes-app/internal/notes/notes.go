package notes

import (
	"fmt"
	"notes-app/internal/db"
)

func CreateNote(username, title, content string) error {
	db.InitDB()
	var userID int
	err := db.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		return fmt.Errorf("пользователь не найден: %v", err)
	}

	_, err = db.DB.Exec("INSERT INTO notes (user_id, title, content) VALUES (?, ?, ?)", userID, title, content)
	if err != nil {
		return fmt.Errorf("ошибка при создании заметки: %v", err)
	}
	return nil
}

func GetNotesByUser(username string) ([]*NoteResponse, error) {
	db.InitDB()
	var userID int
	err := db.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден: %v", err)
	}

	rows, err := db.DB.Query("SELECT id, title, content FROM notes WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении заметок: %v", err)
	}
	defer rows.Close()

	var notes []*NoteResponse
	for rows.Next() {
		var note NoteResponse
		if err := rows.Scan(&note.Id, &note.Title, &note.Content); err != nil {
			return nil, fmt.Errorf("ошибка при чтении данных заметки: %v", err)
		}
		notes = append(notes, &note)
	}
	return notes, nil
}

func EditNoteByID(id, title, content string) error {
	db.InitDB()
	_, err := db.DB.Exec("UPDATE notes SET title = ?, content = ? WHERE id = ?", title, content, id)
	if err != nil {
		return fmt.Errorf("ошибка при редактировании заметки: %v", err)
	}
	return nil
}

func DeleteNoteByID(id string) error {
	db.InitDB()
	_, err := db.DB.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении заметки: %v", err)
	}
	return nil
}

func GetNoteByID(id int) (*db.Note, error) {
	db.InitDB()
	return db.GetNoteByID(id)
}
