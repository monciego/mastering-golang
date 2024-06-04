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
