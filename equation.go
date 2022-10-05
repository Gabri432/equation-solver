/*
The 'equationsolver' package allows to solve Linear, Quadratic and Cubic equations.

How to write an equation:

1) 'x' is the only variable you can use;

2) '^' is the symbol to indicate the powering, ex: x^3 is x powered 3.

Example:

x^3+x^2+3=3+x-2x^2 (valid)

y^2=x**2 (invalid)

If you need to report an issue go on https://github.com/Gabri432/equation-solver/issues/new

If you appreciate the work consider putting a star on https://github.com/Gabri432/equation-solver

Check its correctness with this link: https://www.calculatorsoup.com/calculators/algebra/cubicequation.php

This project is under the MIT license: https://github.com/Gabri432/equation-solver/blob/master/license
*/
package equationsolver

import (
	"fmt"
	"strings"
)

// Check if the expression given is an actual equation.
//
// It looks for:
//
// - the presence of the '=' equal sign
//
// - the absence of '**' (consecutives multiplication signs)
//
// - the absence of '^^' (consecutives powering signs)
//
// - the absence of '//' (consecutives division signs)
//
// - the absence of 'xx' (consecutives variable signs)
func ValidateEquation(equation string) (errorMessage string) {
	hasEqualSign := strings.Contains(equation, "=")
	hasDoubleMultiplySign := strings.Contains(equation, "**")
	hasDoublePowerSign := strings.Contains(equation, "^^")
	hasDoubleDivisionSign := strings.Contains(equation, "//")
	hasDoubleVariableSign := strings.Contains(equation, "xx")
	switch {
	case !hasEqualSign:
		return NO_EQUAL_SIGN_ERROR
	case hasDoubleMultiplySign:
		return DOUBLE_MULTIPLY_SIGN_ERROR
	case hasDoublePowerSign:
		return DOUBLE_POWER_SIGN_ERROR
	case hasDoubleDivisionSign:
		return DOUBLE_DIVISION_SIGN_ERROR
	case hasDoubleVariableSign:
		return DOUBLE_VARIABLE_SIGN_ERROR
	}
	return ""
}

// It takes the user equation and solves it.
//
// It can only take Linear, Quadratic and Cubic equations.
//
// It returns the set of real solutions, complex solution, and eventually an error message.
func EvaluateEquation(equation string) EquationSolution {
	if errorMessage := ValidateEquation(equation); errorMessage != "" {
		fmt.Println(errorMessage)
		return EquationSolution{ErrorDescription: errorMessage}
	}
	eqVariables, eqConstants := splitEquation(replaceEquation(equation))  // Separating variables and constants
	firstDegVar, secondDegVar, thirdDegVar := separatePowers(eqVariables) // Separating variables of different power
	xElevated1 := sumVariableValues(firstDegVar)                          // Sum of variables powered 1
	xElevated2 := sumVariableValues(secondDegVar)                         // Sum of variables powered 2
	xElevated3 := sumVariableValues(thirdDegVar)                          // Sum of variables powered 3
	constantsSum := sumConstantValues(eqConstants)                        // Sum of constants
	polynom := createSamplePolynom(xElevated3, xElevated2, xElevated1, constantsSum)
	solution := complexToReal(evaluatePolynomDeg(polynom))
	return solution
}
