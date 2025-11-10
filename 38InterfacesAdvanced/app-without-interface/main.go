package main

// If we want to use db1 in our code

import "advanced-interfaces-app-without-interface/db1"

type UserApplication struct {
	db *db1.MySql
}

func newUserApplication(db *db1.MySql) *UserApplication{
	return &UserApplication{db:db}
}

func main() {
	db,_ := db1.New("dummy-user","dummy-password","dummy-host","dummy-PORT","dummy-db")
	newApp := newUserApplication(db)
	newApp.db.CreateUser("dummyUser")
	newApp.db.ReadUser("dummyUserID")
	newApp.db.UpdateUser("dummyUserID","dummyUser")
	newApp.db.DeleteUser("dummyUserID")
}

// Now, if we want to use db2 in our code, we need to do the following changes
/*
import "advanced-interfaces-app-without-interface/db2"

type UserApplication struct{
	db *db2.Postgres
}

func newUserApplication(db *db2.Postgres) *UserApplication{
	return &UserApplication{db: db}
}

func main(){
	db, _ := db2.New("dummy-user","dummy-password","dummy-host","dummy-PORT","dummy-db")
	newApp := newUserApplication(db)
	newApp.db.CreateUser("dummyUser")
	newApp.db.ReadUser("dummyUserID")
	newApp.db.UpdateUser("dummyUserID","dummyUser")
	newApp.db.DeleteUser("dummyUserID")
}
*/


// Here, above we can see that our Application is tightly coupled with the database we are using. 
// If we have to change the database, we need to make a lot of changes in the code
// This can be avoided by using Interfaces