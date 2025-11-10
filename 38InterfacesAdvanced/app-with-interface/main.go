package main

import "advanced-interfaces-app-with-interface/db2"

type Database interface {
	// Any struct/type, which implements all the below functions, is a member of Database interface
	CreateUser(userInfo string) error
	ReadUser(userId string) string
	UpdateUser(userId string, userInfo string) error
	DeleteUser(userId string) error
}

type UserApplication struct {
	db Database
}

func newUserApplication(db Database) *UserApplication {
	return &UserApplication{db: db}
}

func main() {
	// Code that needs changes in case of DB change now
	//Start
	//For using db1
	// db, _ := db1.New("dummy-user","dummy-password","dummy-host","dummy-PORT","dummy-db")

	//For using db2
	db, _ := db2.New("dummy-user","dummy-password","dummy-host","dummy-PORT","dummy-db")
	//End
	/*************************/

	// Below code remains unchanged and the db is changed
	newApp := newUserApplication(db)
	newApp.db.CreateUser("dummyUser")
	newApp.db.ReadUser("dummyUserID")
	newApp.db.UpdateUser("dummyUserID", "dummyUser")
	newApp.db.DeleteUser("dummyUserID")
}

// Since both db1 and db2 implements the functions needed by Database interface so both becomes the member of Database interface
// Since, any member of the Database interface can be used to create UserApplication object so we can use both db1 and db2