package repository

import (
	"database/sql"
	"retrognome/internal/types"
)

type RetroRepository struct {
	DB *sql.DB
}

func NewRetroRepository(db *sql.DB) *RetroRepository {
	return &RetroRepository{DB: db}
}

func (r *RetroRepository) GetAllRetros() ([]*types.Retro, error) {
	rows, err := r.DB.Query("SELECT * FROM retros")
	if err != nil {
		return nil, err
	}

	retros := []*types.Retro{}
	for rows.Next() {
		retro := &types.Retro{}
		err = rows.Scan(&retro.ID, &retro.Title, &retro.Description)
		if err != nil {
			return nil, err
		}
		retros = append(retros, retro)
	}
	return retros, nil
}

func (r *RetroRepository) CreateRetro(retro *types.Retro) (*types.Retro, error) {
	result, err := r.DB.Exec("INSERT INTO retros (title, description) VALUES (?, ?)", retro.Title, retro.Description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	retro.ID = int(id)
	return retro, nil
}

func (r *RetroRepository) GetRetroByID(id int) (*types.Retro, error) {
	// Implement logic to retrieve a Retro by ID from the database
	// Use r.DB.QueryRow() or a similar method
	return nil, nil
}

func (r *RetroRepository) UpdateRetro(retro *types.Retro) error {
	// Implement logic to update a Retro in the database
	// Use r.DB.Exec() or a similar method
	return nil
}

func (r *RetroRepository) DeleteRetro(id int) error {
	// Implement logic to delete a Retro by ID from the database
	// Use r.DB.Exec() or a similar method
	return nil
}
