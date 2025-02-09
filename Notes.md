# Important points:

### Golang:

1. Golang is a compiled language. Golang compiles everything and the runtime is baked into the final product. So as long as you have `go` installed you dont need any other thing to run the go code.
2. Go executables are different for different OS.
3. We dont have classes in Golang.
4. There are no try catch in this language and you even dont need that.

### Golang Coding:

1. The main() function acts as entry point for golang code.
2. You dont have to manually add or remove packages. Go automatically do this for you
3. `go env` will give you info about all the env variables set in golang.

### Lexer in Golang

1. The job of the lexer, in any programming language, is to check if you are using the proper grammer of the language while writting code.
2. Lexer does its job at the time of pre-compilation of the code.

### Types in Golang:

1. types are case sensitive.
2. They sometimes define if you are going to make a thing public or private.
3. If the initial letter of the variable or function is Uppercase then it is **Public** else **Private**.
4. In Golang, the variable should be known in advance as it is a compiled language.
5. Everything is a Type in Golang.
6. Basic Types:
   1. String
   2. Bool
   3. Integer -> uint8 uint64 int8 int64 uintptr
   4. Floating -> float32 float64
   5. Complex
7. Advance Types:
   1. Arrays
   2. Slices
   3. Maps
   4. Structs
   5. Pointers
   6. Functions
   7. Channels
   8. Mutex

### Variables:

1. In go, if you have created a variable and have not used it. You will get an error.
2. float64 has more precision than float32
3. If you dont initialize the variable, go automatically initializes it with default value.
4. If you have not defined a variable type but have initialized it then Lexer will automatically put the type based on the value assigned.
5. We also can declare and define variables using syntax,
   - `<variable-name> := value`
   - This is used when the type of the variable is not known previously
6. You cannot define a global variable using `:=`

### Comma ok Syntax || Error ok syntax:

1. In golang, we dont have try catch. Go expects the problems and solutions to be treated as variables.
2. `_` can be used to ignore any value

### Builds for Windows, Mac and Linux

1. Run `go env` and find for `GOOS`.
2. GOOS can be used to build for different OS.
3. `go build` can be used to build the code. It will automatically search for `main.go`.
4. If you want to build for a different OS, run `GOOS=<OS> go build`.This command may not work on windows as we can't put env variables at the start of the command. Here GOOS value can be 'darwin`, 'linux', 'windows' etc

### Memory Management

1. Memory allocation and deallocation happens automatically in Golang
2. Usually 2 methods are used: new() and make(). Both these methods helps us to allocate some memory and in that memory we can put up some data structure.
3. new():
   - Memory is allocated but not initialized
   - You will get a memory address for reference
   - It gives zeroed storage (Here you cannot put any data initially)
4. make():
   - Memory is allocated and initialized as well.
   - You will get a memory address for reference
   - It gives non-zeroed storage (Here you can put any data initially)
5. Garbage collection happens automatically.
6. The **GOGC env** variable set the initial threshold for Garbage collector to trigger. After the memory allocation threshold is met, the garbage collector get triggered for collection.
7. Out of scope and nil values get garbage collected.
8. `runtime` package is important.

### Arrays

1. Arrays are very less used data structures in Golang.
2. In arrays, you always have to mention the size of it. Eg, `var fruits [10]string`
3. Due to the reserved memory in arrays, the `len(array)` will always be the defined length. It does not rely the actual number of values present in the array.
4. The array is initially initialized with the default value of the declared type.
5. Not many operations are present on the arrays.

### Slices

1. Slices are basically arrays under the hood. These arrays once they get an abstraction layer and some more features, they are called as slices.
2. Here, we can add as many values as we need.

### Maps

1. Maps are key value pairs in golang

### Structs

1. There is no inheritance in Golang. Also, we dont have the concept of super/parent

### Loops

1. There is only for loop in golang, which can be used as while loop as well

### Functions

1. You are not allowed to write functions inside the functions

### Defer

1. `defer` keyword makes a statement, in a function, to not execute at the moment but at the last before the function is returned.
2. Defered statements are executed in LIFO manner.

### Handling request in Golang

1. `net/http` is the package that is mostly used in handling
2. Whenever you make any request to the website, you will get a Response object back.
3. Note: ReadResponse or Response.Write never closes the connection. Means, Whenever you make a request, neither your reader or writer is closing the request.
4. It is your reponsibilty to close that response.

### Go mod

1.  Go mod is one of the tools given by Golang.
2.  `go mod init <package-name>` generaly, package name should be like `github.com/aniket/<name>` which gives info about what is the source and who is managing the package.
3.  go mod file will be having info about the parties involved in running the code and also the go version it is dependent upon.
4.  go mod is basically the Golang's Module dependency system.
5.  If you want to install new dependency, we use `go get` command. `go get -u github.com/<package-name>`. This `go get` will go to the internet and bring all the assets from the web. It is important as some applications hosted on some server might not have access to internet.
6.  `go get` command will add a new line like `require github.com/<package-name> <version>` stating the module used and its version. If you see `//indirect` at the end of `<package-name>`, it means we are not using the module or it is not being directly used.
7.  go.sum file. This file will be created after `go get` or `go mod sum`. It contains info about `github.com/<package-name> <version> h1:<hash>`. As we are pulling lot of things from github, we need some check to avoid pullijg malicious code.
8.  At the time of writing code, we have specified version with a hash for the github code. If someone maliciously pushed some code in the same version, then the hash will be changed thus leading to indication of some changes.
9.  `go mod verify` will be used to check these mismatches in hashes.
10. Even if you have done the go get, you will not be able to see the fetched files/modules in your working directory. To see those, you need to run `go env` and check the GOPATH value.
11. Go to the `<GOPATH>/pkg/mod` path and you will see 'cache' folder init. All the modules that you get from the web will go into this folder and not in your working directory.
12. Next time, Whenever we fetch them, they will be brought from cache and not from the web.
13. `go mod tidy`. This command helps us in managing the packages being used in our code. Suppose you have included some packages but you are not using them in your code, tidy command will remove them. Suppose you are using some package in your code but it got deleted from go.mod file mistakenly then tidy command will bring it back into the mod file. Also, initially for some package, `//indirect` is being shown in mod file but the package is being used in our code then it will remove the `//indirect`.
14. `go list all` will output all the packages being installed.
15. `go list -m all` will output all the packages being used in our code.
16. `go list -m -versions <package-name>` will output all the versions publically available for teh package.
17. `go mod why <package-name>` will output all the packages dependent on `<package-name>`
18. `go mod graph` will out the graph of the dependencies like `<package 1> <package 2>`. Here it means package1 is dependent on package2.
19. `go mod vendor`. This will generate a vendor folder. This vendor folder is like node-modules folder in JS. All the modules being used will be brought in this vendor folder and can be directly used by our code.
20. Even if you have this vendor folder, packages will not be brought from it by default. We need to pass on a flag.
21. `go run main.go` will still be bringing packages from web or cache. `go run -mod=vendor main.go` then it will first check in the vendor folder and if not found then go to the web/cache.
22. Note: All the mod operations are expensive ones, just keep an eye when its ran.

### Golang Folder structure:

1. In root directory of the project, golang expects only one '.go' file. The file name can be anything, mainly main.go.
2. init() is a special function which is going to run at the very first time the entire application is going to execute.

### Context:

1. Whenever you are making API calls to some machines outside your machine, then you need to provide context stating for how long I can make a request, what happens when the request goes off and if the connection is active then there should be a context on which the machines are working on.
2. `context.Background()` is used when you want something that keeps on happening in the background. The connection made with this has no deadline.
3. `context.TODO()` is when you dont have any idea which context to use.

### Concurrency vs Parallelism

1. Concurrency includes context switching resembling parallel execution
2. Parallelism includes executing tasks on multiple CPUs simultaneously.

### Goroutines

1. Goroutines helps in achieving parallelism.
2. Difference between thread and goroutines:
   1. Threads are managed by OS and goroutines are managed by Go runtime.
   2. Threads have fixed stack (1 MB) while goroutines have flexible stack.
3. Goroutines are triggered using `go` keyword. Note: once you fire up a go routine in a function and didn't wait and the main function completes its execution, then you will not see any result from the go routine.
4. Thus, you need to wait for go routine to complete, if you want to see some result before main function completes execution and returns.

### Slogan from Golang

1. Do not communicate by sharing memory; instead share memory by communicating.

### Mutex

1. As the goroutines are managed by go runtime and not the OS, then who will be locking and unlocking the memory address when multiple threads trying to write something.
2. For controlling the locking and unlocking, we need to use mutexes.
3. Mutexes lock the memory address when some go routine is writing to it and we can then unlock it.

### Race Condition

1. When parallel processes try to write data at the same memory address at the same time.
2. `go run --race .` can be used to check if there is any race condition in your code.

### Channels

1. Channels are a way in which multiple go routines can talk to each other. They will still not be aware of whats happening inside the other go routines or how long the other routines will execute.
2. These are used where a goutine needs some information from another go routine. The other go routine will still be running but the info can be used by the goroutine for th execution.
3. Here, the other go routine need not finish its job and then the info will be utilized by the go routine rather this thing can happen on the go while the other go routine is still running.
4. Channels are a way/pipleine through which multiple go routines interact.
5. You can both push and pop values from channel in a function. It will lead to deadlock.
6. Suppose you have the following code in function:

   ```go
   func(){
   myCh := make(chan int) //line1
   myCh <- 5 //line2
   fmt.Println(<-myCh) //line3
   }
    /*
   Channel says you are only allowed to push values to me only if somebody is already listening/consuming it.
   Here above, line2 is pushing value to channnel but we are listening/consuming that value at line3.
   Both the operations are not happening simultaneously.
   */

   //For channels to work, the code should be in following manner,
   func(){
   myCh := make(chan int) //line1
   fmt.Println(<-myCh) //line2
   myCh <- 5 //line3
   }
   /*
   Here, we are listening first, but what we are listening to in channel has no value as the value is comming at line3 here.
   This is kind of Deadlock condition.
   */
   ```

7. Note :- Adding value to a channel and consuming the value should happen at the same time.
8. Also, the number of pushes to the channel should be equal to the number of consumptions else it will lead to error.
9. You can use a buffered channel as well if you have a single listener going through a loop to consume the values.
