package db

import "fmt"

var (
	psqlConnStr = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		psql.username,
		psql.password,
		psql.dbName,
		psql.host,
		psql.port,
	)
)
