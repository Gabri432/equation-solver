package equationsolver

import (
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
func createSamplePolynom(firstDegVarCoefficient, secondDegVarCoefficient, thirdDegVarCoefficient, constant float64) Polynom

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

func solveLinearEquation(polynom Polynom) {}
func solveQuadraticEquation(polynom Polynom)
func solveCubicEquation(polynom Polynom)

/*
x^2+2 = 0
x^2/+2/=/0
*/
