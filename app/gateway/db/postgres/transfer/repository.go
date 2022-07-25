package transfer

import (
	"github/isabellasouzas/banking-go/app/domain/transfer"

	"github/isabellasouzas1/src/github.com/jackc/pgx/pgxpool"
)

var _ transfer.Repository = &Repository{}

type Repository struct {
	*pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db}
}
