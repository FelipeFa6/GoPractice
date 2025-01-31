package db

import (
    "database/sql"
    "log"
    "dbQuery/config"
    _ "github.com/lib/pq"
)

func InitConnection() (*sql.DB, error) {
    connectionString, err := config.GetDBConnectionString()
    if err != nil {
        return nil, err
    }

    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    return db, nil
}

