package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context) error {
	return nil
}