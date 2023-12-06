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
