# Advent of Code 2023

My solutions for the 2023 edition of [Advent of Code](https://adventofcode.com/).

## How is this repo structured?

The entire repo builds to a single binary, with the *main()* function living in advent.go. To print the solutions for a given day (e.g., the first day), I run the following command:
~~~
go run ./advent.go 1
~~~
The **solutions** for individual days each get their own directory. That directory contains a *lib.go* file with all the code doing the heavy lifting, as well as a *\*.go* file that exposes a *Solve()* function. That function's job is to read in the puzzle input, pass it off to a parsing function, handle the high-level logic, and return the solutions. This function is called from the *main()* function. In the unlikely case I think a solution is particularly interesting or clever, I might add a separate *readme.md* file in that directory.

Additionally, there is a **helpers** directory, containing a *helpers* package. This package exposes useful functions that aren't linked to a single puzzle or day, but can make life a little easier. E.g: returning all ints from a string, or a *Dequeue()* function that allows me to use a slice as a queue. 