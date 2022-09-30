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
