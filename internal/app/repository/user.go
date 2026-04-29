package repository

import (
	"database/sql"
	"log"
	"user/internal/app/entity"
)

type UserRepository interface {
	Store(firstName, lastName, email, phone string) (*entity.User, error)
	FindByCriteria(filter entity.UserFilter) ([]entity.User, error)
}

type userRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return &userRepository{connection}
}

func (r *userRepository) Store(firstName, lastName, email, phone string) (*entity.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, phone) VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email, phone, created_at, updated_at
	`
	var user entity.User

	err := r.connection.QueryRow(query, firstName, lastName, email, phone).Scan(
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

func (r *userRepository) FindByCriteria(filter entity.UserFilter) ([]entity.User, error) {
	var wb whereBuilder

	if filter.ID != nil {
		wb.add("id = $%d", *filter.ID)
	}

	query := `SELECT id, first_name, last_name, email, phone, created_at, updated_at FROM users` + wb.clause()

	rows, err := r.connection.Query(query, wb.args...)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close error: %v", err)
		}
	}()

	var users []entity.User
	for rows.Next() {
		var user entity.User

		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, rows.Err()
}
