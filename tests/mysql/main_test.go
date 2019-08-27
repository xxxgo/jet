package mysql

import (
	"database/sql"
	"flag"
	"github.com/xxxgo/jet/tests/dbconfig"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pkg/profile"
	"os"
	"testing"
)

var db *sql.DB

var source string

const MariaDB = "MariaDB"

func init() {
	flag.StringVar(&source, "source", "", "MySQL or MariaDB")
	flag.Parse()
}

func sourceIsMariaDB() bool {
	return source == MariaDB
}

func TestMain(m *testing.M) {
	defer profile.Start().Stop()

	var err error
	db, err = sql.Open("mysql", dbconfig.MySQLConnectionString)
	if err != nil {
		panic("Failed to connect to test db" + err.Error())
	}
	defer db.Close()

	ret := m.Run()

	os.Exit(ret)
}
