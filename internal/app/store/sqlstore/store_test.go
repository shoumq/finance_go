package sqlstore_test

import (
	"os"
	"testing"
)

var (
	DatabaseURL string
)

func TestMain(m *testing.M) {
	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		DatabaseURL = "user=postgres port=5433 password=Andrew1095 dbname=finbase_test sslmode=disable"
	}

	os.Exit(m.Run())
}
