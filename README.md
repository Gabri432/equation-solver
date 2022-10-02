# equation-solver
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Gabri432/equation-solver)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/Gabri432/equation-solver)
![GitHub](https://img.shields.io/github/license/Gabri432/equation-solver)

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
    solution := eq.EvaluateEquation("x^3-4x=5x^2-1")
    fmt.Println("real solutions: ", solution.RealSolutions, "\ncomplex solutions:", solution.ComplexSolutions)
    myPolynom := eq.Polynom{A: 1, B: 5, C: -4, D: 1}
    fmt.Println(myPolynom.SolveEquation())
}

```
=== Output ===

real solutions: [-0.8752... 0.2013... 5.6739...]

complex solutions: []

{[0.38196... 2.61803...] [] }

## Functions
```go
EvaluateEquations(equation string) EquationSolution
```
  - Takes the equation in form of string and returns the result of type EquationSolution

```go
ValidateEquation(equation string) string
```
  - Checks if the user inserted equation is valid.

## Types
### Polynom
```go
// Creates a polynom in the form: ax^3+bx^2+cx+d
type Polynom struct {
	A float64
	B float64
	C float64
	D float64
}
```
### EquationSolution
```go
type EquationSolution struct {
	RealSolutions    []float64    // set of real solutions
	ComplexSolutions []complex128 // set of complex solutions
	ErrorDescription string       // error message
```

## Project Structure
- (main)
  - equation.go
  - equation_test.go
  - internal_functions.go
  - internal_functions_test.go
  - license
  - README.md
  - README.it.md
  - .github
    - CONTRIBUTING.it.md
    - CONTRIBUTING.md

## Contribuire al progetto
- If you want to add a feature or making a fix check out this page explaining how to do it: [Contributing to equation-solver](https://github.com/Gabri432/equation-solver/blob/master/.github/CONTRIBUTING.it.md)

## Notes
- If you want to check the accuracy of the program use the following links:
  - [Cubic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/cubicequation.php)
  - [Quadratic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/quadratic-formula-calculator.php)