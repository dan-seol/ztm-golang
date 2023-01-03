# Generics in Go
- Generic functions reduce code duplication
- Generic constraints are interfaces
    - Func parameters / return types are consrrained to a specific set of interfaces
- Use `tilde(~)` to allow approximate types
    - Approximation permits checking underlying types
- `constraints` package has commonly used constraints
-  `comparable` constraint is always available

## Syntax
```
func name[T constraintA, U constraintB | constraintC](a T, b U) T {
    // ...
}
```

## Example
```
func IsEqual[T comparable](a, b T) bool {
    return a == b
}

IsEqual(2, 2)
IsEqual("foo", "bar")
IsEqual('a', 'b')
IsEqual[uint8](4, 4)
```
## Creating a Constraint

```
type Integers32 interface {
    int32 | uint32
}

func SumNumbers[T Integers32](arr []T) T {
    var sum T
    for i:= 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}

nums := []int32{1,2,3}
nums2 := []uint32{1,2,3}
total := SumNumbers(nums)
total2 := SumNumbers(nums2)
```

## Constraints and Type Aliases

```
type Integers32 interface {
    int32 | uint32
}

func SumNumbers[T Integers32](arr []T) T {
    var sum T
    for i:= 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}

type MyInt int32
nums := []MyInt{MyInt(1),MyInt(2),MyInt(3)} // this would cause an error as MyInt, the type alias, is not implementing Integers32 -> has to be an exact match
```

## Approximation


```
type Integers32 interface {
    ~int32 | ~uint32
}

func SumNumbers[T Integers32](arr []T) T {
    var sum T
    for i:= 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}

type MyInt int32
nums := []MyInt{MyInt(1),MyInt(2),MyInt(3)} // this would succeed
```

## Builtin Constraints
| Constraint  | Description |
|---|---|
| any | Any Type |
| comparable | Anything that can be compared for equality |
| Unsigned | All unsigned integers |
| Signed | All signed integers |
| Ordered | Sortable types (numbers, strings) |
| Integer | All integers |
| Float | All floating point numbers |
| Complex | All complex numbers |

## Generic Structure
```
import "constraints"
type MyArray[T constraints.Ordered] struct {
    inner []T
}
func (m *MyArray[T]) Max() T {
    max := m.inner[0]
    for i := 0; i < len(m.inner); i++ {
        if m.inner[i] > max {
            max = m.inner[i]
        }
    }
    return max
}
arr := MyArray[int]{inner: []int{6, 4, 8, 9, 4, 0}}
fmt.Println(arr.Max())
```