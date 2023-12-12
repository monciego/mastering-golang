# Go Basic Types

Go comes with a couple of built-in basic types:

- `int`: A number WITHOUT decimal places (e.g., -5, 10, 12, etc).
- `float64`: A number WITH decimal places (e.g., -5.2, 10.123, 12.9, etc).
- `string`: A text value (created via double quotes or backticks: "Hello World", `Hi everyone`).
- `bool`: `true` or `false`.

Noteworthy "niche" basic types:

- `uint`: An unsigned integer, strictly non-negative (e.g., 0, 10, 255, etc).
- `int32`: A 32-bit signed integer with a range from -2,147,483,648 to 2,147,483,647 (e.g., -1234567890, 1234567890).
- `rune`: An alias for `int32`, represents a Unicode code point (i.e., a single character), used for Unicode characters (e.g., 'a', 'ñ', '世').
- `uint32`: A 32-bit unsigned integer, representing values from 0 to 4,294,967,295.
- `int64`: A 64-bit signed integer, with a range from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
- There are more types like `int8` or `uint8` which work similarly (integer with a smaller number range).

## Null Values

All Go value types come with a "null value," the value stored in a variable if no other value is explicitly set.

For example:

```go
var age int // age is 0
```

Here's a list of the null values for the different types:

`int`: 0
`float64`: 0.0
`string`: "" (empty string)
`bool`: false

# `fmt.Scan()` Limitations

The `fmt.Scan()` function in Go is a convenient tool for capturing and utilizing user input from the command line. However, it comes with a notable limitation: handling multi-word input values is not straightforward with this function.

When attempting to retrieve text that consists of more than a single word, challenges arise with `fmt.Scan()`. The function may not handle whitespace-separated inputs as expected, leading to complications in processing multi-word responses.

Consider alternative methods or input handling functions if your application requires the capture of multi-word inputs.

# Working with Packages in Go

In the world of Go programming, mastering packages is essential for organizing and structuring your code effectively. This involves:

## 1. Splitting Code Across Multiple Files

- Break down your code into logical units by splitting it across multiple files. This not only enhances readability but also makes maintenance and collaboration more manageable.

## 2. Splitting Code Across Multiple Packages

- Embrace modularity by dividing your code into separate packages. Each package encapsulates related functionalities, promoting a clean and scalable architecture.

## 3. Importing and Using Custom Packages

- Leverage the power of reusable code by importing and utilizing custom packages. Whether built-in or external, packages streamline development and foster code reuse.
