package db1

import (
	"database/sql"
	"fmt"
)

const (
	dbDriver = "mysql"
)

type MySql struct {
	db *sql.DB
}

func New(dbUser, dbPassword, dbHost, dbPort, dbName string) (*MySql, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		fmt.Printf("mysqldb connection failure: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("mysqldb ping failure: %v", err)
		return nil, err
	}

	return &MySql{db: db}, nil
}

func (d *MySql) CreateUser(userInfo string) error{
	d.db.Exec("Insert...")
	return nil
}

func (d *MySql) ReadUser(userId string) string{
	d.db.Exec("Select...")
	return fmt.Sprintf("Got user with id: %v",userId)
}

func (d *MySql) UpdateUser(userId string, userInfo string) error{
	d.db.Exec("Update...")
	return nil
}

func (d *MySql) DeleteUser(userId string) error{
	d.db.Exec("delete...")
	return nil
}