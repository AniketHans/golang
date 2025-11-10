# Advanced Golang Concepts

## Closures

1. Suppose we have the following code

   ```go
    package main

    import "fmt"
    func activate100RsGiftCard() func(int) int {
        amount := 100
        debitFunc := func(debitAmount int) int {
            amount -= debitAmount
            return amount
        }
        return debitFunc
    }

    func main() {
        dF := activate100RsGiftCard()
        fmt.Println("Bought chocolate for 25 Rs, the giftcard amout left is", dF(25)) // amount left is 75
        fmt.Println("Bought chips for 15 Rs, the giftcard amout left is", dF(15)) // amount left is 60
        fmt.Println("Bought toffees for 10 Rs, the giftcard amout left is", dF(10)) // amount left is 50
    }
   ```

2. Here above, the `activate100RsGiftCard()` returns a func that takes an integer value as parameter and returns an integer value. The function uses a local variable of the activate100RsGiftCard() and perform operations on it.
3. Now, after the code line `dF := activate100RsGiftCard()`, ideally, the variable `amount` should gets created and destroyed as activate100RsGiftCard() returned the debtFunc and exited from the callstack. The scope of the amount variable is over after activate100RsGiftCard() returned
4. But, in further calls, the debitAmount sent to the debitFunc is still being debited from the amount variable although `amount` variable's scope should be over
5. This is the concept of closures, as the `debitFunc()` is still dependent on the `amount` variable that is defined outside of itself.
6. **The golang will enclose both the function and any variables that it depends on, defined/declared outside the function, and store it as an enclosed unit known as closure**. The function can access the outside variables and perform operations on it whenever required.
   As should below:
   ```
    // In memory
    // as closure
    amount = 100
    debitFunc := func(debitAmount int) int {
            amount -= debitAmount
            return amount
        }
   ```

## Generics

1. Suppose you have slices of different number (int and float) types and you need to write functions to get the sum of all the values in them:
   ```go
        numbers1 := []int{1, 2, 3, 4, 5}
        numbers2 := []int32{1, 2, 3, 4, 5}
        numbers3 := []int64{1, 2, 3, 4, 5}
        numbers4 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
        numbers5 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
   ```
2. Ideally, we should write separate functions to handle the slices of each type which increases the code duplicacy
3. We can define a type Number which can expect any int and float
   ```go
        type Number interface {
            int | int32 | int64 | float32 | float64
        }
   ```
4. Now, we can define a function `NumberSum` which can expect any type defined in Number by using Generics syntax.
5. The generic syntax:- `func <name>[<type-name> <type>](<params>){}`
   For example: `func NumberSum[T Number](values []T) T`. Here Number can be anything amoung `int | int32 | int64 | float32 | float64` and T is the type Number. So if Number is int64 then the T will also be int64
   It also means `T` can accept anytype defined in Number
6. Thus, we can write a common function using generics as follows

   ```go
        func NumberSum[T Number](values []T) T {
            var s T;
            for _, val := range values {
                s += val
            }
            return s
        }

   ```

## Interfaces

1. Suppose you have the main application which needs to connect to some database.
2. We need to create database package that will have some functions needed by the main application to perform the database tasks.
3. Suppose you have user info in main application and your want to perform CRUD operations in database using the user info. For this, you will need the following functions:
   1. createUser(userInfo)
   2. readUser(userId)
   3. updateUser(userId, userInfo)
   4. deleteUser(userId)
4. For any database package to be used by our main application, we need to implement the above functions in it.
5. Also, we might change the database in future if needed. For example, we might start using NoSQL database.
6. If we pass the database object of a specific type in the main application and in future we have to change to the database, then we have to make changes to the main application code as well. As we can see here: [App using database without intefaces](./38InterfacesAdvanced/app-without-interface/main.go)
7. We can avoid the above behaviour by using Interfaces.
8. In our application, instead of creating new application by passing an db object, we can pass an interface to it. Any db package which implements the interface is eligible for crreating out application object. As we have implemented here: [App using database with interface](./38InterfacesAdvanced/app-with-interface/main.go)
