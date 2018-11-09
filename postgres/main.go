package postgres

import (
	"fmt"

	"github.com/jackc/pgx"
)

func OpenDB(dbUser string, dbPass string, dbName string, dbHost string) (*pgx.ConnPool, error) {
	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", dbUser, dbName, dbPass, dbHost)

	var err error

	connConfig, err := pgx.ParseConnectionString(connection)

	if err != nil {
		return nil, err
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 5 /* https://wiki.postgresql.org/wiki/Number_Of_Database_Connections#How_to_Find_the_Optimal_Database_Connection_Pool_Size */})

	if err != nil {
		return nil, err
	}

	return pool, nil
}
