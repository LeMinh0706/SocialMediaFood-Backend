package dump

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

// Fake data here
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot connect config:", err)
	}
	fmt.Println("Run:", config.DBDriver)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect db:", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
