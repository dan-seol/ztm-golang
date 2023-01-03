# Closure/Lambda's: Function Literals in Go
- `Function literals` can be passed to other functions as arguments
    -  They can capture surrounding variables
- Type aliases are helpful when passing function literals to other functions
- Function literals can be returned from a function directly, or assigned to a variable

## Example
- Basic example
```
func helloWorld() {
    fmt.Printf("Hello, ")
    world := func() {
        fmt.Printf("World\n")
    }
    world()
    world()
    world()
    world()
}
```
- Surrounding a variable
```
func customMsg(fn func(m string), msg string) {
    msg = string.ToUpper(msg)
    fn(msg)
}

func surround() func(msg string) {
    return func(msg string) {
        fmt.Printf("%.*s\n", len(msg), "-----------")
        fmt.Println(msg)
        fmt.Printf("%.*s\n", len(msg), "-----------")
    }
}

customMsg(surround(), "hello")
```
- Function literal as a variable
```
func calculatePrice(
    subTotal float64,
    discountFn func(subTotal float64) float64) float64 {
        return subTotal - (subTotal * discountFn(subTotal))
    }


discount := 0.1
discountFn := func(subTotal float64) float64 {
    if subTotal > 100.0 {
        discount += 0.1
    }
    if discount > 0.3 {
        discount = 0.3
    }
    return discount
}

totalPrice := calculatePrice(20.0, discountFn)
```
- Function literal with a type alias
```
type DiscountFunc func(subTotal float64) float64

func calculatePrice(
    subTotal float64,
    discountFn DiscountFunc) float64 {
        return subTotal - (subTotal * discountFn(subTotal))
    }


discount := 0.1
discountFn := func(subTotal float64) float64 {
    if subTotal > 100.0 {
        discount += 0.1
    }
    if discount > 0.3 {
        discount = 0.3
    }
    return discount
}

totalPrice := calculatePrice(20.0, discountFn)
```