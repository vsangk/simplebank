package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // implicitly required for sql.Open to work
	"github.com/vsangk/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

// This TestMain is a pretty common convention to handle initialization code (I think?)
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot connect to the config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
