package equationsolver

import (
	"math"
	"strconv"
	"strings"
)

type Polynom struct {
	thirdDegVarCoefficient  float64
	secondDegVarCoefficient float64
	firstDegVarCoefficient  float64
	constant                float64
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

// Sums all the variables value of a specific power.
//
// Example:
//
// input ==> [x +3x -5x], 1
//
// output ==> -x
func sumVariableValues(variablesList []string, power int) (totalCoeff float64) {
	coefficients := []float64{}
	for _, ax := range variablesList {
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
func createSamplePolynom(firstDegVarCoefficient, secondDegVarCoefficient, thirdDegVarCoefficient, constant float64) Polynom {
	return Polynom{
		firstDegVarCoefficient:  firstDegVarCoefficient,
		secondDegVarCoefficient: secondDegVarCoefficient,
		thirdDegVarCoefficient:  thirdDegVarCoefficient,
		constant:                constant,
	}
}

func evaluatePolynomDeg(polynom Polynom) {
	switch {
	case polynom.thirdDegVarCoefficient != 0:
		solveCubicEquation(polynom)
	case polynom.secondDegVarCoefficient != 0:
		solveQuadraticEquation(polynom)
	case polynom.firstDegVarCoefficient != 0:
		solveLinearEquation(polynom)
	}
}

// It solves a linear equation of type ax+b=0.
func solveLinearEquation(polynom Polynom) (result float64) {
	return (polynom.constant / polynom.firstDegVarCoefficient) * -1
}

// It solves a quadratic equation of type ax^2+bx+c=0.
func solveQuadraticEquation(polynom Polynom) (x1, x2 float64) {
	a := polynom.secondDegVarCoefficient
	b := polynom.firstDegVarCoefficient
	c := polynom.constant
	delta := b*b - 4*a*c
	x1 = (b - math.Sqrt(delta)) / (2 * a)
	x2 = (b + math.Sqrt(delta)) / (2 * a)
	return
}

// It solves a cubic equation of type ax^3+bx^2+cx+d=0 using the General Cubic formula
//
// This is how it solve it https://en.wikipedia.org/wiki/Cubic_equation
func solveCubicEquation(polynom Polynom) (x0, x1, x2 string) {
	a := polynom.thirdDegVarCoefficient
	b := polynom.secondDegVarCoefficient
	c := polynom.firstDegVarCoefficient
	d := polynom.constant
	deltaZero := b*b - 3*a*c
	deltaOne := 2*b*b*b - 9*a*b*c + 27*a*a*d
	C := math.Cbrt(deltaOne + math.Sqrt(deltaOne*deltaOne-4*deltaZero*deltaZero*deltaZero)/2)
	// e := "-1/2 + (1.73/2)i" // It should be [-1 + squareRoot(-3)]/2 == -1/2 + (1.73/2)i
	// e2 := "17/324 + (4/9)i" // e^2
	x0 = strconv.FormatFloat(-(1/3*a)*(b+C+deltaZero/C), 'f', 3, 64)
	x1 = strconv.FormatFloat(-(1/3*a)*b, 'f', 3, 64) + epsilonPoweredOneMultiplyingC(a, C, deltaZero)
	x2 = strconv.FormatFloat(-(1/3*a)*b, 'f', 3, 64) + epsilonPoweredTwoMultiplyingC(a, C, deltaZero)
	return
}

// C * -1/3a * 17/324 + [C * -1/3a * 4/9]i

// It calculates e*C + deltaZero/e*C, where e = -1/2 + (1.73/2)i, C is explained here https://en.wikipedia.org/wiki/Cubic_equation
func epsilonPoweredOneMultiplyingC(a, C, deltaZero float64) (solution string) {
	y1 := strconv.FormatFloat(((-1/3*a)*C)*(-1/2)*deltaZero, 'f', 3, 64)
	y2 := strconv.FormatFloat(((-1/3*a)*C)*(1.73/2)*deltaZero, 'f', 3, 64) + "i"
	y3 := strconv.FormatFloat(deltaZero, 'f', 3, 64)
	solution = y1 + " " + y2 + " " + "( " + y3 + "/ (" + y1 + " " + y2 + "))"
	return
}

// It calculates (e^2)*C + deltaZero/(e^2)*C, where e^2 = 17/324 + (4/9)i, C is explained here https://en.wikipedia.org/wiki/Cubic_equation
func epsilonPoweredTwoMultiplyingC(a, C, deltaZero float64) (solution string) {
	y1 := strconv.FormatFloat(((-1/3*a)*C)*(17/324)*deltaZero, 'f', 3, 64)
	y2 := strconv.FormatFloat(((-1/3*a)*C)*(4/9)*deltaZero, 'f', 3, 64) + "i"
	y3 := strconv.FormatFloat(deltaZero, 'f', 3, 64)
	solution = y1 + " " + y2 + " " + "( " + y3 + "/ (" + y1 + " " + y2 + "))"
	return
}

/*
x^2+2 = 0
x^2/+2/=/0
*/
