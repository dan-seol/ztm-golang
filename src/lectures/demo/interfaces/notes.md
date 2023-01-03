# Interfaces in Go

- Interfaces are implicitly implemented
    - When a type has all rec-func's req. by interface, then it counts as implemented
- Func's operating on interfaces should never accept a pointer to an interface
    - Caller determines whether pointer or value is used
- Prefer multiple interfaces with a few functions over one large interface
- When implementing a pointer receiver func, all func accepting the interface will only accept pointers
    - if self-modification is needed, implement all interface func as receiver func for consistency
- it is sometimes needed to access the underlying type that implements an interface
 - call functions, make modifications, etc.
 ```
 func ResetWithPenalty(r Resetter) {
    if player, ok := r.(Player); ok {
        player.health = 50
    } else {
        r.Reset()
    }
 }
 ```