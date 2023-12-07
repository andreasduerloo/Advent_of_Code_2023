# Advent of Code 2023

These are my solutions for the 2023 edition of [Advent of Code](https://adventofcode.com/).

I see Advent of Code as a great way to become a better programmer (as well as a fun puzzle to share and discuss with friends and colleagues). As such my goal is to do better than I did last year, and to score **more than 30 stars** during the event (in other words, before December 26th).

## Progress

- **Goal progress: 14/31 :star:**
- **Total score: 14/50 :star:**

## How is this repo structured?

The entire repo builds to a single binary, with the `main()` function living in advent.go. To print the solutions for a given day (e.g., the first day), I run the following command:
~~~
go run ./advent.go 1
~~~
The **solutions** for individual days each get their own directory. That directory contains a *lib.go* file with all the code doing the heavy lifting, as well as a *\day_\*\*.go* file that exposes a ´Solve()´ function. That function's job is to read in the puzzle input, pass it off to a parsing function, handle the high-level logic, and return the solutions. This function is called from the ´main()´ function.

> [!Note]
> In the unlikely case I think a solution is particularly interesting or clever, I might add a separate *readme.md* file in that directory.

Additionally, there is a **helpers** directory, containing a *helpers* package. This package exposes useful functions that aren't linked to a single puzzle or day, but can make life a little easier. E.g: returning all ints from a string, or a `Dequeue()` function that allows me to use a slice as a queue.

## Solutions

- [X] Day 1: [Trebuchet?!](https://github.com/andreasduerloo/Advent_of_Code_2023/tree/main/day_01) :star::star:
- [X] Day 2: [Cube Conundrum](https://github.com/andreasduerloo/Advent_of_Code_2023/tree/main/day_02) :star::star:
- [X] Day 3: [Gear Ratios](https://github.com/andreasduerloo/Advent_of_Code_2023/tree/main/day_03) :star::star:
- [X] Day 4: [Scratchcards](https://github.com/andreasduerloo/Advent_of_Code_2023/tree/main/day_04) :star::star: - with readme :bulb:
- [X] Day 5: [If You Give A Seed A Fertilizer](https://github.com/andreasduerloo/Advent_of_Code_2023/tree/main/day_05) :star::star: - needs rework
- [X] Day 6: [Wait For It](https://github.com/andreasduerloo/Advent_of_Code_2023/tree/main/day_06) :star::star: - with readme :bulb:
- [X] Day 7: [Camel Cards](https://github.com/andreasduerloo/Advent_of_Code_2023/tree/main/day_07) :star::star: