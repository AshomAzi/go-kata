## The Mathematical Divider (The Zero-Trap)

* **The Goal**: Understand how multi-value returns pass custom mathematical failures back up the execution stack.

Create a function named SafeDivide that handles numeric parameters without risking an application panic state.

* **The Rules**:

The function accepts two float values: numerator and denominator.

It must check if the denominator is exactly 0. If it is, return a result value of 0.0 along with a custom error message: "division by zero is undefined".

Otherwise, return the calculated quotient and a nil error.