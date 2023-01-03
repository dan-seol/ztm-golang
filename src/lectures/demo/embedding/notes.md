# Embedded Types/Structs
- Embedding interfaces allows multiple interfaces to be used as one
    - Changes in embedded interfaces automatically propagate throughout codebase
- Embedding structs promotes the embedded struct's fields and methods
    - Easy access to inner struct fields
    - If you define receiver functions sharing the same name as promoted method, they will work as an override

## Embedded Types
```
type Whisperer interface {
    Whisper() string
}

type Yeller interface {
    Yell() string
}

type Talker interface {
    Whisperer
    Yeller
}

func talk(t Talker) {
    fmt.Println(t.Yell())
    fmt.Println(t.Whisper())
}
```

## Embedded Structs
```
type Account struct {
    accountId int
    balance int
    name string
}

type ManagerAccount struct {
    Account
}

mgrAcct := ManagerAccount{Account{2,30,"Cassandra"}}
```

### Promoted Fields and Methods
```
func (a *Account) GetBalance() int {
    return a.balance
}

func (a Account) String() string {
    return fmt.Sprintf("Standard (%v) $%v \"%v\"",
    a.accountId,
    a.balance,
    a.name)
}

func (a ManagerAccount) String() string {
    return fmt.Sprintf("Manager (%v) $%v \"%v\"",
    a.accountId,
    a.balance,
    a.name)
}
```