# Goroutines
- Goroutine allow functions & closures to run concurrently
- Use the  `go` keyword to create a new goroutine
- The function that starts a goroutine will `not` wait for it to finish (non-blocking)
    - Both the calling function and goroutine will run to completion
- Closure captures are shared among all goroutines
    - Easy to parallelize code

## Example
- Basic
```
func count(amount int) {
    for i:= 1; i <= amount; i++ {
        time.Sleep(100.time.Millisecond)
        fmt.Println(i)
    }
}

func main() {
    go count(5)
    fmt.Println("wait for goroutine")
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("end program)
}
```
- Closure
```
counter := 0
wait := func(ms time.Duration) {
    time.Sleep(ms * time.Millisecond)
    counter++
}

fmt.Println("Launching goroutines")
go wait(100)
go wait(900)
go wait(1000)

fmt.Println("Launched.")
fmt.Println("Counter =", counter)
time.Sleep(1100 * time.Millisecond)
fmt.Println("Waited 1100ms.")
fmt.Println("Counter =", counter)
```