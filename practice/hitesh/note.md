## Go Tutorial with Hitesh choudhury
# [YouTube playlist Link](https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa)

## Type conversion:
In go you can convert a variable from one type to another using the following methods.
#### Numeric Conversions:
- int to float: `float64(myInt)`
- float to int:` int(myFloat)` (truncates the decimal part)
  
#### String Conversions:
- string to int: `strconv.Atoi(myString)` or strconv.
- `ParseInt(myString, 10, 64)`
- int to string: `strconv.Itoa(myInt)`
- string to []byte: `[]byte(myString)`
- []byte to string: `string(myBytes)`

## Time Function in Go:
- `time.Now()` returns the current time.
- `time.Now().Add(time.Duration(1 * time.Second))` adds 1 second to the current time.
- `time.Now().Sub(time.Now())` returns the time difference between two times. 
- `time.Now().Unix()` returns the time in seconds since the Unix epoch (January 1, 1970).
- `time.Now().UnixNano()` returns the time in nanoseconds since the Unix epoch.
- `time.Now().Format(time.RFC3339)` returns the time in RFC3339 format. 
  
ETC.

