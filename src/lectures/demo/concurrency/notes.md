# Concurrency in Go
- Concurrent code allows full utilization of available compute resources
- Go automatically abstracts threads and async operations
    - Threaded code is used to make parallel computations on cores
    - Asynchronous code may pause/resume and is used for waiting on resources (like networks)
- Synchronized code runs deterministically
- Concurrent code runs non-deterministically
    - Synchronization or other techniques are required to ensure proper program behavior

# Synchronization in Go
- Managing data across multiple goroutines can become problematic and hard to debug
    - Multiple goroutines can change the same data leading to unpredictable results
    - Using channels to communicate - not always ideal
- `Synchronization` solves this issue, enabling:
    - Waiting for goroutines to finish
    - Prevents multiple goroutines from modifying data simultaneously

## Mutex
- `mut`ual `ex`clusion
- Provides a way to `lock` and `unlock` data
    - Locked data cannot be accessd by any other goroutine until it is locked
    - While locked, all other goroutines are `blocked` until the mutex is unlocked
        - Execution waits until lock is available, or if `select` is used
- Helps reduce bugs when working with multiple goroutines

```
import "sync"

type SyncedData struct {
    inner map[string]int
    mutex sync.Mutex
}

func (d *SyncedData) Insert(k string, v int) {
    d.mutex.Lock()
    d.inner[k] = v
    d.mutex.Unlock()
}

func (d *SyncedData) Get(k string) int {
    d.mutex.Lock()
    data := d.inner[k]
    d.mutex.Unlock()
    return data
}

func main() {
    data := SyncedData{inner: make(map[string]int)}
    data.Insert("sample", 5)
    data.Insert("test", 2)
    fmt.Println(data.Get("sample"))
    fmt.Println(data.Get("test"))
}
```

### Deferred Unlock
- `defer` can be used the mutex gets unlocked
```
func (d *SyncedData) Insert(k string, v int) {
    d.mutex.Lock()
    defer d.mutex.Unlock()
    d.inner[k] = v
}

func (d *SyncedData) Get(k string) int {
    d.mutex.Lock()
    defer  d.mutex.Unlock()
    return d.inner[k]
}
```

## Wait Groups
- Wait groups enable an application to wait for goroutines to finish
- Operates by incrementing a counter whenever a goroutine is added and decrementing when it finishes
    - Waiting on the group will block execution until the counter is 0
```
var wg sync.WaitGroup
sum := 0
for i := 0; i < 20; i++ {
    wg.Add(1)
    value := i
    go func() {
        defer wg.Done()
        sum += value
    }()
}
wg.Wait()
fmt.Println("sum = ", sum)
```

## Recap - Synchronization
- Data can be safely accessed across goroutines using a `mutex`
    - Lockng a mutex prevents other goroutines from locking it
    - Always remember to unlock a mutex
- It is possible to wait for goroutines to finish with a `wait group` - a counting semaphore
    - Add 1 per goroutine to the wait group, then use `.Done()` in each goroutine to decrement the group counter
- Using defer makes it simple to unlock mutexes and when working with wait groups

# Concurrency Patterns
- Channels and goroutines must be cleaned up at some point
    - Failing to do so is a resource leak
- Concurrency patterns offer ways to clean up goroutines
    - Also help with managing data flow
- Multiple patterns available based on the situation

## Pipelines
- are multiple stages of operations connected using channels
    - at least 1 input channel and 1 output channel
- use goroutines to manage the sending/receiving of values
- multiple options for managing and combining stages
- each stage of the pipeline depends on the previous stage

## Fan-in
- Multiple input channels, 1 output channel
- Examples:
    - Serializing requests to purchase an item
    - Segmented processing of a single item

## Close / Cancellation
- Closing channel indicates end of data stream
- Pipeline stage should close channel when work is done
    - Ckeanup goroutines / no resource leaks / unblock

## Request Quit
- Dedicated "quit" channel
- Stage listens on incoming "data" channel and "quit" channel
    - if data/signal comes in on "quit" channel, goroutine shuts down
- Can use one or multiple "quit" channels
    - Multiple: must send enough signals for each stage to quit

## Context
- Similar to request quit
- No need to manually manage channels
- Calling a quit function cancels all operations using the Context
- Can associate data with each Context
    - IP addresses, session Id's, node identifiers, etc.

## Generators
- On-demand yielding of items
- Items generated until bounded channel fills
    - Reduced processing
- Whenever items are read from the channel, new ones are calculated as needed

## Recap - Concurrency Patterns
- Concurrency patterns are used to manage goroutines
- Pipelines are broken into "stages" wuth a goroutine for each stage
    - Stages communicate with channels
- Multiple ways to clean up pipelines:
    - Close channels / send "quit" signal / use Context
- Generators can be used to yield items on-demand