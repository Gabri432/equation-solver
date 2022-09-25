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
