# Golang - Advanced

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

### Value and Reference Types Table

| **Value Types**                   | **Reference Types**                                       |
|-----------------------------------|-----------------------------------------------------------|
| **int**                           | **slices**                                                |
| **float**                         | **maps**                                                  |
| **string**                        | **channels**                                              |
| **bool**                          | **pointers**                                              |
| **structs**                       | **functions**                                             |

- Values types => Use pointers to change these things in a function.
- Reference types => Don't worry about pointers with these.

## Data Types

- string
- bool
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- byte
- rune
- float32
- float64
- complex64
- complex128

