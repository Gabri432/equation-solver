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
	solution := EvaluateEquation("x^2+3=1")
	if len(solution.complexSolutions) != 2 {
		t.Fatalf("Expected to have 2 complex solutions, got %d.", len(solution.complexSolutions))
	}
	x1 := imag(solution.complexSolutions[0])
	x2 := imag(solution.complexSolutions[0])
	p := func(x float64) float64 {
		return 1*x*x + 3
	}
	if p(x1)+p(x2)*-1 != 0 {
		t.Fatalf("The two results aren't completely or at all correct: %f, %f", p(x1), p(x2))
	}
	solution2 := EvaluateEquation("+3=1-x^2")
	y1 := imag(solution2.complexSolutions[0])
	y2 := imag(solution2.complexSolutions[0])
	sum1 := p(x1) + p(x2)*-1
	sum2 := p(y1) + p(y2)*-1
	if sum1 != sum2 {
		t.Fatalf("Expected '+3=1-x^2' and 'x^2+3=1' to have the same result: %f, %f", sum1, sum2)
	}
}

func TestEvaluateEquation2(t *testing.T) {
	solution1 := EvaluateEquation("2x^3+x^2=-4-x")
	solution2 := EvaluateEquation("2x^3+x^2+x+4=0")
	if len(solution1.complexSolutions) != len(solution2.complexSolutions) {
		t.Fatalf("Expected same complex solutions amount for both equations, got %d and %d.", len(solution1.complexSolutions), len(solution2.complexSolutions))
	}
	if len(solution1.realSolutions) != len(solution2.realSolutions) {
		t.Fatalf("Expected real solutions amount for both equations, got %d and %d.", len(solution1.realSolutions), len(solution2.realSolutions))
	}
	if solution1.realSolutions[0] != solution2.realSolutions[0] {
		t.Fatalf("Expected same real solution for both equations, got %f and %f", solution1.realSolutions[0], solution2.realSolutions[0])
	}
}
