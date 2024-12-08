package notes

import (
	"context"
	"errors"
	"notes-app/internal/db"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UnimplementedNotesServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	_, err = db.DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", req.Username, string(hash))
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{Message: "Пользователь зарегистрирован"}, nil
}

func (s *Service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	var user User
	err := db.DB.QueryRow("SELECT id, password_hash FROM users WHERE username = ?", req.Username).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, errors.New("Неверный логин или пароль")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("Неверный логин или пароль")
	}
	token := "some-jwt-token"
	return &LoginResponse{Token: token}, nil
}

func (s *Service) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	return &LogoutResponse{Message: "Выход выполнен"}, nil
}

func (s *Service) CreateNote(ctx context.Context, req *CreateNoteRequest) (*NoteResponse, error) {
	result, err := db.DB.Exec("INSERT INTO notes (user_id, title, content) VALUES (?, ?, ?)", 1, req.Title, req.Content)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &NoteResponse{Id: int32(id), Title: req.Title, Content: req.Content}, nil
}

func (s *Service) UpdateNote(ctx context.Context, req *UpdateNoteRequest) (*NoteResponse, error) {
	_, err := db.DB.Exec("UPDATE notes SET title = ?, content = ? WHERE id = ? AND user_id = ?", req.Title, req.Content, req.Id, 1)
	if err != nil {
		return nil, err
	}
	return &NoteResponse{Id: req.Id, Title: req.Title, Content: req.Content}, nil
}

func (s *Service) DeleteNote(ctx context.Context, req *DeleteNoteRequest) (*EmptyResponse, error) {
	_, err := db.DB.Exec("DELETE FROM notes WHERE id = ? AND user_id = ?", req.Id, 1)
	if err != nil {
		return nil, err
	}
	return &EmptyResponse{}, nil
}

func (s *Service) ListNotes(ctx context.Context, req *ListNotesRequest) (*ListNotesResponse, error) {
	rows, err := db.DB.Query("SELECT id, title, content FROM notes WHERE user_id = ?", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*NoteResponse
	for rows.Next() {
		var note NoteResponse
		if err := rows.Scan(&note.Id, &note.Title, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}
	return &ListNotesResponse{Notes: notes}, nil
}
