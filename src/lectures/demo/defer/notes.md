# Defer in Go
- Useful to run operations after functions complete
- The defer keyword can be used to execute code after a function runs
    - Clean up resources, reset data, etc.
```
file, err := os.Open("sample.txt")
if err != nil {
    fmt.Println(err)
    return
}

defer file.Close()

buffer := make([]byte, 0, 30)
bytes, err := file.Read(buffer)

if err != nil {
    fmt.Println(err)
    // File will close with deferred call
    return
}

fmt.Printf("%c\n", bytes)
// File will close with deferred call
```