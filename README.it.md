# equation-solver
Un programma che risolverà una qualunque equazione lineare, quadratica e cubica.

## Come usare la libreria
- Scarica il package 'equationsolver':
```
go get -u github.com/Gabri432/equation-solver
```

- Esempio di utilizzo
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

## Funzioni
```go
EvaluateEquations(equation string) EquationSolution
```
  - Prende l'equazione in forma di stringa e ritorna il risultato di tipo EquationSolution

```go
ValidateEquation(equation string) string
```
  - Controlla se l'equazione inserita dall'utente è valida.

## Tipi
### Polynom
```go
// Crea un polinomio nella forma: ax^3+bx^2+cx+d
type Polynom struct {
	a float64
	b float64
	c float64
	d float64
}
```
### EquationSolution
```go
type EquationSolution struct {
	realSolutions    []float64    // insieme di soluzioni reali
	complexSolutions []complex128 // insieme di soluzioni complesse
	errorDescription string       // messaggio d'errore
```

## Struttura del progetto
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

## Note aggiuntive
- Se vuoi verificare l'accuratezza del programma usa i seguenti link:
  - [Cubic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/cubicequation.php)
  - [Quadratic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/quadratic-formula-calculator.php)