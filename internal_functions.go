package equationsolver

import (
	"math"
	"strconv"
	"strings"
)

type Polynom struct {
	a float64
	b float64
	c float64
	d float64
}

type EquationSolution struct {
	realSolutions    []float64
	complexSolutions []complex128
	errorDescription string
}

// Checks if the given string is either a '+' plus, or '-' minus, or '=' equal sign.
//
// If so it returns true, otherwise false.
func isSign(c string) bool {
	switch c {
	case "+", "-", "=":
		return true

	}
	return false
}

// Converts an equation like x^2+2-3x+5=0 in:
//
// x^2; +2; -3x; +5; =; 0
func replaceEquation(equation string) (newEquation string) {
	for i := 0; i < len(equation); i++ {
		expr := string(equation[i])
		if isSign(expr) && i != 0 {
			newEquation += ";" + expr
			if expr == "=" {
				newEquation += ";"
			}
		} else {
			newEquation += expr
		}
	}
	return
}

// Splits the equation whenever it meets the ';' semicolon sign and then separates constants and variables
func splitEquation(equation string) (xVariables, constants []string) {
	var isOverEqualSign bool
	splittedEquation := strings.Split(equation, ";")
	for _, value := range splittedEquation {
		if isOverEqualSign {
			value = changeSign(value)
		}
		switch {
		case strings.Contains(value, "x"):
			xVariables = append(xVariables, value)
		case strings.Contains(value, "="):
			isOverEqualSign = true
		default:
			constants = append(constants, value)
		}
	}
	return
}

// Replace the '-' minus sign with the '+' plus sign, and viceversa
func changeSign(value string) string {
	switch {
	case strings.Contains(value, "+"):
		return strings.ReplaceAll(value, "+", "-")
	default:
		return strings.ReplaceAll(value, "-", "+")
	}
}

// Separates variables of different degree:
//
// Example: input ==> [x x^3 x^2]
//
// output ==> [x], [x^2], [x^3]
func separatePowers(variablesList []string) (firstDegVar, secondDegVar, thirdDegVar []string) {
	for _, variable := range variablesList {
		switch {
		case strings.Contains(variable, "x^3"):
			thirdDegVar = append(thirdDegVar, variable)
		case strings.Contains(variable, "x^2"):
			secondDegVar = append(secondDegVar, variable)
		case strings.Contains(variable, "x"):
			firstDegVar = append(firstDegVar, variable)
		}
	}
	return
}

// Sums all the variables value that are powered at the same exponential value.
//
// Example:
//
// input ==> [x +3x -5x]
//
// output ==> -x
func sumVariableValues(variablesList []string) (totalCoeff float64) {
	coefficients := []float64{}
	for _, ax := range variablesList {
		if ax == "x" {
			ax = "1x"
		} else if ax == "-x" {
			ax = "-1x"
		}
		coeff, _ := strconv.ParseFloat(strings.Split(ax, "x")[0], 64)
		coefficients = append(coefficients, coeff)
	}
	for _, coeff := range coefficients {
		totalCoeff += coeff
	}
	return
}

// Sums all the constants values
func sumConstantValues(constantsList []string) (total float64) {
	for _, num := range constantsList {
		value, _ := strconv.ParseFloat(num, 64)
		total += value
	}
	return
}

// Create a Polynom of type: ax^3+bx^2+cx+d=0
func createSamplePolynom(a, b, c, d float64) Polynom {
	return Polynom{
		a: a,
		b: b,
		c: c,
		d: d,
	}
}

func evaluatePolynomDeg(polynom Polynom) (solution EquationSolution) {
	switch {
	case polynom.a != 0:
		x0, x1, x2 := solveCubicEquation(polynom)
		return EquationSolution{
			complexSolutions: []complex128{x0, x1, x2},
		}
	case polynom.b != 0:
		x1, x2 := solveQuadraticEquation(polynom)
		return EquationSolution{
			realSolutions: []float64{x1, x2},
		}
	case polynom.c != 0:
		return EquationSolution{
			realSolutions: []float64{solveLinearEquation(polynom)},
		}
	default:
		return EquationSolution{
			realSolutions:    []float64{0},
			complexSolutions: []complex128{0},
			errorDescription: "Missing variable error.",
		}
	}
}

// It solves a linear equation of type ax+b=0.
func solveLinearEquation(polynom Polynom) (result float64) {
	return (polynom.d / polynom.c) * -1
}

// It solves a quadratic equation of type ax^2+bx+c=0, as long as delta is not negative.
func solveQuadraticEquation(polynom Polynom) (x1, x2 float64) {
	a := polynom.b
	b := polynom.c
	c := polynom.d
	delta := b*b - 4*a*c
	x1 = (-b - math.Sqrt(delta)) / (2 * a)
	x2 = (-b + math.Sqrt(delta)) / (2 * a)
	return
}

// It solves a cubic equation of type ax^3+bx^2+cx+d=0 using the General Cubic formula
//
// This is how it solve it https://en.wikipedia.org/wiki/Cubic_equation
func solveCubicEquation(polynom Polynom) (x0, x1, x2 complex128) {
	a := polynom.a
	b := polynom.b
	c := polynom.c
	d := polynom.d
	deltaZero := b*b - 3*a*c
	deltaOne := 2*b*b*b - 9*a*b*c + 27*a*a*d
	deltaDifference := deltaOne*deltaOne - 4*deltaZero*deltaZero*deltaZero
	C := math.Cbrt(deltaOne + math.Sqrt(deltaDifference)/2)
	epsilon := complex(-1/2, 1.73/2)
	epsilon2 := epsilon * epsilon
	if C == 0 {
		x0 = complex((-1/(3*a))*(b+C), 0)
		x1 = complex(-1/(3*a), 0) * (complex(b, 0) + complex(C, 0)*epsilon)
		x2 = complex(-1/(3*a), 0) * (complex(b, 0) + complex(C, 0)*epsilon2)
	} else {
		x0 = complex((-1/(3*a))*(b+C+(deltaZero/C)), 0)
		x1 = complex(-1/(3*a), 0) * (complex(b, 0) + complex(C, 0)*epsilon + complex(deltaZero, 0)/(complex(C, 0)*epsilon))
		x2 = complex(-1/(3*a), 0) * (complex(b, 0) + complex(C, 0)*epsilon2 + complex(deltaZero, 0)/(complex(C, 0)*epsilon2))
	}
	return
}

// It converts complex numbers to real ones if complex coefficient is zero.
func complexToReal(solution EquationSolution) EquationSolution {
	for i, complexSol := range solution.complexSolutions {
		if imag(complexSol) == 0 {
			solution.realSolutions = append(solution.realSolutions, real(complexSol))
			solution.complexSolutions = append(solution.complexSolutions[:i], solution.complexSolutions[i+1:]...)
		}
	}
	return solution
}
