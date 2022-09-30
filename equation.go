/*
The 'equationsolver' package allows to solve linear, Quadratic and Cubic equations.

How to write an equation:

1) 'x' is the only variable you can use;

2) '^' is the symbol to indicate the powering, ex: x^3 is x powered 5.

Check its correctness with this link: https://www.calculatorsoup.com/calculators/algebra/cubicequation.php
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
func ValidateEquation(equation string) (errorMessage string) {
	hasEqualSign := strings.Contains(equation, "=")
	hasDoubleMultiplySign := strings.Contains(equation, "**")
	hasDoublePowerSign := strings.Contains(equation, "^^")
	if !hasEqualSign {
		return "Error: No equal sign detected."
	} else if hasDoubleMultiplySign {
		return "Error: Double multiply sign detected."
	} else if hasDoublePowerSign {
		return "Error: Double power sign detected."
	}
	return ""
}

func EvaluateEquation(equation string) {
	if errorMessage := ValidateEquation(equation); errorMessage != "" {
		fmt.Println(errorMessage)
		return
	}
}
