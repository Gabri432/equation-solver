package equationsolver

import (
	"math"
	"strconv"
	"strings"
)

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
			value = changeSign(value, true)
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
func changeSign(value string, flag bool) string {
	switch {
	case strings.Contains(value, "+") && flag:
		return strings.ReplaceAll(value, "+", "-")
	case strings.Contains(value, "-") && flag:
		return strings.ReplaceAll(value, "-", "+")
	case !strings.Contains(value, "+") && flag:
		return "-" + value
	case !strings.Contains(value, "-") && flag:
		return "+" + value
	}
	return ""
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
		switch ax {
		case "x", "+x":
			ax = "1x"
		case "-x":
			ax = "-1x"
		case "x^2", "+x^2":
			ax = "1x^2"
		case "x^3", "+x^3":
			ax = "1x^3"
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
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func evaluatePolynomDeg(polynom Polynom) (solution EquationSolution) {
	switch {
	case polynom.A != 0:
		x0, x1, x2 := solveCubicEquation(polynom)
		return EquationSolution{
			ComplexSolutions: []complex128{x0, x1, x2},
		}
	case polynom.B != 0:
		return solveQuadraticEquation(polynom)
	case polynom.C != 0:
		return EquationSolution{
			RealSolutions: []float64{solveLinearEquation(polynom)},
		}
	default:
		return EquationSolution{
			RealSolutions:    []float64{0},
			ComplexSolutions: []complex128{0},
			ErrorDescription: "Missing variable error.",
		}
	}
}

// It solves a linear equation of type ax+b=0.
func solveLinearEquation(polynom Polynom) (result float64) {
	return (polynom.D / polynom.C) * -1
}

// It solves a quadratic equation of type ax^2+bx+c=0.
func solveQuadraticEquation(polynom Polynom) (solution EquationSolution) {
	a := polynom.B
	b := polynom.C
	c := polynom.D
	delta := b*b - 4*a*c
	if delta < 0 {
		deltaToUse := math.Sqrt(math.Abs(delta))
		x1 := complex((-b / (2 * a)), -deltaToUse/(2*a))
		x2 := complex((-b / (2 * a)), deltaToUse/(2*a))
		solution = EquationSolution{
			ComplexSolutions: []complex128{x1, x2},
		}
		return
	}
	x1 := (-b - math.Sqrt(delta)) / (2 * a)
	x2 := (-b + math.Sqrt(delta)) / (2 * a)
	solution = EquationSolution{
		RealSolutions: []float64{x1, x2},
	}
	return
}

// It solves a cubic equation of type ax^3+bx^2+cx+d=0
//
// More details on https://en.wikipedia.org/wiki/Cubic_equation and https://proofwiki.org/wiki/Cardano's_Formula
func solveCubicEquation(polynom Polynom) (x0, x1, x2 complex128) {
	a := polynom.A
	b := polynom.B
	c := polynom.C
	d := polynom.D
	if d == 0 {
		solution := solveQuadraticEquation(Polynom{A: b, B: c, C: d})
		if len(solution.ComplexSolutions) == 0 {
			x1 = complex(solution.RealSolutions[0], 0)
			x2 = complex(solution.RealSolutions[1], 0)
		} else if len(solution.RealSolutions) == 0 {
			x1 = solution.ComplexSolutions[0]
			x2 = solution.ComplexSolutions[1]
		}
		return
	}
	q := (3*a*c - b*b) / (9 * a * a)
	r := (-2*b*b*b + 9*a*b*c - 27*a*a*d) / (54 * a * a * a)
	if (q*q*q)+(r*r) < 0 {
		sol1, sol2, sol3 := depressedCubic(polynom)
		x0 = complex(sol1, 0)
		x1 = complex(sol2, 0)
		x2 = complex(sol3, 0)
		return
	}
	difference := math.Sqrt((q * q * q) + (r * r))
	s := math.Cbrt(r + difference)
	t := math.Cbrt(r - difference)
	x0 = complex(s+t-(b/3*a), 0)
	x1 = complex(-(s+t)/2-b/(3*a), (s-t)*(1.73/2))
	x2 = complex(-(s+t)/2-b/(3*a), -(s-t)*(1.73/2))
	return
}

// It finds the three real solutions of a general cubic equation using a depressed one.
//
// More details on: https://en.wikipedia.org/wiki/Cubic_equation#Depressed_cubic
func depressedCubic(polynom Polynom) (x1, x2, x3 float64) {
	a := polynom.A
	b := polynom.B
	c := polynom.C
	d := polynom.D
	p := (3*a*c - b*b) / (3 * a * a)
	q := (2*b*b*b - 9*a*b*c + 27*a*a*d) / (27 * a * a * a)
	t := func(k int) float64 { // https://proofwiki.org/wiki/Cardano%27s_Formula/Trigonometric_Form
		arccosineArgument := ((3 * q) / (2 * p)) * math.Sqrt(-3/p)
		cosineArgument := (math.Acos(arccosineArgument) - 2*3.14159265*float64(k)) / 3
		return 2 * math.Sqrt(-p/3) * math.Cos(cosineArgument)
	}
	x1 = t(0) - (b / (3 * a))
	x2 = t(1) - (b / (3 * a))
	x3 = t(2) - (b / (3 * a))
	return
}

// It converts complex numbers to real ones if complex coefficient is zero.
func complexToReal(solution EquationSolution) EquationSolution {
	for i, complexSol := range solution.ComplexSolutions {
		if imag(complexSol) == 0 {
			solution.RealSolutions = append(solution.RealSolutions, real(complexSol))
			if i == 0 && i+1 < len(solution.ComplexSolutions) {
				solution.ComplexSolutions = append(solution.ComplexSolutions[:i], solution.ComplexSolutions[i+1:]...)
			} else if i == 1 {
				solution.ComplexSolutions = append(solution.ComplexSolutions[:i], solution.ComplexSolutions[i+1:]...)
			} else {
				solution.ComplexSolutions = []complex128{}
			}
		}
	}
	return solution
}

// Returns polynom solutions
func (p Polynom) SolvePolynom() EquationSolution {
	return complexToReal(evaluatePolynomDeg(p))
}
