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
