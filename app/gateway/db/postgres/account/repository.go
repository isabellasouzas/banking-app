package account

import (
	"banking-app/app/domain/usecases/account"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ account.Repository = &Repository{}

type Repository struct {
	*pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Pool: db,
	}
}
