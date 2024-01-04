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