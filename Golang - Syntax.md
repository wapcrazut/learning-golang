# Golang - Syntax

## Arrays and Slices

Arrays have a fixed size, slices are dynamic.

```go
// Array
var fruitArr [2]string

// Assign values
fruitArr[0] = "Apple"
fruitArr[1] = "Orange"

// Declare and assign
fruitArr := [2]string{"Apple", "Orange"}

// Slice
fruitSlice := []string{"Apple", "Orange", "Grape"}
```

## Variable Declaration

```go
var name string = "John Doe"
var age int = 30
var isCool = true
```

## Shorthand

```go
name := "John Doe"
age := 30
isCool := true
```

## Function Declaration

```go
func add(num1 int, num2 int) int {
    return num1 + num2
}
```

## Function Declaration (Shorthand)

```go
func add(num1, num2) int {
    return num1 + num2
}
```

## Literal Function

```go
// Without parameters
func() {
    fmt.Println("Hello")
}()

// With parameters
func(name string) {
    fmt.Println("Hello", name)
}("John")
```


## Loop Declaration

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

## Receivers Declaration

```go
type person string

func (p Person) sayHello() {
    fmt.Println("Hello", p)
}
```

## Conditionals

```go
x := 5
y := 10

// If else
if x < y {
    fmt.Printf("%d is less than %d\n", x, y)
} else {
    fmt.Printf("%d is less than %d\n", y, x)
}
```

## Structs

Structs are similar to classes in other languages but without inheritance.

```go
// Define person struct
type person struct {
    firstName string
    lastName string
} 

p := person{firstName: "John", lastName: "Doe"}
``` 

### Alternative

```go
p := person{"John", "Doe"}
```

### Alternative

```go
var p person
p.firstName = "John"
p.lastName = "Doe"
```

## Pointers

Pointers allow you to point to the memory address of a value and are used to share a value between different function calls.

`&` operator gets the memory address of a value.
`*` operator gets the value of a memory address.

```go
a := 5
b := &a

fmt.Println(a, b) // 5 0xc0000b4008
fmt.Printf("%T\n", b) // *int

// Use * to read val from address
fmt.Println(*b) // 5
fmt.Println(*&a) // 5

// Change val with pointer
*b = 10
fmt.Println(a) // 10
```

### Pointer Receiver

```go
type person struct {
    firstName string
    lastName string
}

func (p *person) updateName(newFirstName string) {
    (*p).firstName = newFirstName
}

func main() {
    p := person{"John", "Doe"}
    p.updateName("Jane")
    fmt.Println(p) // {Jane Doe}
}
```

## Maps

Maps are key-value pairs and are similar to objects in JavaScript.

Keys must be of the same type, values must be of the same type.

```go
// Define map
colors := map[string]string{
    "red":  "#ff0000",
    "blue": "#0000ff",
}

// Delete from map
delete(colors, "red")
```

### Alternative

```go
// Define map
colors := make(map[string]string)

// Assign key values
colors["white"] = "#ffffff"
```

### Alternative

```go
// Define map
var colors map[string]string

// Assign key values
colors["white"] = "#ffffff"
```

## Interfaces

Interfaces are collections of method signatures that an object can implement.

As general rules:

- If a type has a function with the same signature as the interface, it automatically implements that interface.
- We don't need to explicitly say that it implements the interface.

```go
type bot interface {
    getGreeting() string
}

type englishBot struct{}

func (englishBot) getGreeting() string {
    return "Hello"
}
```

## Go Routines

Go routines are used to run functions concurrently.

```go
func main() {
    go print("Hello")
    print("World")
}

func print(text string) {
    fmt.Println(text)
}
```

### Alternative

```go
func main() {
    go func(text string) {
        fmt.Println(text)
    }("Hello")
    fmt.Println("World")
}
```

## Channels

Channels are used to communicate between go routines.

`<-` operator is used to send and receive values to and from a channel, the direction of the arrow indicates the direction of the data flow.

The close function is used to close a channel, it is optional but it is a good practice to close channels when they are not used anymore.

```go
func main() {
    c := make(chan string) // Create channel
    go print("Hello", c)
    fmt.Println(<-c) // Receive value from channel
}

func print(text string, c chan string) {
    c <- text // Send value to channel
    close(c) // Close channel (optional)
}
```

