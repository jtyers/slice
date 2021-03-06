# Godash Slice Utilities

An implementation of utility functions inspired by [Lodash](https://lodash.com) for the Go language, focusing on efficiency without the use of reflection.

This is based on the original [slice](https://github.com/go-dash/slice) repo with support for Go modules and code generation in a more idiomatic style (for example, to use with `go generate`).

All slice utilities opt for immutability so always a new copy of the array is returned instead of a different slice over the same array. Some of the motivation behind the original library is discussed in this [blog post](https://medium.freecodecamp.org/lodash-in-go-language-without-reflection-1d64b5115486).

## Installation

1. Go get the library:

    ```
    go get github.com/jtyers/slice
    ```
  
2. In your project, use `go generate` to generate `slice` types for the types you need.

    ```go
    //go:generate go-dash-slice -type string

    ```

    This will string `string.go` in the `go-dash` subdirectory. You can change the directory and package name for the generated code with `-dir` and `-package` respectively. You can also use `-build-tag` to add build tags to the head of the generated files.
  
3. Run `go generate`.

4. Use in your code.
    
    ```go
    import (
      "<your project>/godash"
    )
    
    godash.NewStringSlice()

    ```


You should check generated code into your repository and re-run `go generate` at least whenever you update this library or change the `//go:generate` comment in your code. I tend to run `go generate` as part of my tests (as in `go generate && go test`).

&nbsp;
## Methods

* [`Reverse`](#_reverseslice)
* [`Uniq`](#_uniqslice)
* [`Filter`](#_filterslice-func)
* [`Map`](#_mapslice-func)
* [`Reduce`](#_reduceslice-func-initial)
* [`Contains`](#_containsslice-slice)
* [`Concat`](#_concatslice-slice)
* [`First`](#_firstslice)
* [`Last`](#_lastslice)
* [`Drop`](#_dropslice-n)
* [`DropRight`](#_droprightslice-n)
* [`Chain`](#_chainsliceactionactionvalue)
* [`Value`](#_chainsliceactionactionvalue)

&nbsp;
#### `_.Reverse(slice)`

Returns a new array in reverse order.

```go
_int.Reverse([]int{1, 2, 3})
// => []int{3, 2, 1}
```

#### `_.Uniq(slice)`

Returns a new array without duplicates (all elements are unique).

```go
_int.Uniq([]int{1, 2, 1, 3, 3})
// => []int{1, 2, 3}
```

#### `_.Filter(slice, func)`

Returns a new array of all elements which the function predicate returns true for.

```go
even := func (element int, index int) bool {
  return element % 2 == 0
}
_int.Filter([]int{1, 2, 3, 4}, even)
// => []int{2, 4}
```

#### `_.Map(slice, func)`

Returns a new array of all elements after the function has been executed on them.

```go
double := func (element int, index int) int {
  return 2 * element
}
_int.Map([]int{1, 2, 3, 4}, double)
// => []int{2, 4, 6, 8}
```

#### `_.Reduce(slice, func, initial)`

Reduces the slice to a single value which is the accumulated result of running each element through the function.

```go
sum := func (acc int, element int, index int) int {
  return acc + element
}
_int.Reduce([]int{1, 2, 3, 4}, sum, 0)
// => int(10)
```

#### `_.Concat(slice, slice)`

Returns a new array which is the first slice with the second concatenated at its end.

```go
_int.Concat([]int{1, 2, 3}, []int{4, 5})
// => []int{1, 2, 3, 4, 5}
```

#### `_.Contains(slice, item)`

Returns `true` if the given item is present in the slice. Equality (`==`) is used for comparisons, which means that structs with equal field values will be considered equal. Where `slice` is a slice of pointers, dereferenced values will also be checked for equality, meaning that two different pointers to the same underlying variable will also be considered equal.

```go
_int.Contains([]int{1, 2, 3}, 3)
// => true

_int.Contains([]int{1, 2, 3}, 4)
// => false
```

#### `_.First(slice)`

Returns the first element in the slice.

```go
_int.First([]int{1, 2, 3, 4})
// => int(1)
```

#### `_.Last(slice)`

Returns the last element in the slice.

```go
_int.Last([]int{1, 2, 3, 4})
// => int(4)
```

#### `_.Drop(slice, n)`

Returns a new array where n elements are dropped from the beginning.

```go
_int.Drop([]int{1, 2, 3, 4, 5}, 2)
// => []int{3, 4, 5}
```

#### `_.DropRight(slice, n)`

Returns a new array where n elements are dropped from the end.

```go
_int.DropRight([]int{1, 2, 3, 4, 5}, 2)
// => []int{1, 2, 3}
```

#### `_.Chain(slice).Action().Action().Value()`

Chains multiple actions together and runs each on the result of the previous one. `Value()` returns the final result.

```go
_int.Chain([]int{1, 2, 1, 3}).Uniq().Reverse().Value()
// => []int{3, 2, 1}
```

&nbsp;
## Working with different types

In order to avoid inefficient reflection, the library creates dedicated implementations for each type you need.

In the original [Lodash](https://lodash.com), the library is used through the underscore character `_`. For example: `_.uniq()`.

We keep the same convention, except that the underscore is followed by the type. For example: `_int.Uniq()` for integers, `_string.Uniq()` for strings.

#### Primitive types (int, string, etc)

Simply import the relevant subset of the library with the type appearing after the underscore:

```go
import "github.com/go-dash/slice/_string"

func main() {
    _string.Uniq([]string{"aa", "bb", "aa"})
    // => []string{"aa", "bb"}
}
```

#### Custom types (structs)

Do the same thing, just add a comment afterwards of where the struct is defined.

```go
import "github.com/go-dash/slice/_Person" // github.com/my-user/my-repo/person

func main() {
    _Person.Uniq([]Person{Person{"John", 18}, Person{"Rachel", 17}, Person{"John", 18}})
    // => []Person{Person{"John", 18}, Person{"Rachel", 17}}
}
```

&nbsp;
## Running the tests

From the root directory run `./test.sh`.

&nbsp;
## Motivation

This is mostly a thought experiment coming from the love of lodash and missing basic functional operations for slices in golang. Implementations based on [reflection](https://github.com/robpike/filter) miss the point IMO since they're inefficient in runtime. Since golang doesn't currently support generics, code generation is the closest we can come.

If you like this approach and want to contribute to porting all of lodash to golang, don't hesitate to come help!

&nbsp;
## License

MIT

&nbsp;
## Who made this

Godash is developed by the [Orbs.com](https://orbs.com) engineering team. Orbs is a public blockchain infrastructure for decentralized consumer applications with millions of users. Orbs core has an open source [implementation](https://github.com/orbs-network/orbs-network-go) in golang.
