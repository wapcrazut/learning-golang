# Golang - Advanced

## Data Types

- string: Max 2^32-1 bytes
- bool: true or false
- int: Max 2^32-1 bytes
- int8: Max 2^7-1 bytes
- int16: Max 2^15-1 bytes
- int32: Max 2^31-1 bytes
- int64: Max 2^63-1 bytes
- uint: Max 2^32-1 bytes
- uint8: Max 2^8-1 bytes
- uint16: Max 2^16-1 bytes
- uint32: Max 2^32-1 bytes
- uint64: Max 2^64-1 bytes
- byte: alias for uint8
- rune: alias for int32
- float32: Max 2^32-1 bytes
- float64: Max 2^64-1 bytes
- complex64: Max 2^64-1 bytes
- complex128: Max 2^128-1 bytes


## Variable Declaration Comparison

| **Aspect**                        | **Short Variable Declaration (`:=`)**                   | **Long Variable Declaration (`var`)**                           |
|-----------------------------------|--------------------------------------------------------|-----------------------------------------------------------------|
| **Conciseness**                   | Concise and reduces redundancy.                        | More verbose, requires both `var` and explicit type.            |
| **Readability**                   | Easier to read, especially for local variables.         | Clear, but may be more cluttered, especially for simple cases.  |
| **Type Specification**            | Infers type from assigned value.                        | Allows explicit type specification for clarity.                |
| **Scope**                         | Limited to the block or function where declared.        | Variables have a broader scope, potentially leading to conflicts.|
| **Initialization Separation**     | Declaration and initialization in a single line.        | Allows separation of declaration and initialization.           |
| **Package-Level Variables**       | Not suitable for package-level variables.               | Appropriate for package-level or global variables.             |
| **Redundancy**                    | Reduces redundancy in simple cases.                     | May introduce redundancy when the type is evident from the value.|
| **Use Cases**                     | Suitable for local variables with short lifetimes.      | Suitable for variables with broader scope or explicit typing needs.|

## Zero Values 

| **Type**                          | **Zero Value**                                            |
|-----------------------------------|-----------------------------------------------------------|
| **Boolean**                       | `false`                                                   |
| **Integer**                       | `0`                                                       |
| **Floating Point**                | `0`                                                       |
| **Complex**                       | `0i`                                                      |
| **String**                        | `""`                                                      |
| **Pointer**                       | `nil`                                                     |
| **Function**                      | `nil`                                                     |
| **Interface**                     | `nil`                                                     |
| **Slice**                         | `nil`                                                     |
| **Channel**                       | `nil`                                                     |
| **Map**                           | `nil`                                                     |

## Pointers

Pointers are used to store memory addresses of other variables, and are used to share a value between different function calls.

Structs are value types, so when you pass them into a function, they are copied. If you want to modify the original struct, you need to pass in a pointer to it.

In the other hand, maps and slices are reference types, so when you pass them into a function, they are not copied. If you want to modify the original map or slice, you don't need to pass in a pointer to it.

### Value and Reference Types

Value types are types that store their own data, while reference types are types that store a reference to the underlying data.

- Values types: Use pointers to change these things in a function.
- Reference types: Don't worry about pointers with these.

| **Value Types**                   | **Reference Types**                                       |
|-----------------------------------|-----------------------------------------------------------|
| **int**                           | **slices**                                                |
| **float**                         | **maps**                                                  |
| **string**                        | **channels**                                              |
| **bool**                          | **pointers**                                              |
| **structs**                       | **functions**                                             |

## Maps and Structs Comparison

Maps and structs are used to represent a collection of properties.

| **Maps**                          | **Structs**                                               |
|-----------------------------------|-----------------------------------------------------------|
| **All keys must be the same type**| **Values can be of different type**                       |
| **All values must be the same type**| **Useful for grouping data together**                    |
| **Keys are indexed**              | **Keys don't support indexing**                           |
| **Don't need to know all the keys at compile time**| **Need to know all the different fields at compile time**|
| **Reference type**                | **Value type**                                            |
| **Use to represent a collection of related properties**| **Use to represent a "thing" with a lot of different properties**|

## Concrete Types and Interfaces Types

Concrete types are types that you can create a value directly from.

Examples of concrete types: `string`, `int`, `float`, `bool`, `struct`, `map`, `slice`, `array`.

Interfaces are a type of type that you can't create a value directly from.

Examples of interfaces: `Bot`, `Reader`, `Writer`, `http.ResponseWriter`, `http.Request`.

## Interfaces

Interfaces can be used to define a set of methods that a type must implement.

At the same time, interfaces can implement other interfaces. Meaning, if a type implements the main interface it needs to satisfy all the methods of the other interfaces that it implements to be considered as a valid implementation.

In Go, interfaces are implicit, meaning that you don't need to explicitly say that a type implements an interface, it is enough to implement the methods of the interface to be considered as a valid implementation.

For example:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

### Reader Interface

This interface is defined in the `io` package, and it is used to read data from a source. 

Implemented by: `os.File`, `bytes.Buffer`, `strings.Reader`, `bufio.Reader`, `net.TCPConn`, `net.Conn`, `net.UnixConn`, `net.IPConn`, `net.UDPConn`, `net.UnixPacketConn`, `net.PacketConn`...

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Writer Interface

This interface is defined in the `io` package, and it is used to write data to a destination.

Implemented by: `os.File`, `bytes.Buffer`, `strings.Builder`, `bufio.Writer`, `bufio.ReadWriter`, `net.TCPConn`, `net.Conn`, `net.UnixConn`, `net.IPConn`, `net.UDPConn`, `net.UnixPacketConn`, `net.PacketConn`...

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### io.Copy

This function is defined in the `io` package, and it is used to copy data from a source to a destination.

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

## Go Routines

Go routines are used to run functions concurrently by creating a new rutine with the `go` keyword.

The main rutine is the one that runs when you execute a program, and it is the one that creates the other go routines, and it will not wait for the other go routines to finish, so you need to use channels to communicate between them.  

Also, you need to consider that it does not guarantee that the function will run concurrently, but it gives the Go scheduler the option to run it concurrently.

General recommendations for go routines:

- Don't create go routines in libraries.
- When you create a go routine, know how it will end.
- Concurrency != Parallelism.

### Channels

Channels are data types used to communicate between go routines. Meaning, they are used to send and receive data between go routines.

### Concurrency vs Parallelism

Concurrency is about dealing with lots of things at once, while parallelism is about doing lots of things at once.

By default, Go programs run on a single core, so they are concurrent but not parallel.

See the following table:

| **Concurrency**                   | **Parallelism**                                           |
|-----------------------------------|-----------------------------------------------------------|
| **Single-core CPU**               | **Multi-core CPU**                                        |
| **One thread**                    | **Multiple threads**                                      |
| **Asynchronous**                  | **Synchronous**                                           |
| **Concurrency is about dealing with lots of things at once**| **Parallelism is about doing lots of things at once**|
| **Concurrency is about structure**| **Parallelism is about execution**                        |

### Sleep Function

This function is defined in the `time` package, and it is used to pause the execution of a go routine for a certain amount of time.

```go
func main() {
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("Hello")
    }()
    fmt.Println("World") // This will be printed first
}
```

### Mutex

Mutex is a data type defined in the `sync` package, and it is used to synchronize access to shared data between go routines.

```go
type Counter struct {
    value int
    mutex sync.Mutex 
}

func (c *Counter) Increment() {
    c.mutex.Lock()
    defer c.mutex.Unlock() // This will be executed after the function finishes
    c.value++
}

func (c *Counter) Value() int {
    return c.value
}

func main() {
    counter := Counter{}
    for i := 0; i < 1000; i++ { // This will be executed sequentially
        go func() {
            counter.Increment() // This will be executed concurrently
        }()
    }
    time.Sleep(1 * time.Second)
    fmt.Println(counter.Value())
}
```

### Deferring Functions

This feature is used to execute a function after the current function finishes.

```go
func main() {
    defer fmt.Println("World") // This will be executed after the function finishes
    fmt.Println("Hello")
}
```