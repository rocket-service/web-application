package postgres

import (
	"context"
	"errors"
	"rocket-web/internal/storage"
	"rocket-web/internal/storage/models"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) SaveUser(ctx context.Context, username, password string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.GetUser(ctx, username)
	if err == nil {
		return 0, storage.ErrUserAlreadyExists
	}

	stmt := "INSERT INTO Users (username, passwordhash) VALUES ($1, $2) RETURNING id"

	var id int64
	err = s.conn.QueryRow(ctx, stmt, username, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetUser(ctx context.Context, username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := "SELECT * FROM Users WHERE username = $1"

	var user models.User
	err := pgxscan.Get(ctx, s.conn, &user, stmt, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, storage.ErrUserNotFound
		}

		return models.User{}, err
	}

	return user, nil
}
