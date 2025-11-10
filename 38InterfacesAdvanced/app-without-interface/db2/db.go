package db2

import (
	"database/sql"
	"fmt"
)

const (
	dbDriver = "postgres"
)

type Postgres struct {
	db *sql.DB
}

func New(dbUser, dbPassword, dbHost, dbPort, dbName string) (*Postgres, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		fmt.Printf("postgres connection failure: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("postgres ping failure: %v", err)
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (d *Postgres) CreateUser(userInfo string) error{
	d.db.Exec("Insert...")
	return nil
}

func (d *Postgres) ReadUser(userId string) string{
	d.db.Exec("Select...")
	return fmt.Sprintf("Got user with id: %v",userId)
}

func (d *Postgres) UpdateUser(userId string, userInfo string) error{
	d.db.Exec("Update...")
	return nil
}

func (d *Postgres) DeleteUser(userId string) error{
	d.db.Exec("delete...")
	return nil
}