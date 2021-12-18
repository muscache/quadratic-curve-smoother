# quadratic-curve-smoother
Stupid simple curve smoothing using quadratic polynomials.

Turn this:
![](https://github.com/muscache/quadratic-curve-smoother/blob/master/static/rough.png?raw=true)

Into this:
![](https://github.com/muscache/quadratic-curve-smoother/blob/master/static/smooth.png?raw=true)


The code is written in Go, but it can be easily ported to any language.

## How it works
We first decide on a window size (`WindowSize`). Larger values will make the curves smoother, but will miss large fluctuations in the graph. Smaller values capture high frequency fluctuations better, but do not make the curve as smooth.

After splitting the data into windows, we operate on a single window at a time.
First, determine the local maxima and minima for the window. In the case of more than one maxima, the last value is chosen. Of course, to capture *all* peaks, you need to use a larger window size.

We have this simple equation:
![](https://github.com/muscache/quadratic-curve-smoother/blob/master/static/initial_equation.png?raw=true)

Where `Max` is the local maxima, `Min` is the local minima and `ArgMax` is the index of the local maxima

Our next step is to determine `Slope`. For that, simply transpose the above (while setting y = 0) to get this:
![](https://github.com/muscache/quadratic-curve-smoother/blob/master/static/transposed_equation.png?raw=true)

What should `x` be here? It should be the value furthest away from `ArgMax`!
* If `ArgMax` is closer to zero, `x` should be `WindowSize-1`.
* Else if `ArgMax` is closer to `WindowSize-1`, `x` should be zero.
* Else, `x` can be either `WindowSize-1` or zero. We use zero.

This `x` value is referred to as "poi" in the code.
Why we have this rule-based selection of "poi" is left as an exercise to the reader.

At this point we have everything we need to actually smooth the curve. Simply iterate over the points in the window and set the `i`th value to:
![](https://github.com/muscache/quadratic-curve-smoother/blob/master/static/final_equation.png?raw=true)


## Limitations
Since it uses second degree polynomials, this algorithm cannot correctly smooth windows with more than one local maxima (if there are `n` local maximas, we need an `n+1` degree polynomial to handle it). If someone has a solution that generates such polynomials, please open an issue, I would love to see it!

## Comparison with the Savitsky-Golay filter
My project solves a very specific problem, which is to ease into "exploding" values.
For example, you might want to ease an LED on and off instead of toggling it sharply. This code solves that problem. In this way, my project is more similar to polynomial regression than to the S/G filter.
