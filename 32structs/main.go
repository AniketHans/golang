package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       uint64
}

// ************** EMBEDDED STRUCT ***************
type LoginInfo struct{
	username string
	password string
}

type UserData struct{
	name string
	age uint
	loginInfo LoginInfo

}

type UserData2 struct{
	name string
	age int
	LoginInfo // Here a field of type and name LoginInfo will be created.
}

// **********************************************

func (u UserData) Print(){
	fmt.Printf("\n%+v\n",u)
}

func main() {
	person1 := Person{"a","b",1}
	fmt.Println(person1) // {a b 1}
	person2 := Person{firstname: "c", age: 4, lastname: "d"}
	fmt.Println(person2) // {c d 4}
	var person3 Person
	fmt.Println(person3) // {  0}, basically default values "" "" 0
	fmt.Printf("\n%+v",person3) //{firstname: lastname: age:0}

	person3.firstname="e"
	person3.lastname="d"
	person3.age=5
	fmt.Printf("\n%+v\n",person3) // {firstname:e lastname:d age:5

	var person4 Person =Person{
		firstname: "m",
		lastname: "n",
		age:10,
	}
	fmt.Println(person4) //{m n 10}

	//***** Using embedded structs *********

	user1 := UserData{
		name: "user1",
		age: 25,
		loginInfo: LoginInfo{
			username: "user_1",
			password: "user1pass",
		},
	}
	fmt.Printf("\n%+v\n",user1) // {name:user1 age:25 loginInfo:{username:user_1 password:user1pass}}

	user2 := UserData2{
		name: "user1",
		age: 25,
		LoginInfo: LoginInfo{
			username: "user_1",
			password: "user1pass",
		},
	}
	fmt.Printf("\n%+v\n",user2) // {name:user1 age:25 LoginInfo:{username:user_1 password:user1pass}}

	// calling the reciever function
	user1.Print() // {name:user1 age:25 loginInfo:{username:user_1 password:user1pass}}
}


// type User struct{
//       name string
//       age uint
//    }

//    func updateName(u *User){
//       fmt.Println("Old name value",u.name)
//       u.name = "haha"
//       fmt.Println("New name value",u.name)
//    }

//    func (u *User) updateAge(){
//       fmt.Println("Old age value",u.age)
//       u.age = 90
//       fmt.Println("New age value",u.age)
//    }

//    func main(){
//    user := User{
//       name: "user1",
//       age: 25,
//    }

//    fmt.Println("Old name before passing",user.name)
//    updateName(&user)
//    fmt.Println("New name after passing",user.name)

//    fmt.Println("Old age before passing",user.age)
//    user.updateAge()
//    fmt.Println("New age after passing",user.age)
// }


// //Output
// /*
// 	Old name before passing user1
// 	Old name value user1
// 	New name value haha
// 	New name after passing haha
// 	Old age before passing 25
// 	Old age value 25
// 	New age value 90
// 	New age after passing 90
// */