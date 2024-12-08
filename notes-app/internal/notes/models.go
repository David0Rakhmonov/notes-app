package notes

type User struct {
	ID           int
	Username     string
	PasswordHash string
}

type Note struct {
	ID        int
	UserID    int
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt string
}
