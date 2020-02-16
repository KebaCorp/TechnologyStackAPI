package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=db_technology_stack_000000_test user=userexample password=passwordexample sslmode=disable"
	}

	os.Exit(m.Run())
}
