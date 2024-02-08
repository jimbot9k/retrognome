package repository

import (
	"database/sql"
	"log"
	"retrognome/internal/types"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) CreateUser(user *types.User) error {
	_, err := u.DB.Exec("INSERT INTO users (email, password, salt) VALUES (?, ?, ?)", user.Email, user.Password, user.Salt)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (*types.User, error) {
	userRowData := u.DB.QueryRow("SELECT id, email, password, salt FROM users WHERE email = ?", email)
	user := &types.User{}

	err := userRowData.Scan(&user.ID, &user.Email, &user.Password, &user.Salt)
	if err == sql.ErrNoRows {
		return user, nil
	} else if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}
