package db

import (
	"AnalyticsAndReporting/util"
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	// Data
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot connect to config file: ", err)
	}

	// DB
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
