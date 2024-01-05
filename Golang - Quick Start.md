# Golang - Quick Start

Golang is a statically typed, compiled programming language designed at Google.

It does not follow the object oriented paradigm, but it does support some of its features.

For example, it does not have classes, but it does have structs.

The main features of Golang are: 

- Simplicity
- Fast compilation
- Garbage collection
- Built-in concurrency
- Static typing
- Compiled language

## Table of Contents

- [Golang - Quick Start](#golang---quick-start)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Linux](#linux)
  - [Hello World Example](#hello-world-example)
  - [Go CLI](#go-cli)
    - [go build](#go-build)
    - [go run](#go-run)
    - [go fmt](#go-fmt)
    - [go install](#go-install)
    - [go get](#go-get)
    - [go test](#go-test)
  - [Packages](#packages)
    - [Installing packages](#installing-packages)
  - [Organizing code](#organizing-code)
    - [Package declaration](#package-declaration)
    - [Importing packages](#importing-packages)
    - [Syntax](#syntax)
  - [Resources](#resources)

## Installation

### Linux

```bash
# Download the latest version
$ wget https://golang.org/dl/go1.21.4.linux-amd64.tar.gz

# Extract the archive
$ sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz

# Add the following line to ~/.profile
export PATH=$PATH:/usr/local/go/bin

# Reload the profile
$ source ~/.profile

# Verify the installation
$ go version
```

## Hello World Example

```bash
# Create a new project
$ mkdir hello-world && cd hello-world

# Create a new file
$ touch main.go
```

Edit file `main.go` and add the following content:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run the project:

```bash
$ go run main.go
```

## Go CLI

This is list of cli commands that can be used to manage Go projects.

### go build

Build compiles the packages named by the import paths, along with their dependencies, but it does not install the results.

```bash
# Build the project
$ go build

# Build the project and install the results
$ go build -i
```

### go run

Run compiles and runs the file.

```bash
# Run the project
$ go run main.go
```

### go fmt

Formats all go code in the project files.

```bash
# Format the project
$ go fmt
```

### go install

Install compiles and installs the packages named by the import paths, along with their dependencies.

```bash
# Install the project
$ go install
```

### go get

Downloads the source code of the project and its dependencies.

```bash
# Download the project
$ go get github.com/bitpanda-platform/go-template
```

### go test

Test compiles and runs all tests.

```bash
# Run all tests
$ go test ./...
```

## Packages

Package is a collection of source files in the same directory that are compiled together.

`Package == Project == Workspace`

There are two types of packages:

- Executable: Builds a file we can run.
- Reusable: Code used as `helpers`. Good place to put reusable logic.

If the package name is `main`, it will create an executable file. Must have a `func main()`.

If the package name is anything else, it will create a reusable package.

### Installing packages

The official list of packages can be found at [pkg.go.dev](https://pkg.go.dev/std).

To use them in a project, we need import them in the file we want to use them.

```go
import "fmt"
```

Alternatively, we can use the `go get` command to download and install a package, then import it.

```bash
# Install a package
$ go get github.com/bitpanda-platform/go-template

# Install a specific version
$ go get github.com/bitpanda-platform/go-template@v1.0.0
```

## Organizing code

### Package declaration

The first line of every file must be a package declaration.

```go
package main
```

### Importing packages

The second line of every file must be an import declaration.

```go
import "fmt"
```

### Syntax  

- `package main`: Defines a package that can be compiled and run.
- `import "fmt"`: Import a package.
- `func main()`: Defines a function that can be run.
- `fmt.Println()`: Call a function from a package.

Thats it, you are ready to go :trollface:!

## Resources

- https://www.udemy.com/course/go-the-complete-developers-guide
- https://github.com/StephenGrider/GoCasts/tree/master/diagrams
