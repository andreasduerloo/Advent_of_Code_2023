# Day 4 - Scratchcards

Two things are going on with the **second star** of this problem that are worth pointing out. Before I get into those, here's how a naive approach to this problem would work (and it's very possible it would still compute within an acceptable time):

1. Make a long list containing all your scratchcards.
2. Look at each item in the list in turn, and append the correct scratchcards to the end of the list.
3. Keep going (front-to-back) until you hit the end.
4. Take the length of the list.

Two things can be optimized here:

1. We are making a *very* long list, which brings potential memory and performance problems.
2. We are recalculating the same thing a lot of times.

## 1. Layers of abstraction

This problem reminded me of [Advent of Code 2021 day 14](https://adventofcode.com/2021/day/14). In that problem, you had to count the number of elements in a polymer which roughly doubled in length every 'tick'. After enough ticks (2<sup>n</sup>, after all) the polymer would become far too long to iterate over efficiently. That meant the 'naive' approach of **simulating** the entire thing was not the optimal (or even a viable) solution. Additionally, we **were not being asked to simulate the entire polymer**, we just had to count each of the elements in it, and there were only a limited number of those.

The same ting goes here: **we are not being asked to produce the entire list of scratchcards**, we just have to count how many there will be at the end. This can be achieved without adding a single item to a list.

> [!Note]
> Two takeaways:
> 1. Don't lose track of what is being asked and what isn't.
> 2. It's often far better to represent the relevant attributes of what we are simulating, rather than actually simulating it (which is usually our first instinct).

## 2. Working back-to-front (or how I learnt to stop calculating and use the cache)

Even if we calculate the number of tickets rather than simulate it, we still have a problem. Consider this example:

> The first scratchcards has two winning numbers, so we add the total cards for scratchcard 2 and 3. However, the total scratchcards for card 2 equals one plus the sum of the total for cards 3, 4,... etc.

This is a recursive problem: we can keep expressing the amount of cards for a given card as a sum of other cards (and they are again sums of other cards, etc.) until we hit a base case. This approach might work, but you will be recalculating the same thing over and over again. In the example above, we would already be calculating the number of cards for scratchcard 2 twice, and probably far more in reality. Compare to calculating a fibonacci number recursively: a massive chunk of your recursion tree is just calculating fib(1) over and over again.

There are two answers to this, and I'm using both of them:

1. **Memoization**, i.e. storing the result of a calculation so that I calculate things once and just look up the result the next time I would need to calculate it again.
2. Working back-to-front rather than front-to-back.

To elaborate on the second point: I start from the base case (the last card), for which we know it will not add a new card. In other words, that card counts for one card (itself), the nwe go to the card before that, which - if it has at least one winning number - counts for itself and the last card, so two. This way I am building and using the cache so efficiently that I only ever calculate the value for the card I am currently looking at (the rest **will** already be calculated and cached). I am never calculating something to calculate something else (as opposed to the example above). The two points are complementary: **working backwards maximizes the number of cache hits and minimizes the number of calculations.**

In a very basic way, this is an example of bottom-up dynamic programming. Working front-to-back recursively (with memoization) would be top-down dynamic programming.

> [!Note]
> Two more takeaways:
> 1. Working back-to-front might be less intuitive, but **can** be far more efficient.
> 2. Memoization is a cheat code, all the more combined with recursion.
