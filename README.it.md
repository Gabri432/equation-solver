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

## Struttura del progetto
- (main)
  - equation.go
  - equation_test.go
  - internal_functions.go
  - internal_functions_test.go
  - license
  - README.md
  - README.it.md

## Note aggiuntive
- Se vuoi verificare l'accuratezza del programma usa i seguenti link:
  - [Cubic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/cubicequation.php)
  - [Quadratic Equations Solver](https://www.calculatorsoup.com/calculators/algebra/quadratic-formula-calculator.php)