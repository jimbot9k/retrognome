package repository

import (
	"database/sql"
	"log"
	"retrognome/internal/types"
)

type SessionRepository struct {
	DB *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{DB: db}
}

func (s *SessionRepository) CreateSession(session *types.Session) error {
	_, err := s.DB.Exec("INSERT INTO sessions (user_id, token, created_at) VALUES (?, ?, ?)", session.UserID, session.Token, session.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionRepository) GetSessionByToken(token string) *types.Session {
	row := s.DB.QueryRow("SELECT id, user_id, token, created_at FROM sessions WHERE token = ?", token)
	session := &types.Session{}
	err := row.Scan(&session.ID, &session.UserID, &session.Token, &session.CreatedAt)
	if err != nil {
		log.Print(err)
		log.Print(session)
		return &types.Session{}
	}
	log.Print(session)
	return session
}

func (s *SessionRepository) DeleteSessionByToken(token string) error {
	_, err := s.DB.Exec("DELETE FROM sessions WHERE token = ?", token)
	if err != nil {
		return err
	}
	return nil
}
