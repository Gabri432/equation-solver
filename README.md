# equation-solver
A program that will solve any linear, quadratic and cubic equation.

## How to use it
- Download the 'equationsolver' package:
```
go get -u github.com/Gabri432/equation-solver
```

- Example of usage
```go
package main

import (
    "fmt"
    eq "github.com/Gabri432/equation-solver"
)

func main() {
    solution := eq.EvaluateEquation("x^3-4x=5x^2+0")
    fmt.Println("real solutions: ", solution.realSolutions, "\ncomplex solutions:", solution.complexSolutions)
}

```
=== Output ===

real solutions: [-0.8752..., 0.2013..., 5.6739...]

complex solutions: []

## Project Structure
- (main)
  - equation.go
  - equation_test.go
  - internal_functions.go
  - internal_functions_test.go
  - license
  - README.md
  - README.it.md

## Notes
- If you want to check the accuracy of the program use the following links:
  - [Cubic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/cubicequation.php)
  - [Quadratic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/quadratic-formula-calculator.php)