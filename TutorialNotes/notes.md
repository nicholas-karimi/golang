### Intro
`Go` is not an _object oriented programming language_. It does not support classes and inheritance.

### Formatting Strings in Go

From the standard library
1. `fmt.Printf` - Prints a formatted string to standard output.
2. `fmt.Sprint` - Prints a formatted string

#### Examples
**%v** - Interpolate the default representation
The `%v` variant prints the Go syntax representation of a value.

```Go
fmt.Printf("I am %v years old", 27)
// I am 27 years old

fmt.Printf("I am %v years old", "wqay too many")
```

**%s** - Interpolate a string
`fmt.Printf("I am %s years old", "way too many")` 

**%d** - Interpolate an Integer in Decimal form
`fmt.Printf("I am %d years old", 27)` 

**%f** - Interpolate a Decimal
```Go
    fmt.Printf("I am %f years old", 27.99)
    // I am 27 years
    // Round to 2 decimal places
    fmt.Printf("I am %.2f years old", 27.990)
    // I am 27.99 years old
```
#### Print Type insteaf of a value
```Go
age := 27
fmt.Printf("His => age %T\n", age)
```

```Go
// Use Sprintf to store the value instead of printing to standard output
package main
import "fmt"

func main() {
    const name = "Nicholas Karimi"
    const openRate = 100.5

    msg := fmt.Sprintf("Hi %s, your rate is %.1f percent", name, openRate)


    fmt.Println(msg)
}
```

#### Passing Variables by Value
In Go, variables are passed by value not by reference.
  i.e when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the callers data.

  ```Go
  fun main(){
    x := 5
    increment(x)

    fmt.Println(x)
    <!-- prints 5 because the increment func received a copy of x -->
  }

    func increment(x int) {
        x++
    }
```

#### Ignoring Return Values
A function can return a value that the cller doesnt care about. 
Go does not allow you to have unused variables. We use `_` to explicitly ignore variables

```Go
func getPoint() (x int, y int) {
    return 3, 4
}

// ignore y value
x, _ := getPoint
```

#### Named Return Values
if return values are given names, they are trated the same as if they were new variables defined at the top of the function.
> a return statement qithout argument return the named return value. Also known as `naked return`. Naked return should only be used only in short functions.

```Go
func getCoords()(x, y int){
    // x and y are intialized with zero values

    return // automatically returns x and y
}
// same as 
func getCoords()(int, int){
    var x int
    var y int
    return x, y
}
```

#### Early Returns (Guard clauses)
ability to return early from a function.
Guard clauses leverage the ability to return early from a function (or `continue` thro a loop) to make nested conditions one dimensional.
Instead of using `if/else` chains, we just return early from the function at the end of each conditional block.

```Go
func divide(divinded, divisor int)(int, error) {
    if divisor == 0{
        return 0, errors.New("Can't divide by zero")
    }
    return divinded/divisor, nil
}

```

### STRUCTS
Structs are used to rep structed data in a key value pair manner
```Go
type carMan Struct {
    Make string
    Model string
    Year int
	Height int
	Width int
    FrontWheel Wheel
    BackWheel Wheel

}
// \Nestef
type Wheel Struct {
    Radius int
    Material string
}

// insantiate the struct
// myCar = carMan{}
car := carMan{
    Make: 'make',
    Model: 'model',
    Year: 'year',
    Height: 'height',
    Width: 'width',
}

myCar.FrontWheel.Radius = 5

```

### Anonymous Struct
its defined without a name nd threfore cannot be referenced elsewhere.
To create anonymous struct, just instantiate immediately using a second pait of brackets after declaring the type

```Go
myCar := struct {
    Make string
    Model string
}{
    Make: "tesla",
    Model: "model 3",
}
```

### Embeded vs Nested Struct
An embedded struct's fields are accessed at the top level, unlike nested struct.
Promoted fields can be accessed like normal fields except they cant be used in composite literals.

```go
type car Struct{
    make string
    model string
}

type truck Struct{
    // car is embedded so the definition of a "truck" now additonally contains all of the fields of car struct.
    car
    bedSize int
}

// instantiate
lanesTruck := truck{
    bedSize: 10,
    car: car{
        make: "Toyota",
        model: "GTI",
    },
}

fmt.Println(lanesTruck.bedSize) 

// embeded fields promoted to top-level - instead of lanesTtuck.car.make
fmt.Println(lanesTruck.make)
fmt.Prinln(lanesTruck.model)
```

#### STRUCT Methods in GO
methods are functions that have a receiver
a receiver is a special parameter that syntactically goes before the name of the method.

```go
type rect struct {
    width int
    height int
}
// area has a receiver of (r rect)

func (r rect) area() int {
    return r.width * r.height
}

r := rect{
    widht: 5,
    height: 10,
}

fmt.Println(r.area())

```

## INTERFACES
Interfaces are a collection of method signatures. A type implemens an interface if it has all methods given interface defined on it..
_an interface is an abstract type that represents other types_
eg. a shape interface that must return its area and perimeter of the shape

```go
type shape struct{
    area() float64
    perimeter() float64
}

type rect struct{
    widht, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct{
    radius float64
}
func (c circle) area() float64 {
    return math.PI * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    rerurn 2*math.PI * c.radius 
}
```

#### Interface Implementation
interfaces are implemented implicitly
A type never declared that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type autmatically fulfils that inteface

#### Type Assertion in GO
every onec in a while you'll need to access the underlying type of an interface value. You can cast an interface to its underlying type using type assertion.

```go
type shape interface{
    area () float64
}

type circle struct{
    radius float64
}

// c is a new circle cast from s which is an instance of a shape. "ok" is a bool that is 
// true if s was a circle or false if s isn't a circle.

c, ok := s.(circle)

```
#### Type switches
type switch makes it easy to do several types assertions in a series.
A type switch is similar to a regular switch statement, but the cases specifify the types instead of values.

syntax
```go
func printNumericValues(num interface{}){
    switch v := num.(type){
        case int:
            fmt.Printf("%T\n", v);
        case string:
            fmt.Printf("%T\n", v);
        default:
            fmt.Printf("%T\n", v);
    }
}

func main(){
    printNumericValues(1) // int
    printNumericValues("1") // string
    printNumericValues(struct{}{}) // struct {}


}

```
#### Clean Interfaces
- rule of thumb
1. Keep Interfaces Small
   interfaces are meant to define the minimal beaviour necessary to accurately rep the idea or concept.
2. Interfaces should have no knowledge of satisfying types
   An interface should define what is necessary for other types to classify as  memebers of that interface. They shouldnt be aware of any types that happen to satisfy the interface at design time.
3. Interfaces are NOT Classes
   they're slimmer
   dont have constructores and deconstructors that require data is creatd or destroyable.

## ERRORS
go expresses erorrs with `error` values. An error is any type that implements the simple built in error-interface.

```go
type error interface{
    Error() string
}
```
Error is either `nil` or `!nil`. If it is `nil`, it means that everything is fine else if error is `!nil` something went wrong.

ASCII to Interger conversion
```go
// Atoi converts  a stringfied number to integer 
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert", err)
    // because 42b is not a valid integer, we print:
    // couldnt convert: strconv.Atoi: Parsing "42b" invalid syntax error
    return
}
```
- an `errr` in a `nullable` string representing what went wrong or nothing.
  
#### Custom Error Interface 
- error interface has only one single method that needs to ne defined returning a string `Error`.

```go
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account: %v", e.name)
}
```
#### The Errors Package
Go std library provides an `errors` package `errors.New()` for dealing with errors.
`var err error = errors.New("something went wrong")`



## Loops 
Loops in `go` can omit sections fo the loop. eg. the `CONDITION` (middle part) can be omitted qhich causes the loop to run forever.
```go
for INTIAL; ; AFTER;{
    // do something
}
```

##### While loop
- there is no `while loop` in go. A while loop is similar to a _for_ loop only that is does not have the **intial** and **after** statement. It just runs until some condition is nolonger true.
- Go allows omission of sections of a for loop.
- A while lloop is just  a for loop with only one condition.
```go 
for CONDITIONS {
    // DO SOMETHING WHILE CONDITION IS TRUE
}
```
__eg__
```go
plantHegight := 1
for plantHegight < 5{
    fmt.Println("still growing; current height: " + plantHegight)
    plantHegight++
}
fmt.Println("plant has grown to ",  plantHegight, " inches")
```

### Logical Operators
#### AND
`true &&  false` `// false`
`true && true` `// true`

#### OR
`true || false` `// true`
`false || false` `//false`

#### Continue
The `continue` keyword stops the current iteration of a loop and continues the next iteration.
Its a powerful way to use the `guard clause` pattern within loops.
It helps bail out early.

```go
for i := 0; i < 10; i++ {
    if i == 0 {
        continue
    }
    fmt.Println(i)
}
```

#### Break
The `break` keyword stops the current iteration of a loop and exits.
```go
for i := 0; i < 10; i++ {
    if i == 5{
        break
    }
    fmt.Println(i)
}
```

### Arrays
I go `arrays` have a fixed sze.
The type `[n]T` is an array of n valuesof type `T`.

`var myInts [10] int`

use intialixe literals if you know what goes where in an array.
`primes := [6] int{1, 2, 3, 4, 5, 6}`

### Slices
A `slice` is a dynamically sized flexible view into an array. They wrap arrays.
create a slice from an array

`myArray := [6] int{1,2,3,4,5,6}`

create a slice from an array from index 1 to index 4

`mySlice := myArray[1:5]`

slices are built ontop of array for memory management reasons.
slices hold references to underlying array, and if you assign one slice to another both refer to same array.
slices are stored in contigous memory

#### Syntax
``func make([]T, len, cap) []T`
`mySlice := make([]int, 5, 10)`
capacity argument can be omited and slice defaults to the length of the underlying array.
`s=make([]int,10)` // initalizes 0 values 
Slices have built-in functions `len` and `cap` tht gets the length of the slice and the capacity of the slice respectively.

##### Panic
**panicking** means that `runtime` ERROR that is unrecoverable.
`panicking` refers to a situation
where the normal flow of execution is stopped, and if it’s within a function, the function returns to its caller.

#### Variadic
a variadic function receives a variadic argument as a slice.
function in the standard library can take arbitrary nmber of final arguments. this is accomplished using the `"..."` syntax in the function signature.
```go
func sum(nums ...int) int {
    // nums is a slice 
    for i := 0; i < len(nums); i++ {
        nums[i] = num
    }
}

func main() {
    total := sum(1,2,3,4)
    fmt.Println(total) //10
}
```

#### Spread Operator
allows us to pass a slice into a `variadic` function. It consists of 3 dots following the sluce in the function call.
```go
func printStrings(strings ...string) {
    for i := 0; i < len(strings); i++ {
        fmt.Println(strings[i])
    }
}

func main() {
    names := [] string{"bob", "nic", "joy"}
    printStrings(names...)
}
```

#### Slice Append
built in function is a variadic function.
syntax is
`func append(slice []Type, elems, ...Type) [] Type`
`mySLice = append(mySLice, thingOne)`
`mySLice = append(mySLice, thingOne, thing2)`
`mySLice = append(mySLice, anotherSlice...)`

`append` changes the underlying array of its parameters and returns a new slice. Using append on anything other than its self is a `BAD` idea.

print slice memory addrss use `&slice[0]`

_dont do on append_
`someslice = append(otherSlice, elelemt)`


### RANGE
`Go` provides syntatic suggar to iterate over elements of a slice.
```go
for INDEX, ELEMENT := range SLICE {

}
```
```go
fruits := []string {"apple", "orange", "banana", "mango"}
for i, fruit := range fruits{
    fmt.Println(i, fruit)
}
```

### MAPS
`maps` are similar to js objects, python dicts and ruby hashes.
They're data structures that provide `key-value` mapping.
The `zero` value of a map is `nil`
Use the `make` function to create a map.

`ages := make(map[string]int)`
`ages['John'] = 22`
```go
ages = map[string]int{
    "John": 22,
    "Mary": 102,
}
```
`len` on maps returns total numbe rof key-value pairs.
`len(ages)`

maps are more efficient than slices.

#### Map mutations
1. Insert an element
`m[key] = elem`

2. Get and element
`elem = m[key]`
3. Delete an element
`delete(m, key)`
4. Check if a key exists
`elem, ok := m[key]`
if key is in m, ok is true else false.
if `key` is not in the map, then elem is zero value for the maps element type.

#### Key Types
What makes a type qualify to be used a map key?
Any type can be used as a value in the map, but keys are more _restrictive_
__maps keys maybe any type that is comparable__ what is absent from the list is the `slices,maps and funcs` as they cannot be compared using the `==` ops.

#### Nested Maps
Maps can contain, maps creating nested structures.
`map[string]map[string]int`
`map[rune]map[string]int`
`map[string]map[string]map[string]int`

in go we rep individual characters as `runes` instead of strings.


## First class and high order functions.
a `high-order` function allows us to pass around funcs as other values.
a function that accepts another func or returns another dunc as a values is a `first-class` function.

`firstclass` is a function that can be trated like any other value. A function type is dependent on the types of the parameters and return values.
eg `func() int`| `func(string) int`

Use cases
`HTTP API` handlers
`Pub/Sub` handlers
`Onclick` callback handlers

*High order functions* is a function that takes a function as an argument or returns a function as a return value.

#### Function Currying
practice of writing functions that takes a function as input and returns a new function
best usecase for `middleware functions`

#### Defer Keyword
a unique feature in Go. It allows a function to be executed automatically just before its closing function returns.
the defer call argument are evaluated immediately, but the function call is not executed until the sorrounding function returns.


### Closures
A closure is a function that references a variable fro outside its own function body.

in the example below the `concatter` function  returns a function that has references to enclosed `doc` value. Each successive call to `harryPotterAggregator` mutates the same `doc` variable.

```go
func concatter() func (string) string{
    doc := ""
    retrun func(word string) string {
        doc += word + ""
        return doc
    }
}

fun main( ){
    harryPotterAggregator := concatter()
    harryPotterAggregator("Mr.")
    harryPotterAggregator("and")
    harryPotterAggregator("Mrs.")
    harryPotterAggregator("Dursley")
    harryPotterAggregator("of")
    harryPotterAggregator("number")
    harryPotterAggregator("four,")
    harryPotterAggregator("Privet")

    
    fmt.Println(harryPotterAggregator("Drive"))


}

```

### POINTERS
`pointers` are all about how we store memory in our computers.
reference  a pointer value using `&variable`. 
`eg. x := 5 to create a z reference z address ``z := &x`.
update value of `x` without having to access original value use the _de-reference ops_ `*z=6`

a variable is a named location in memory that stores a value.

#### A pointer is a Variable
a pointer is a variable that stores the memory address of another variable. that is, a pointer `points to` the location of where data is stored NOT the actual data itself.

_syntax_
`var p *int`

The `&` operator generates  a pointer to its operand.
 `mySTring := "hello"`
 `mySTringPtr = &mySTring`

 Pointers allow us to manipulate data in memory directly without making copies or duplicating data.

 The `* dereference` operator to gain access to the value,
 `fmt.Println(*mySTringPtr)` // read mystring through the pointer
 `*mySTringPtr = "world"` // write mystring through the pointer

 #### NIL Pointers
 if a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (`panic`)
 crashing the program.
 When using pointers check if its `nil` before dereferencing.

 #### POINTER Recievers
 methods with pointer receivers can modify the value to which the receiver points.

 ```go
 type car struct {
    color string
 }
 func (c *car) setColor(color string){
    c.color = color
 }

 func main(){
    c := car{
        color: "black",

    }

    c.setColor("blue")
    fmt.println(c.color) // blue
 }
 ```
 pointer receivers are more common than value receivers.


 ## Local Development
 ### PACKAGES
 every go program is made of packages.
 a package named "main" has the entey point at the `main()` function and is compiled into an executable program.
 a package with any other name other than "main" is a `library package` and have no entry point.


#### Naming Conventions
a package name is the same sa the last element of its import path. eg. the `math/rand` package comprises files that begin with `package rand`.
 packages live at the directory level.

 ### MODULES
 go programs are organized into packages. A package is a directory of Go code that's all compiled together.
 A `module` is a releasele collection of `go` packages.
 _a go repository contains only one module, located at the root of the repository_
 a file named `go.mod` at the root of the project declares a module. It contains:
 1. the module path
 2. the version of go language the project requires
 3. optional external dependencies the project has.

the go.mod
```go
module github.com/botdotdev/exampleproject
go 1.21
require github.com/botdotdev/examplepackage v1.3.0
```

>In go ecosystem, there's no central location for 3rd party packages like `npmjs.com` in javascript. Instead the Go toolchain works on top of git and uses the import path as rhe remote url where go can download code.
>modeule github.com/ or gitlab.com

- an `import path` is a module path + package subdirectory.

#### Initialize a go module
`go mod init github.com/nicholas-karimi/hellogo`
where `github.com/nicholas-karimi/ is the remote repository for the package and hellogo is the folder where the package is located.

### Run GO
`go help run`
to 

1. `go run pacjage.go` 
used to quickly compile and run the Go package.
the compliled binary is not saved in your working directory

2. `go build`
used to compile production executables.

#### Build an executable
`go build`

run the new program
`./buildname`

Compile and run
`go build && ./buildname`


3. `go install`
compiles and installs the program locally.


running `go build` on a non-main package  does not output an `executable` program but _a compiled package silently saved to the local build cache_

go dependency management is heavily based on git and remote url.

4. `go get`
downloads and installs the remote package.
`go get github.com/nicholas-karimi/go-tinytime`

### Clean Packages

#### Rule of thumb

1. Hide Internal logics
ref pillars of OOP - encapsulation.

2. Do not Change APIs

3. Do not export functions from the main package
A `main` package isnt a library, theres no need to export functions from it.

4. Packages shouldnt know about the dependencies


## Chap 13: CHANNELS AND CONCURRENCY

#### CONCURRENCY
CONCURRENCY is the ability to perform multiple tasks at the same time.
Typically, code is executed one line at a time, one after another. This is known as `sequential or synchronous` execution.

Syntax
We use the `go` keyword when calling a function.
`go doSomething()`
the go keyword will spwawn a new `goroutine`.
when used, we're unable to capture the return value of the function.

#### CHANNELS
Channels are typed, thread safe. It allows different go routines to communicate.
> since go routines does not support returning values, channels are used to re-synchronize the code.
>
SYntax
`ch := make(chan int)`

##### Send Data to a channel
`ch <- 69`
The `<-` is called the `channel operator`. Data flows in the direction of the arrow. This operator will block until another go routine is ready to receive a value.

##### Receive Data from a channel
`v := <-ch`
this reads and removes a value from the channel and saves it in the variable v.

Empty structs are ofted used as `token` in Go programs. A token is a `unary` value i.e we dont care what is passed through the channel. We care when and if it was passed.

We can block and wait unitl something is sent on the channel using the syntax of `<-ch`. This will block until it pops a single item off the channel, then continue to discard it.

#### BUFFERRED CHANNELS
Channels can be optionally buffered.
buffered channels helps stored information in them.
A buffer of a length that allows senders to send things to it until its if full. When the receiver is available to pop them, it will read the 1 by 1 and pop them.
`ch := make(chan int, 100)`
sending on a buffered channle only blocks until the buffer is full.
receiving blocks only when the buffer is empty

#### CLOSING CHANNELS IN GO
We close the channel to indicate that were done with the channel.
A channel should only be closed from the `sending` side.

SYntax:
`ch := make(chan int) //do something close(ch)`

#### Check if channel is closed
use the `ok` value similar to accessing `maps`. receivers can check the ok value when receiving from a channel to test if the channel is closed
`v, ok := <-ch`
ok is false if the channel is closed.
if channel is buffered, ok is true until the channel buffer is drained.

#### Dont send on a closed channel
sending on a closed channel will cause panic. A panic on main goroutine will cause entire program crash.

#### Range keyword in Channel
channels similar to maps and slices, can be ranged.
```go
for item := ch {
    // item is the next value received from the channel
}
```

#### CHANNEL SELECT
A `select` statement is used to listen to multiple channels at the same time. its similar to `switch` but for channels.
```go
select {
    case i, ok := <- chInts:
    fmt.Println(i)
    case s, ok := <- chStrings:
    fmt.Println(s)

    default:
     // receiving from channel will block
}
```

#### STd lib channels
1. `time.Tick()` -a std libe that returns a channel that sends a value on a given interval.
2. `time.After()` - sends a value once after duration has passed
3. `time.Sleep()` - blocks the current  go routine for a given duration

#### Read Only Channels

```go
func main() {
    ch := make(chan int)
    readCh(ch)
}

func readCh(ch <-chan int) {
    // ch can only be read in this function
}
```

#### Write Only Channels
Same for the writers
- writes int the channel.
```go
func writeCh(ch <- chan int){
    // ch can only be written in this channel.
}
```

## Chap 14: MUTEXES
**Mutexes** allow users to lock access to data. This ensures that we can control which goroutines can access certain data at what time.
Go std lib provides built-in mutex implementations using `sync.Mutex` type. 
Mutex methods are:
1. `.Lock()`
2. `.Unlock()`

```go
protected(){
    mux.Lock();
    defer mux.Unlock();
    // the rest of the function is protected
    // any other call to mux.Lock() will blocked
}
```
Mutex == mutal exclusion

#### MAPS ARE NOT THREAD SAFE
maps are not safe for concurrent use. If you have multiple go routines accessing the same map, and atleast one of them is writing the map, you must lock the map in mutex.

> Race condition - when to diff goroutines racing to get access to specific resource.
> mutal exclusion - because a mutex excludes different threads(or go routines) from accessing the same data same time.

### RW MUTEX
read/write mutex
has additional methods to sync.RWMutex()
1. Lock()
2. Unlock()

the sync.RWMutex() has
1. RLock()
2. RUnlock()

## Chap 15: GENERICS IN GO
generics allow is to use variables to refer to specific types.

```go
func splitAnySlice[T any] (s []T)([]T, []T){
    mid := len(s)/2
    return s[:mid] + s[mid:]
}

fistInts, secondInts := splitAnySlice([]int{1, 2, 3})
fmt.Printfln(fistInts, secondInt)
```
generics help us to achieve the `DRY` principle.

Why Generics
1. Generics reduce repititive code.
2. Used in libraries and packages more often

### CONSTRAINTS

#### INTERFACE TYPE LIST
a new wat of writing interfaces introudced when generics were released.

###  PARAMETER CONSTRAINTS



#### NAMING GENERIC TYPES
capital `T` is commonly used when there is only a single type param for a given function.

## Chap 16: GO Proverbs
by Robert Pike
1. Do not communicate by sharing memory, share memory by communicating.
2. Concurrency is not parrallelism
3. Channels orchestrates, mutexes serialize
4. The bigger the interface, the weaker the abstraction
5. Make the zero useful value
6. Interface{} says nothing 
7. Gofmts style is no ones favorite, yet gofmt is everyones favorite.
8. A little copying is better than alittle dependency
9. Syscall must always be guarded with build tags
10. Cgo must always be guarded with build tags
11. Cygo is not Go
12. With the unface package theres no gurantee
13. Clear is better than clever
14. Reflection is never clear
15. Errors are values
16. Dont just check errors, handle them gracefully
17. Design the architecture, name the components, document the details
18. Documentation is for users
19. Dont panic.
