package repository

import (
	"database/sql"
	"user/internal/app/entity"
)

type UserRepository interface {
	Store(firstName, lastName, email, phone string) (*entity.User, error)
}

type userRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return &userRepository{connection}
}

func (r *userRepository) Store(firstName, lastName, email, phone string) (*entity.User, error) {
	sql := `
		INSERT INTO users (first_name, last_name, email, phone) VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email, phone, created_at, updated_at
	`
	var user entity.User

	err := r.connection.QueryRow(sql, firstName, lastName, email, phone).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
