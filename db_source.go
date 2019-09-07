package dbs

import (
	"database/sql"
	"fmt"
)

type DBSource struct {
	Driver     string `json:"driver"`
	ServerName string `json:"server_name"`
	Port       int    `json:"port"`
	Name       string `json:"name"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

func (dbSource *DBSource) Connection() (*sql.DB, error) {
	db, err := sql.Open(dbSource.Driver, dbSource.Source())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dbSource *DBSource) Source() string {
	if dbSource.Driver == MYSQL {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbSource.User, dbSource.Password, dbSource.ServerName, dbSource.Name)
	}

	if dbSource.Driver == POSTGRES {
		return fmt.Sprintf(
			"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable search_path=public",
			dbSource.ServerName,
			dbSource.User,
			dbSource.Password,
			dbSource.Name,
		)
	}

	if dbSource.Driver == SQLITE3 {
		return dbSource.Name
	}

	return ""
}
