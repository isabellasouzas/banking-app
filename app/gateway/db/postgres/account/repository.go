package account

import (
	"github/isabellasouzas/banking-go/app/domain/account"

	"github/isabellasouzas1/src/github.com/jackc/pgx/pgxpool"
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
