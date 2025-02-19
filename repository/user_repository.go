package repository

import (
	"database/sql"
	"errors"
	"to-do-list-app/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)
	return user, nil
}

func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
