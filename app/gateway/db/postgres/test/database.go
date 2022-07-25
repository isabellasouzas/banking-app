package test

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"testing"

	"banking-go/app/config"
	"banking-go/app/gateway/db/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewTestDatabase(t *testing.T) Database {
	test := t.Name()
	name := createRandomName()

	if (env == config.DatabaseConfigurations{}) {
		log.Fatal("SetupTest should be called first")
	}

	err := createDatabase(name)
	if err != nil {
		log.Fatalf("Could not setup database for test %s: %v", test, err)
	}

	pool, err := connect(name, env)
	if err != nil {
		log.Fatalf("Could not setup database for test %s: %v", test, err)
	}

	return Database{
		Name:     name,
		TestName: test,
		Pool:     pool,
		Context:  context.Background(),
	}
}

type Database struct {
	Name     string
	TestName string
	Pool     *pgxpool.Pool
	Context  context.Context
}

func (d Database) Drop() {
	const script = `drop database if exists `
	d.Pool.Close()

	_, err := db.Exec(d.Context, script+d.Name)
	if err != nil {
		log.Printf("Could not drop database for test %s: %v", d.TestName, err)
	}
}

func createRandomName() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	return fmt.Sprintf("database_%d", n)
}

func createDatabase(name string) error {
	_, err := db.Exec(context.Background(), "create database "+name)
	return err
}

func connect(database string, env config.DatabaseConfigurations) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := postgres.ConnectAndMount(ctx, postgres.Config{
		User:     env.User,
		Password: env.Password,
		Host:     env.Host,
		Name:     database,
		SSLMode:  env.SSLMode,
	})

	return pool, err
}
