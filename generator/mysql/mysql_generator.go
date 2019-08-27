package mysql

import (
	"database/sql"
	"fmt"
	"github.com/xxxgo/jet/v2/generator/internal/metadata"
	"github.com/xxxgo/jet/v2/generator/internal/template"
	"github.com/xxxgo/jet/v2/internal/utils"
	"github.com/xxxgo/jet/v2/mysql"
	"path"
)

// DBConnection contains MySQL connection details
type DBConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	Params   string

	DBName string
}

// Generate generates jet files at destination dir from database connection details
func Generate(destDir string, dbConn DBConnection) error {
	db, err := openConnection(dbConn)
	if err != nil {
		return err
	}
	defer utils.DBClose(db)

	fmt.Println("Retrieving database information...")
	// No schemas in MySQL
	dbInfo, err := metadata.GetSchemaInfo(db, dbConn.DBName, &mySqlQuerySet{})

	if err != nil {
		return err
	}

	genPath := path.Join(destDir, dbConn.DBName)

	err = template.GenerateFiles(genPath, dbInfo.TableInfos, dbInfo.EnumInfos, mysql.Dialect)

	if err != nil {
		return err
	}

	return nil
}

func openConnection(dbConn DBConnection) (*sql.DB, error) {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConn.User, dbConn.Password, dbConn.Host, dbConn.Port, dbConn.DBName)
	if dbConn.Params != "" {
		connectionString += "?" + dbConn.Params
	}
	db, err := sql.Open("mysql", connectionString)

	fmt.Println("Connecting to MySQL database: " + connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
