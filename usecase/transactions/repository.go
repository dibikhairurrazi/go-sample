package repository

import (
	"context"
	"database/sql"
	"go-sample/usecase/transactions"
)

type Repository interface {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) GetSomething(ID uint) interface{} {
	err := r.db.Query("")
	return nil
}