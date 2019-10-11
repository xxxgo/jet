package mysql

import (
	"database/sql"
	"github.com/xxxgo/jet/v2/generator/internal/metadata"
	"github.com/xxxgo/jet/v2/internal/utils"
	"strings"
)

// mySqlQuerySet is dialect query set for MySQL
type mySqlQuerySet struct{}

func (m *mySqlQuerySet) ListOfTablesQuery() string {
	return `
SELECT table_name
FROM INFORMATION_SCHEMA.tables
WHERE table_schema = ? and table_type = ?;
`
}

func (m *mySqlQuerySet) PrimaryKeysQuery() string {
	return `
SELECT k.column_name
FROM information_schema.table_constraints t
JOIN information_schema.key_column_usage k
USING(constraint_name,table_schema,table_name)
WHERE t.constraint_type='PRIMARY KEY'
  AND t.table_schema= ?
  AND t.table_name= ?;
`
}

func (m *mySqlQuerySet) ListOfColumnsQuery() string {
	return `
SELECT COLUMN_NAME, 
	IS_NULLABLE, IF(COLUMN_TYPE = 'tinyint(1)', 'boolean', DATA_TYPE), 
	IF(DATA_TYPE = 'enum',  CONCAT(TABLE_NAME, '_', COLUMN_NAME), ''), 
	COLUMN_TYPE LIKE '%unsigned%'
FROM information_schema.columns 
WHERE table_schema = ? and table_name = ?
ORDER BY ordinal_position;
`
}

func (m *mySqlQuerySet) ListOfEnumsQuery() string {
	return `
SELECT (CASE c.DATA_TYPE WHEN 'enum' then CONCAT(c.TABLE_NAME, '_', c.COLUMN_NAME) ELSE '' END ), SUBSTRING(c.COLUMN_TYPE,5)
FROM information_schema.columns as c
	INNER JOIN information_schema.tables  as t on (t.table_schema = c.table_schema AND t.table_name = c.table_name)
WHERE c.table_schema = ? AND DATA_TYPE = 'enum';
`
}

func (m *mySqlQuerySet) GetEnumsMetaData(db *sql.DB, schemaName string) []metadata.MetaData {

	rows, err := db.Query(m.ListOfEnumsQuery(), schemaName)
	utils.PanicOnError(err)
	defer rows.Close()

	ret := []metadata.MetaData{}

	for rows.Next() {
		var enumName string
		var enumValues string
		err = rows.Scan(&enumName, &enumValues)
		utils.PanicOnError(err)

		enumValues = strings.Replace(enumValues[1:len(enumValues)-1], "'", "", -1)

		ret = append(ret, metadata.EnumMetaData{
			EnumName: enumName,
			Values:   strings.Split(enumValues, ","),
		})
	}

	err = rows.Err()
	utils.PanicOnError(err)

	return ret

}
