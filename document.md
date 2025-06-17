# Documentation for Go learning

# Variable

Variables in Go are used to store data that can be referenced and manipulated throughout your program. They must be declared before use and can be of various types such as int, float64, string, and more.

##### Example

Here is an example of variable declaration in **Go**:

```go

import "fmt"

func main() {

// variable type

van myNum int = 10

//constant type

const x int = 10

//dynamic type

y:= 15

fmt.Println(myNum)

}

```

# Condition

Conditions in Go are used to execute code based on whether a certain condition is true or false. They can be implemented using the `if`, `else if`, and `else` statements, allowing for decision-making in your programs.

##### Example

```go

if x > 0 {

fmt.Println("x is positive")

}

else if x < 0 {

fmt.Println("x is negative")

}

else {

fmt.Println("x is zero")

}

```

# [Loop](https://go.dev/doc/effective_go#for "visit the official Go documentation for more information on loops")

Loops in Go are used to execute a block of code repeatedly for a specified number of times or until a certain condition is met. They can be implemented using the `for` statements only.

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
 // short declaration
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}

// with range clause
for key, value := range oldMap {
    newMap[key] = value
}
```

# Function

<p style='text-align: justify;'> In Go, functions are defined using the `func` keyword, followed by the function name, parameters in parentheses, and an optional return type. They allow you to encapsulate logic and can be called multiple times with different arguments to perform tasks efficiently.Functions in Go are blocks of code that perform a specific task and can be reused throughout your program. They are defined using the `func` keyword followed by the function name, parameters, and return type.</p>

##### Example

1. Normal single value return function

```go

func functionName (param type) return type{

// function body goes here

return result

}

```

2. Function with multiple return values

```go

func getValues(param type) (int, string) {

//something happen here

return 10, "example"

}

```

# Scope in Go

Global scope, local scope, function scope are same like js scope. but the different is in package scope.

# Package scope:

The directory name is the package name. a single package name arn't same in same directory. first create a create a directory with the same name like package name.

# Custom package

when you want to use a custom package, you need to create a directory with the same name like package name. and then you can use it in your main package. After that you should run a command to compile the package.

```sh

go mod init example.com

```

the above command will create a go.mod file in your directory. Package export function is name should be start with `Uppercase` letter.

## Variable shadowing:

<p style='text-align: justify;'> Variable shadowing is when a variable in an outer scope has the same name as a variable in an inner scope. The inner scope variable hides the outer scope variable. This can be useful for creating a new variable with the same name as an outer scope variable, but it can also lead to unexpected behavior if not used carefully.</p>

##### Example:

```go
    package main
    import "fmt"
    var x = 10

    func main() {
        age := 25
        if age>18{
            x := 20
            fmt.Println(x) // prints 20
        }
        fmt.Println(x) // prints 10
    }
```

# Go Function type:

### Standard Function:

Standard function is a function that is defined with the `func` keyword. It can take arguments and return values. The function name is followed by a list of arguments in parentheses, and the return values are specified after the `return` keyword.

##### Example:

```go
    func functionName (arg1 type1, arg2 type2) return _type {
        return
    }
```

### Init function:

<p style='text-align: justify;'> Init function is a special function that is called when the package is initialized. The function name should be `init` and we can't call this function. It is used to initialization the package. The init function is called before the main function. The init function is not called explicitly, it is called automatically by the Go runtime.</p>

```
func init() {
    // initialization code
    fmt.Println("Initialization code")
}
```

### Anonymous and IIFE (Immediately Invoked Function Expression) Function:

<p style='text-align: justify;'>Anonymous function is a function that is defined without a name. It can be assigned to a variable or passed as an argument to another function. IIFE is a function that is defined and called immediately. It is used to create a new scope and avoid polluting the global namespace.</p>

```go
    // anonymous function
    func() {
        fmt.Println("Anonymous function")
    }
    // IIFE
    func() {
        fmt.Println("IIFE")
    }() // IIFE
```

# Go First order function, Higher order function:

1. `First order function` is a function that can assign in a variable. like:

- Standard Function
- Init Function
- Anonymous Function
- IIFE Function
- Function expression

2. `Higher order function` is a function that can pass to another function as an argument or return a function from another function.

- Function that takes another function as an argument
- Function that returns another function
- Function that takes another function as an argument and returns another function both.

# Go Internal Memory:

## Code segment:

- All `function` are store in the code segment.
- <p style='text-align: justify;'>  The code segment is the part of the program that is stored in memory. It is the part of the program that is executed by the CPU. The code segment is read-only. It is stored in the program's memory space. The code segment is loaded into memory when the program is started. It is stored in the program's memory space until the program is terminated. The code segment is divided into several parts, including the program's entry point, the program's data , and the program's stack. </p>

## Data segment:

- All `global` variables are store in the data segment.
- <p style='text-align: justify;'>The data segment is the part of the program that is stored in memory. It is the part of the program that is used to store data. The data segment is read-write. It is stored in the program's memory space. The data segment is loaded into memory when the program is started. It is stored in the program's memory space until the program is terminated. The data segment is divided into several parts, including the program's global variables, the program's static variables, and the program's constants.</p>

## Stack:

- <p style='text-align: justify;'> The stack is a region of memory that is used to store data that is used by the program . It is a LIFO (Last In First Out) data structure. The stack is used to store data that is used by the program, such as function call stack, local variables, and parameters . The stack is divided into several parts, including the program's stack frame, the program's stack pointer, and the program's stack limit.</p>

#### Stack frame:

    when a function call, a new stack frame is created. The stack frame contains the function's local variables, parameters, and return address.

## Heap:

- <p style='text-align: justify;'>The heap is a region of memory that is used to store data that is dynamically allocated by the program. It is a LIFO (Last In First Out) data structure. The heap is used to store data that is dynamically allocated by the program, such as objects, arrays, and strings. The heap is divided into several parts, including the program's heap block, the program's heap pointer, and the program's heap limit.</p>

## GC (Garbage Collector):

- <p style='text-align: justify;'> The garbage collector is a program that is used to automatically manage the memory of the program. It is used to free up memory that is no longer needed by the program. The garbage collector is used to prevent memory leaks and to improve the performance of the program. The garbage collector is divided into several parts, including the program's garbage collector, the program's garbage collector thread, and the program's garbage collector queue. </p>

## Closure:

- Closure is a function that has access to its own scope, the parent scope, and the global scope.

# Struct:

- A struct is a collection of variables of different data types that can be used together as a single unit. It is a user-defined data type that can be used to represent complex data structures. A struct can have fields of different data types, such as integers, floats, strings, and others. A struct can also have methods, which are functions that operate on the struct's fields. A struct can be used to represent a real-world object, such as a person, a car, or a house. A struct can be used to represent a complex data structure, such as a graph or a tree.

##### Example:

```go
    type Person struct {
        name string
        age int
    }

    func main(){
        user := Person{
            name: "John",
            age: 30,
        }
    }

```

# Receiver Function:

- A receiver function is a function that is used to operate on a struct. It is takes a pointer to the struct as an argument. The receiver function is used to access and modify the fields of the struct. Used to implement methods for the struct. This function is also used to perform operations on the struct, such as printing its fields or modifying its fields.

##### Example:

```go
type Person struct {
        name string
        age int
    }
    func (user Person) printInfo() {
        fmt.Printf("Name: %s, Age: %d\n", user.name, user.age )
    }
    func main(){
        user := Person{
            name: "John",
            age: 30,
        }
        user.printInfo()
    }
```

# Array

- An array is a collection of elements of the same data type stored in contiguous memory locations. The elements of an array are accessed using an index, which is a number that identifies the position of an element in the array. Arrays are homogeneous, meaning that all elements must be of the same data type. Arrays are mutable, meaning that their elements can be modified after the array is created.

##### Example:

```go

    var arrayName [array length] type of array element

    var myArray [5] int; // example

    // another type of array
    myArray2 := [5]int{1, 2, 3, 4 , 5}

```

# Pointer:

- A pointer is a variable that holds the memory address of another variable. Pointers are used to access and modify the value of a variable indirectly. Pointers are declared using the asterisk symbol (\ \*) before the pointer name. Pointers are used to implement dynamic memory allocation, where memory is allocated at runtime.

- to print address of memory cell user `&` ampersand symbol for the variable name.
- to print value of memory cell user `*` asterisk symbol for the pointer name.

##### Example:

```go
    var x int = 10
    var ptr = &x
    fmt.Println(ptr) // prints address of x

    fmt.Println(*ptr)// print value of ptr
```

# Slice:

- A slice is a dynamically-sized, flexible view into the elements of an array. Slices are referenced values that refer to an underlying array. Slices are used to represent a contiguous sequence of elements of the same data type. Slices are mutable, meaning that their elements can be modified after the slice is created.

**_Type of slice_** :

1. Conventional Array to slice:

```go
   var myArray = [3] string {"hello", "go", "slice"}
   s:= myArray[1:2] //myArray[startIndex: endIndex]
```

2. Slice to slice:

```go
    s1 := []int{1, 2, 3}
    s2 := s1[1:2]
```

3. Slice literal type:

```go
    s := []int{1, 2, 3} // array literal which is length are not specify.
    s[3]=4
```

4. Make function to create slice:

```go
s := make([]int, 3, 5) // make(type, length, capacity)
```

5. Empty or nil slice: `as the previous example variable type was dynamic but here it is static`

```go
    var s []int
    fmt.Println(s)

    // output []
```

# Variadic function:

- A variadic function is a function that can take a variable number of arguments. The arguments are passed as a slice. The variadic function can be called with any number of arguments, including zero.

##### Example of variadic function:

```go
package main
import "fmt"
// variadic function
func sum(numbers ...int) {
    sum := 0
    for _, num := range numbers {
            sum += num
        }
    fmt.Println(sum)
    }
func main() {
    sum(1, 2, 3, 4, 5)
    sum()
}
```

# Context switching and PCB(Process Control Block):

- Context switching is the process of switching between different processes or threads. It involves saving the current state of a process and restoring the state of the next process to be executed. The process control block (PCB) is a data structure that contains information about a process, such as its program counter, registers, and memory management information. The PCB is used to manage the context switching process with current state of process. When context switching occur then current state of process is saved in PCB and next process state is loaded in PCB.

# Concurrency:

- Go has built-in concurrency features that make it easy to write concurrent programs. The `go` keyword is used to start a new goroutine, which is a lightweight thread. Goroutines are scheduled by the Go runtime, which makes it easy to write concurrent programs without worrying about the underlying threading details.

<h3 style="color: red">what is concurrency ?</h3>

Concurrency is the ability of a program to execute multiple tasks simultaneously. This can improve the responsiveness and throughput of a program. Go's concurrency features make it easy to write concurrent programs.

# Concurrency and Parallelism:

- Concurrency and parallelism are related but distinct concepts. Concurrency refers to the ability of a program to execute multiple tasks simultaneously, while parallelism refers to the execution of multiple tasks at the same time on multiple processors. Go's concurrency features are designed to make it easy to write concurrent programs, but they do not automatically parallelize tasks. To achieve parallelism, you need to use Go's built-in parallelism features , such as the `goroutine` and `channel` types.

# Thread:

# Go Datatype:

Go datatype:

- Int
- Float
- String
- Boolean

## Numerical Datatype:

### Singed Integer

To print decimal number in signed integer we use`"%d"`

| Type  | Bit | Byte | Note                              |
| ----- | --- | ---- | --------------------------------- |
| int   | -   | 4    | Depend on the system architecture |
| int8  | 8   | 1    | Store 8-bit signed integer        |
| int16 | 16  | 2    | Store 16-bit signed integer       |
| int32 | 32  | 4    | Store 32-bit signed integer       |
| int64 | 64  | 8    | Store 64-bit signed integer       |

### Unsigned Integer

Unsigned integer is a number that is always positive..Go has following unsigned integer types:
| Type | Bit | Byte | Note |
| ------ | --- | ---- | --------------------------------- |
| uint | - | 4 | Depend on the system architecture |
| uint8 | 8 | 1 | Store 8-bit unsigned integer |
| uint16 | 16 | 2 | Store 16-bit unsigned integer |
| uint32 | 32 | 4 | Store 32-bit unsigned integer |
| uint64 | 64 | 8 | Store 64-bit unsigned integer |

### Float

To print floating point number we use`"%f"` and specify the precision using `%.xf` where x is the number of digits after the decimal point.
| Type | Bit | Byte | Note |
| ------- | --- | ---- | ---------------------------------- |
| float32 | 32 | 4 | Store 32-bit floating point number |
| float64 | 64 | 8 | Store 64-bit floating point number |

### Byte

Alias for unsigned `unit8` type
| Type | Bit | Byte | Note |
| --- | --- | --- | ---------------------------------- |
| byte | 8 | 1 | Store 8-bit unsigned integer |

## Rune

Rune is a `int32`-bit `Unicode` code point. It is used to represent a single `Unicode` character. Print flag is `"%c"`.

## Boolean Datatype

To print boolean value we use `"%v"`. It prints `true` or `false`.
| Type | Bit | Byte | Note |
| ---- | --- | ---- | ------------------- |
| bool | 1 | 8 | Boolean type are store as 8-bit integer |
