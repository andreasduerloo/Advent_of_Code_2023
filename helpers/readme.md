# Helpers

This package contains functions (and some structs) that I think (hope?) will come in handy.

## Higher-order functions
Unless I missed them, the Go standard library does not have some of the higher-order functions and iterators that I have come to love in Elixir (see the [Enum](https://hexdocs.pm/elixir/1.15/Enum.html) module). I wrote the following functions (and will probably add more):

```
FilterSlice[T any](s []T, f func(T) bool) []T
MapSlice[T, U any](s []T, f func(T) U) []U
AnySlice[T any](s []T, f func(T) bool) bool
AllSlice[T any](s []T, f func(T) bool) bool
UniqSlice[T comparable](s []T) []T
```

## Data structures

### Dequeue([]T) (T, []T)
This function allows me to use a go slice as a queue (along with the built-in append()). It takes a slice as an argument and returns the first item along with the rest of the slice (or an empty slice, if we just dequeued the last item). Example:
~~~
queue := []int{1, 2, 3}

// Add an item
queue = append(queue, 4)

// Dequeue the first item
first, rest := Dequeue(queue)
~~~

## Input parsing

### ReGetInts(string) []int
This function (which is probably the most useful by far), returns all the integers found in a string as a slice *of integers*. Given that it's based on a RegEx (hence 'Re'), the integers don't need to be separated from non-digit runes by whitespace.

When you know the input string only contains a single integer, you can grab it through indexing:
~~~
input := "This line contains 1 integer"
theInt := ReGetInts(input)[0]
~~~

## Working with grids

Puzzles based on grids are not uncommon for Advent of Code. I can think of four ways to represent a grid:
- A single slice or array. (You calculate the index from coordinates using some width/height math.)
- A 2D-slice or array. (Careful with indexes: rows and columns vs. x and y, they're the other way around and that's a great source of bugs.)
- A (hash)map with coordinates as keys. (This is what I used in Elixir.)
- An actual graph of structs with pointers to their neighbors.

The first three options all share a problem: whenever you want to interact with a point's neighbors, you have to check for edges to make sure you don't go out of bounds with the index. That generates a lot of uninteresting and error-prone code. The fourth and final option tackles that issue *once* at creation time, and from then on you can access the neighbors without any additional checking. This makes graph algorithms far easier to implement, because you can do things like this:
~~~
[...]
for _, n := range node.neighbors {
[...]
~~~
For that reason it's by far my preferred way of working with grids, and I have two functions to build them for me:

### IGridAsGraph() and RGridAsGraph()
Both of these functions take a (multiline) string as input, which represents a grid with an arbitrary width and height. The difference is that one of them assumes a grid of integers (and takes care of the conversion from rune to int), and the other assumes a grid of runes. Either way, they will return a slice of nodes with populated neighbor pointers. Here's what I mean by node (taking the int version as an example):
~~~
type inode struct {
    value       int
    neighbors   []*inode
}
~~~