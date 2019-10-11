package postgres

import (
	"database/sql"
	"github.com/xxxgo/jet/generator/internal/metadata"
	"github.com/xxxgo/jet/internal/utils"
)

// postgresQuerySet is dialect query set for PostgreSQL
type postgresQuerySet struct{}

func (p *postgresQuerySet) ListOfTablesQuery() string {
	return `
SELECT table_name 
FROM information_schema.tables
where table_schema = $1 and table_type = $2;
`
}

func (p *postgresQuerySet) PrimaryKeysQuery() string {
	return `
SELECT c.column_name
FROM information_schema.key_column_usage AS c
LEFT JOIN information_schema.table_constraints AS t
ON t.constraint_name = c.constraint_name
WHERE t.table_schema = $1 AND t.table_name = $2 AND t.constraint_type = 'PRIMARY KEY';
`
}

func (p *postgresQuerySet) ListOfColumnsQuery() string {
	return `
SELECT column_name, is_nullable, data_type, udt_name, FALSE
FROM information_schema.columns
where table_schema = $1 and table_name = $2
order by ordinal_position;`
}

func (p *postgresQuerySet) ListOfEnumsQuery() string {
	return `
SELECT t.typname,  
       e.enumlabel
FROM pg_catalog.pg_type t 
   JOIN pg_catalog.pg_enum e on t.oid = e.enumtypid  
   JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace
WHERE n.nspname = $1
ORDER BY n.nspname, t.typname, e.enumsortorder;`
}

func (p *postgresQuerySet) GetEnumsMetaData(db *sql.DB, schemaName string) []metadata.MetaData {
	rows, err := db.Query(p.ListOfEnumsQuery(), schemaName)
	utils.PanicOnError(err)
	defer rows.Close()

	enumsInfosMap := map[string][]string{}
	for rows.Next() {
		var enumName string
		var enumValue string
		err = rows.Scan(&enumName, &enumValue)
		utils.PanicOnError(err)

		enumValues := enumsInfosMap[enumName]

		enumValues = append(enumValues, enumValue)

		enumsInfosMap[enumName] = enumValues
	}

	err = rows.Err()
	utils.PanicOnError(err)

	ret := []metadata.MetaData{}

	for enumName, enumValues := range enumsInfosMap {
		ret = append(ret, metadata.EnumMetaData{
			EnumName: enumName,
			Values:   enumValues,
		})
	}

	return ret
}
