# Day 6 - Wait For It

Calculating the distance a boat will go in a given total time after waiting a given time is relatively trivial: distance = (wait time) * (total time - wait time). This makes it tempting to iterate through all possible waiting times and count those where we beat the record, but there is a far more efficient solution.

What we need to know is the number of integers for which the following statement holds: `record distance < (wait time) * (total time - wait time)`. Knowing that the wait time is our variable, well call it x. We'll call the record distance d and the total time t. We have:

```math
d \lt x * (t - x)
```
```math
d \lt tx - x^2
```
```math
0 \lt -x^2 + tx - d

```

That's a [quadratic equation](https://en.wikipedia.org/wiki/Quadratic_equation) in the format $ax² + bx + c = 0$. As we all remember from high school, we can use the quadratic formula to find the zeroes for that kind of equation.

$$
x = \frac{-b +- \sqrt(b^2 -4ac)}{2a}
$$

The range defined by those zeroes is where our input beats the record. We just have to count how many integers are in that range (including the edges), and we have our answer.

