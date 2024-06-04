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

### Same Line Declarations

You can declare multiple variables on the same line:

```go
mileage, company := 80276, "Tesla"
```

The above is the same as:

```go
mileage := 80276
company := "Tesla"
```

### Type Sizes

`ints`, `uints`, `floats`, and `complex` numbers all have type sizes.

```go
int  int8  int16  int32  int64 // whole numbers

uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)

```

The size (8, 16, 32, 64, 128, etc) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The standard sizes that should be used unless the developer has a specific need are:

- `int`
- `uint`
- `float64`
- `complex128`

Some types can be converted like this:

```go
temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)
```

### Constants

Constants are declared with the `const` keyword. They can't use the `:=` short declaration syntax.

```go
const pi = 3.14159
```

Constants can be character, string, boolean, or numeric values. They can not be more complex types like slices, maps and structs.

As the name implies, the value of a constant can't be changed after it has been declared.

### Computed Constants

Constants must be known at compile time. They are usually declared with a static value:

```go
const myInt = 15
```

However, constants **_can be computed_** as long as the computation can happen at compile time.

For example, this is valid:

```go
const firstName = "Clark"
const lastName = "Kent"
const fullName = firstName + " " + lastName
```

That said, you **_cannot_** declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:

```go
// the current time can only be known when the program is running
const currentTime = time.Now()
```

---

### Questions and Answers

Q: When should you elect to **NOT** use a 'default type'?
A: When performance and memory are the primary concern
