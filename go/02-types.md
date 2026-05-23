
## Boolean

Used to represent `true` or `false`.

```go
var isAdmin bool = true

fmt.Println(isAdmin)
```

---

## Integer

```go
var age int = 25

fmt.Println(age)
```

### Integer Variations

Signed integers (positive and negative numbers):

| Type  | Range                                                   |     |
| ----- | ------------------------------------------------------- | --- |
| int8  | -128 to 127                                             |     |
| int16 | -32,768 to 32,767                                       |     |
| int32 | -2,147,483,648 to 2,147,483,647                         |     |
| int64 | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |     |

```go
var smallNumber int8 = 127
var mediumNumber int16 = 32000
var largeNumber int32 = 2000000000
var hugeNumber int64 = 9000000000000000000
```

`int` depends on your system architecture:
- 32-bit systems -> behaves like `int32`
- 64-bit systems -> behaves like `int64`

---

### Unsigned Integers

Unsigned integers store only positive numbers.

| Type | Range |
|---|---|
| uint8 | 0 to 255 |
| uint16 | 0 to 65,535 |
| uint32 | 0 to 4,294,967,295 |
| uint64 | 0 to 18,446,744,073,709,551,615 |

```go
var points uint8 = 255
var money uint16 = 65000
var population uint32 = 4000000000
```

`uint` also depends on your system architecture.

---

## Float


```go
var price float64 = 19.99

fmt.Println(price)
```

### Float Variations

| Type | Example |
|---|---|
| float32 | 3.14 |
| float64 | 19.999999999 |

```go
var pi float32 = 3.14
var productPrice float64 = 199.99

fmt.Println(pi)
fmt.Println(productPrice)
```

`float64` is more precise than `float32` and is commonly used in Go applications.

---

## String


```go
var name string = "Jonathas"

fmt.Println(name)
```

---

## Rune

Represents a Unicode character.

```go
var letter rune = 'A'

fmt.Println(letter)
```

---

## Byte

Alias for `uint8`.

Usually used for raw binary data.

```go
var value byte = 255

fmt.Println(value)
```

---

## Complex Numbers

Used to work with complex numbers.

```go
var number complex64 = 1 + 2i

fmt.Println(number)
```

### Complex Variations

```go
complex64
complex128
```

---

# Short Variable Declaration

Go also allows short variable declarations using `:=`.

```go
name := "Go"
age := 15
isLearning := true
```

---

# Zero Values

Variables declared without a value receive a default zero value.

| Type | Zero Value |
|---|---|
| bool | false |
| int | 0 |
| float | 0 |
| string | "" |

```go
var name string
var age int
var active bool

fmt.Println(name)
fmt.Println(age)
fmt.Println(active)
```