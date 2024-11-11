package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *pgx.Conn

// Fake data here
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("Cannot connect config:", err)
	}
	fmt.Println("Run:", config.DBDriver)
	testDB, err = pgx.Connect(context.Background(), config.DBProduct)
	if err != nil {
		log.Fatal("Cannot connect db:", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
