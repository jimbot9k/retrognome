package repository

import (
	"database/sql"
)

type SessionRepository struct {
	DB *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{DB: db}
}

func (s *SessionRepository) CreateSession(sessionID string, userID int, token string) error {
	s.DB.Exec("INSERT INTO sessions (session_id, user_id, token) VALUES (?, ?, ?)", sessionID, userID, token)
	return nil
}
