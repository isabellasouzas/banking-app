package transfer

import (
	"banking-app/app/domain/usecases/transfer"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ transfer.Repository = &Repository{}

type Repository struct {
	*pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db}
}
