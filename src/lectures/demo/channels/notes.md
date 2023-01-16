# Channels in Go
- Channels offer bidirectional communication
    - Conceptually the same as a two-ended pipe:
        - Write data in one end and read data out the other
        - `sending` and `receiving`
- Use channels to let goroutines to communicate
    - Can send/receive messages or computational results
- Channel ends can be duplicated across goroutines

## Creation & Usage
```
channel := make(chan int)

// Send
go func() { channel <-1 }()
go func() { channel <-2 }()
go func() { channel <-3 }()

// Receive from channel
first := <-channel
second := <-channel
third := <-channel

fmt.Println(first, second, third)
```

## Details
- Channels can be buffered/unbuffered
    - Unbuffered channels will `block` when sending until a reader is available
    - Buffered channels have a specified capacity
        - Can send messages up to the capacity, even without a reader
- Messages on a channel are FIFO ordering

## Buffered Channel
```
channel := make(chan int, 2) // 2 is the buffer size

// Send
channel <- 1
channel <- 2
go func() { channel <-3 }()

// Receive from channel
first := <-channel
second := <-channel
third := <-channel

fmt.Println(first, second, third)
```

## Channel Selection
- The `select` keyword lets you work with multiple -- potentially blocking -- channels
- `select` vs `switch` - https://stackoverflow.com/questions/38821491/what-is-the-difference-between-switch-and-select-in-go
- Send/Receive attemps are made regardless of blocking status

```
one := make(chan int)
two := make(chan int)

for { //optional also
    select {
        case o := <-one:
            fmt.Println("one:", o)
        case t := <-two:
            fmt.Println("two:", t)
        default: // optional
            fmt.Println("no data to receive")
            time.Sleep(50 * time.Millisecond)
    }
}
```

## Timeout
- The `time` package can be combined with `select` to create timeouts

```
one := make(chan int)
two := make(chan int)

for {
    select {
        case o := <-one:
            fmt.Println("one:", o)
        case t := <-two:
            fmt.Println("two:", t)
        case <-time.Sleep(300 * time.Millisecond):
            fmt.Println("timed out")
            return
    }
}
```

## Recap
- Channels - bidirectional communication pipes
    - send/write end - ... - receive/read end
- The ends of a channel can be duplicated across goroutines
- `select` can be used send/receive on multiple different channels
- Buffered channels are non-blocking, unbuffered channels will block