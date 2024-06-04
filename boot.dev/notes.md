# Notes

## Ch 1: Introduction

### Go is fast, simple and productive

- Go code generally runs faster than interpreted languages and compiles faster than other compiled languages like C and Rust.

### Comparing Go's Speed

Go is generally faster and more lightweight than interpreted or VM-powered languages like:

- Python
- JavaScript
- PHP
- Ruby
- Java

However, in terms of execution speed, Go does lag behind some other compiled languages like:

- C
- C++
- Rust

Go is a bit slower mostly due to its automated memory management, also known as the "Go runtime". Slightly slower speed is the price we pay for memory safety and simple syntax!

### Go programs are easy on memory

Go programs are fairly lightweight. Each program includes a small amount of "extra" code that's included in the executable binary. This extra code is called the Go Runtime. One of the purposes of the Go runtime is to clean up unused memory at runtime.

## Ch 2: Variables

### Basic Types

Go's basic variable types are:

```go
    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // alias for uint8

    rune // alias for int32
        // represents a Unicode code point

    float32 float64

    complex64 complex128
```

### Short Variable Declaration

Inside a function (like the main function) the := short assignment statement can be used in place of a var declaration. The := operator infers the type of the new variable based on the value. It's colloquially called the walrus operator because it looks like a walrus... sort of.

These two lines of code are equivalent:

```go
var empty string

```

```go
empty := ""

```

```go
numCars := 10 // inferred as an integer
temperature := 0.0 // temperature is inferred as a float because it has a decimal
var isFunny = true // inferred as a boolean

```

Outside of a function (in the global/package scope), every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

### Type Inference

To declare a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

```go
var i int
j := i // j is also an int
```
