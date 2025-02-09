> These contains notes of sections after 29

### Go commands

1. `go run`, this command compiles code in one or more file and immediately executes them.
2. `go build`, this command code in one or more files. It creates an executable file that can be run whenever needed.
3. `go install`, compiles and installs a package.
4. `go get`, downloads the raw source code of someone else's package.

### Go Packages

1. Package == Project
2. A package is a collection of common source code files.
3. Each go file belongs to a package.
4. The files serving common objective or files belong to a particular project should belong to same package.
5. There are 2 types of packages in Go:
   1. Executable type, these, when compiled, generates a file that we can run.
   2. Reusable type, these are used as helpers to other projects/packages. These are like code dependencies or libraries.
6. The name of the package defines whether the package is an executable type or reusable type.
7. The word `main` is used to make an executable type package, `package main`. **`package main` ---> go build ---> main.exe**. When the main.exe is executed, the main() present in the function will automatically gets executed.
8. Any other work, except `main`, on building, will not give an executable file.**`package <anything>` ---> go build ---> nothing!!**.
9. Note, the files that are under package main must have a main() inside them. It means, out of all the files under `package main`, only one file should have main().

### `import <package>`

1. The importing of package is done so that our current package can have access to the code inside the imported package.
2. Some packages come as part of standard library of Go like "fmt". You can read about them on `golang.org/pkg`

### Declaring Variables

1. Go is a statically typed language, means we need to predefine the type of variable
2. Method 1 :- `var <variable-name> <variable-type> = <value>`, eg. `var name string = "Aniket Hans"`.
3. Method 2 :- `<variable-name> := <value>`, eg `name := "Aniket Hans"`
4. The `:=` will only be used at the time of declaring the variable, not while re-assigning value to it. E.g `name := "Aniket Hans"; name = "AH"`
5. You can't declare global variables using `:=`, you have to use `var` for declaring global variables.
6. If you dont use any declared variable in your code, then Go will throw error.

### Function Declarations

1. Syntax:- `func <function-name>() <return-type> {}`. E.g. `func getName() string{}`

### Arrays

1. Fixed length list of records
2. Every element in a array should be of same type.

### Slice

1. Slice is an array that can grow or shrink.
2. Every element in a slice should be of same type.
3. Syntax:- `<variable> := []<type of slice>{data1,data2}`, eg. cards := []string{"ace of spade", "king of hearts"}
   `var <variable> = []<type of slice>{data1,data2}`, eg. var cards = []string{"ace of spade", "king of hearts"}
4. Adding new element to slice, `<slice-variable> = append(<slice-variable>, <new value>)`, e.g. cards = append(cards,"Queen of Diamonds")
5. `append()` does not modify the existing slice, instead it returns a new slice which is assigned back to the variable.

### For loop

1. Looping over a slice, `for <index variable>, <value> := range <slice-variable> {}`

### OOPs in Go

1. Go is not an Object Oriented Programming language as there are no classes in Go.
2. In classes, we have properties (variable) in classes and methods (functions inside class), which use the properties to perform operations and give some result.
3. In Go, we first create a custom type and then make a receiver function that will receives the object/variable of custom type to perform some operations and give result.
4. The custom type can be a string, []string, struct etc.

### Defining type

1. Syntax: `type <type-name> <type>` e.g making a type of []string, `type name []string`. Here, name will have all the properties of []string.
2. Both of these are similar:-
   1. humans := []string{"AH","DH"}
   2. humans := name{"AH","DH"}
3. We can also define type on primitive data types. Eg. `type customString string`. We can define customType variables in 2 types:
   1. Using var, `var s1 customString = "hello"`
   2. Using :=, `s2 := customString("hi")`

### Receiver function

1. Syntax:- `func (<variable> <type-of-variable>) <function-name>() <return-type>{}`
2. E.g. for type name, `func (n name) printNames() {}`
3. The receiver function printNames here will have access to the `n` variable which is of type []string and can perform oerations on it.
4. The receiver function is called using the object/variable of type `name`, eg. humans.printNames()
5. We can consider, loosely, the variable defined using `type` as class and properties and the receiver function as class methods which can be invoked by objects of type defined and the methods have access to the properties of calling object.
6. ```go
    type name []string // Its like declaring class name with one property of type []string
   func (n name) printNames(){ // Its like defining the method of class name
        for i, name := range n{
        fmt.Println(i,name)
        }
   }

   humans := name{"AH","DH"} // Its like initializing object humans of name class.
   humans.printNames() // class objects can be used to call the class methods
   ```

7. Note, we have created name of type []string and attached printNames() method to it. You might be thinking, can we create an object of type []string and then try to call printNames(), like :-
   ```go
   humans2 := []string{"AH","DH"}
   humans2.printNames() //This will throw an error.
   ```
   name has kind of inherited the properties of []string and we have enhanced it with printNames() method. The printNames() method is available for object of type name and not []string.
8. Thus, name and []string are diffrent.
9. The receiver, say `n`, can be of pointer type as well.

   ```go
   func (n *name) func1(){} // It is similar to call by reference, any changes made in func1 on n will be refrlected to n1.
   func (n name) func2(){} // It is similar to call by value, any changes made in func2 on n will not be reflected to n2.

   n1.func1()
   n2.func2()
   ```

### Slices

1. Slices are zero indexed.
2. Slice range syntax:- `<slice-name> [startIncludingIndex : uptoNotIncluding]`
3. Eg:
   ```go
   var s = []int32{10,20,30,40}
   s[0:2] // [10,20]
   ```
4. We can omit either of `startIncludingIndex` or `uptoNotIncluding`.
   ```go
   var s = []int32{10,20,30,40}
   s[0:2] // [10,20]
   s[:3] // [10,20,30]
   s[1:] // [20,30,40]
   ```

### Returning multiple values

1. A function can return multiple values.
2. Syntax:
   ```go
   func func1(param1 string, param2 int) (int,string){
      return param2*10, param1+"Hi"
   }
   ```
   - Here, we are returning 2 values of type int and string respectively from the function func1.
3. Note: functions that are defined in the same package can be used in any of the files without any import, given the files also belong to the same package.
4. When the function called, the number of variables, that will be holding the retunred values from the called function, will be equal to the number of returning values from the function.
   ```go
   // If a function, say func1() (string,int){}, returns 2 values
   ret1 := func1() // This will give Error as only one variable is collecting the values returned by the func1()
   ret2, ret3 := func1() // This is alright as the returning 2 values are captured in 2 variables respectively.
   ```
