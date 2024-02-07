package repository

import (
	"database/sql"
	"retrognome/internal/types"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) CreateUser(user *types.User) error {
	u.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, user.Password)
	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (*types.User, error) {
	userRowData := u.DB.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email)
	user := &types.User{}
	userRowData.Scan(&user.ID, &user.Email, &user.Password)
	return user, nil
}
