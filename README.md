# Getting Started with GO

Pointer is a variable which stores memory address of another variable or data structure

UTF stands for Unicode Transformation Foramt which is encoding standard used for electronic communication.

## Data Types

### Declaration

Data types can be declared as:
`var varialbleName String = "Hello World!"`

or

```
    variableName := "Hello World!"
```

### Int Types

Int types can store values up to:

```
int = (Signed) either 32 bits or 64 bits (See int32 or int64)
int8 = 0 to 255
int16 = -32768 to 32767
int32 = -2147483648 to 2147483647
int64 = -9223372036854775808 to 9223372036854775807

uint = (Unsigned) either 32 bits or 64 bits (See uint32 or uint64)
uint8 = 0 to 255
uint16 = 0 to 65535
uint32 = 0 to 4294967295
uint64 = 0 to 18446744073709551615
uintptr = unsigned integer large enough to store the uninterpreted bits of a pointer value
```

`byte` is alias for `uint8`
`rune` is alias for `int32`

### Floating Types

Float and complex numbers sizes

```
float32     IEEE-754 32-bit floating-point numbers
float64     IEEE-754 64-bit floating-point numbers
complex64   complex numbers with float32 real and imaginary parts
complex128  complex numbers with float64 real and imaginary parts
```

### Overflow and Wraparound

Overflow happens when a signed or pointer value exceeds the range of the variable into which it is stored.

Wraparound happens, instead, when an unsigned integer value exceeds the range that its underlying storage can represent.

### Boolean

Boolean can be either `true` or `false`. `bool` represents boolean type in Go.

### Strings

Strings can be used using back quotes("`") or double quotes(`"`).

If back quotes is used, raw string literal is created.

```
a := `Say "hello" to Go!`
fmt.Println(a)
```

Raw String literals can be used to create multiline strings but ignores special characters.

In strings, 


### Arrays

Arrays are ordered sequence of elements. Here, capacity of an array is stored at created time. Size of array is static, therefore it only allocates memory once.

Example Array of Strings:
`[2] string{"hello", "world"}`

### Slices

Similar to array but can change length i.e. it can dyanmically increase allocated size. We can append the slices.

Exmaple Slices or int:
`[] int{-1, 0, 1, 2, 3}`

### Maps

Dictionary which stores data in key value pair.

`map[string]string{"key1": "value1", "key2": "value2"}`
