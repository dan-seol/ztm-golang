# Errors in Go
## Error Handling
- Go has no exceptions
- Errors are obtained as the last return value from a function
    - Encodes failure as a part of the function signature
        - Simple to determine if a function can fail
    - Returns `nil` if no error occurred
- Errors implement the `error` interface from `std`
    - one function to implement: `Error() string` 
## Basics
- The `errors` stdlib package can generate simple errors with the `New()` function
```
import "errors"

func divide(lhs, rhs int) (int, error) {
    if 0 == rhs {
        return 0, errors.New("cannot divide by zero")
    }
    return rhs / lhs, nil
}
```
## interface
```
type error interface {
    Error() string
}
```
## implementation
```
type DivError struct {
    a, b int
}

func (d *DivError) Error() string {
    return fmt.Sprintf("Cannot divide by zero: %d / % d", d.a, d.b)
}
```
- Always implement error as a receiver function
    - prevents comparison problems if error is inspected

## Working with errors
- `errors.Is()` - to check
- `errors.As()` - converting error
- always check if `err != nil` for functions returning error type