package equationsolver

import (
	"testing"
)

func TestValidateEquation(t *testing.T) {
	errorDescription := ValidateEquation("x^2+3")
	if errorDescription == "" {
		t.Fatal("Expected function to return an error message when equal is missing.")
	}
	errorDescription = ValidateEquation("x^^2+3")
	if errorDescription == "" {
		t.Fatal("Expected function to return an error message when consecutives powering signs were present.")
	}
	errorDescription = ValidateEquation("x**2+3")
	if errorDescription == "" {
		t.Fatal("Expected function to return an error message when consecutives multiplication signs were present.")
	}
	errorDescription = ValidateEquation("x//2+3")
	if errorDescription == "" {
		t.Fatal("Expected function to return an error message when consecutives division signs were present.")
	}
	errorDescription = ValidateEquation("xx2+3")
	if errorDescription == "" {
		t.Fatal("Expected function to return an error message when consecutives variable signs were present.")
	}
}

func TestEvaluateEquation(t *testing.T) {
	eqVariables, eqConstants := splitEquation(replaceEquation("x^2+3=1"))
	firstDegVar, secondDegVar, thirdDegVar := separatePowers(eqVariables)
	xElevated1 := sumVariableValues(firstDegVar)
	xElevated2 := sumVariableValues(secondDegVar)
	xElevated3 := sumVariableValues(thirdDegVar)
	constantsSum := sumConstantValues(eqConstants)
	if len(eqVariables) != 1 || len(eqConstants) != 2 {
		t.Fatalf("Expected to have 1 variable and two costants, got %d var, %d const.", len(eqVariables), len(eqConstants))
	}
	if xElevated2 != 1 || xElevated3 != 0 || xElevated1 != 0 {
		t.Fatalf("Expected 0 1 0, got %f %f %f.", xElevated3, xElevated2, xElevated1)
	}
	if constantsSum != 2 {
		t.Fatalf("Expected total sum of constants to be equal to 2, got %f", constantsSum)
	}
	solution := EvaluateEquation("x^2+3=1")
	if len(solution.complexSolutions) != 2 {
		t.Fatalf("Expected to have 2 complex solutions, got %d.", len(solution.complexSolutions))
	}
}
