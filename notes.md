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


