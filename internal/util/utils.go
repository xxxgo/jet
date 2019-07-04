package util

import (
	"github.com/go-jet/jet/internal/3rdparty/snaker"
	"strings"
)

func ToGoIdentifier(databaseIdentifier string) string {
	if len(databaseIdentifier) == 0 {
		return databaseIdentifier
	}
	databaseIdentifier = strings.Replace(databaseIdentifier, " ", "_", -1)
	databaseIdentifier = strings.Replace(databaseIdentifier, "-", "_", -1)

	return snaker.SnakeToCamel(databaseIdentifier)
}

func ToGoFileName(databaseIdentifier string) string {
	databaseIdentifier = strings.Replace(databaseIdentifier, " ", "_", -1)
	databaseIdentifier = strings.Replace(databaseIdentifier, "-", "_", -1)

	return strings.ToLower(databaseIdentifier)
}
