# IO in Go
- Reader, Writer, butio
## Reader
```
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```
- Reader example
```
reader := strings.NewReader("SAMPLE")
var newString strings.Builder
buffer := make([]byte, 4)
for {
    numBytes, err := reader.Read(buffer)
    chunk := buffer[:numBytes]
    newString.Write(chunk)
    fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
    if err == io.EOF {
        break
    }
}
fmt.Printf("%v\n", newString.String())
```

## bufio
- provides Reader/Writer buffering
- No need to manually manage buffers or construct data
```
source := strings.NewReader("SAMPLE")
buffered := bufio.NewReader(source)
//below or buffered.ReadBytes
newString, err := buffered.ReadString('\n')
if err == io.EOF {
    fmt.Println(newString)
} else {
    fmt.Println("something's off here..")
}
```
- `bufio.Scanner` provides convenient functions
```
//reading from standard input
scanner := bufio.NewScanner(os.Stdin)
lines := make([]string, 0, 5)
for scanner.Scan() {
    lines = append(lines, scanner.Text())
}
if scanner.Err() != nil {
    fmt.Println(scanner.Err())
}
fmt.Printf("Line count: %v\n", len(lines))

for _, line := range lines {
    fmt.Printf("Line: %v\n", line)
}
```

## Writer
```
buffer := bytes.NewBufferString("")
numBytes, err := buffer.WriteString("SAMPLE")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("Wrote %v bytes: %c\n", numBytes, buffer)
}
```

## Recap
-  `Reader`/`Writer` are interfaces that allow reading from and writing to I/O sources

- Using `Reader` directly requires manually populating a buffer
- The `bufio` stdlib packages provides auto-buffered reads
- The `bufio.Scanner` type can automatically read and delimit inputs