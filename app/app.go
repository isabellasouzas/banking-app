package app

import (
	"context"

	"banking-app/app/config"
	accUsecase "banking-app/app/domain/usecases/account"
	trUsecase "banking-app/app/domain/usecases/transfer"
	"banking-app/app/gateway/db/postgres"
	accRepo "banking-app/app/gateway/db/postgres/account"
	trRepo "banking-app/app/gateway/db/postgres/transfer"
	"banking-app/app/gateway/http/account"
	trHttp "banking-app/app/gateway/http/transfer"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Application struct {
	Account  account.Handler
	Transfer trHttp.Handler
}

func NewApplication(ctx context.Context, cfg config.Config) (*Application, error) {

	p, err := setupDB(ctx, cfg.Database)
	if err != nil {
		return &Application{}, err
	}

	//repository
	accRepo := accRepo.NewRepository(p)
	transferRepo := trRepo.NewRepository(p)

	// usecase
	accUsecase := accUsecase.NewUseCase(accRepo)
	transferUsecase := trUsecase.NewUsecase(transferRepo)

	// handler
	account := account.NewHandler(accUsecase)
	transfer := trHttp.NewHandler(transferUsecase)

	return &Application{
		Account:  account,
		Transfer: transfer,
	}, nil

}

func setupDB(ctx context.Context, cfg config.DatabaseConfigurations) (*pgxpool.Pool, error) {

	url := postgres.CreateDatabaseURL(cfg.User, cfg.Password, cfg.Host, cfg.Name, cfg.SSLMode)
	conn, err := postgres.NewConnection(ctx, url)
	if err != nil {
		return conn, err
	}

	err = postgres.Migrate(conn.Config().ConnConfig)
	if err != nil {
		conn.Close()
		return conn, err
	}

	return conn, nil
}
