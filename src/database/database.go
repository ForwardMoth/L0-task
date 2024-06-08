package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"go-nuts/internal/config"
	"log"
)

const OrdersTableName = "Orders"

type Database struct {
	DB *sql.DB
}

func New(dbCfg *config.DBConfig) Database {
	dialect, connectionURL := dbCfg.Dialect, CreateConnectionURL(dbCfg)
	conn, err := sql.Open(dialect, connectionURL)
	if err != nil {
		log.Fatal(err)
	}
	return Database{DB: conn}
}

func CreateConnectionURL(dbCfg *config.DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.Username, dbCfg.DBName, dbCfg.Password)
}

func (db *Database) Select() ([]json.RawMessage, error) {
	query := fmt.Sprintf("select data from %s", OrdersTableName)
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	res := make([]json.RawMessage, 0)
	for rows.Next() {
		var data json.RawMessage
		if err = rows.Scan(&data); err != nil {
			return nil, err
		}
		res = append(res, data)
	}
	return res, nil
}

func (db *Database) Insert(uuid string, data json.RawMessage) {
	query := fmt.Sprintf("INSERT INTO %s (uuid, data) VALUES($1, $2)", OrdersTableName)
	row := db.DB.QueryRow(query, uuid, data)
	if err := row.Err(); err != nil {
		log.Fatalf("error of creating new order %s", err)
	}
	log.Printf("created new order with uuid %s", uuid)
}
