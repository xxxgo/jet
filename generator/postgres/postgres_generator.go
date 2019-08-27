package postgres

import (
	"database/sql"
	"fmt"
	"github.com/xxxgo/jet/v2/generator/internal/metadata"
	"github.com/xxxgo/jet/v2/generator/internal/template"
	"github.com/xxxgo/jet/v2/internal/utils"
	"github.com/xxxgo/jet/v2/postgres"
	"path"
	"strconv"
)

// DBConnection contains postgres connection details
type DBConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	SslMode  string
	Params   string

	DBName     string
	SchemaName string
}

// Generate generates jet files at destination dir from database connection details
func Generate(destDir string, dbConn DBConnection) error {

	db, err := openConnection(dbConn)
	defer utils.DBClose(db)

	if err != nil {
		return err
	}

	fmt.Println("Retrieving schema information...")
	schemaInfo, err := metadata.GetSchemaInfo(db, dbConn.SchemaName, &postgresQuerySet{})

	if err != nil {
		return err
	}

	genPath := path.Join(destDir, dbConn.DBName, dbConn.SchemaName)

	err = template.GenerateFiles(genPath, schemaInfo.TableInfos, schemaInfo.EnumInfos, postgres.Dialect)

	if err != nil {
		return err
	}

	return nil
}

func openConnection(dbConn DBConnection) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s %s",
		dbConn.Host, strconv.Itoa(dbConn.Port), dbConn.User, dbConn.Password, dbConn.DBName, dbConn.SslMode, dbConn.Params)

	fmt.Println("Connecting to postgres database: " + connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
