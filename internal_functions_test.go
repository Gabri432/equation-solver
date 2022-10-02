package equationsolver

import (
	"testing"
)

var testPolynom = "x^3+2x^2+x+1=0"

func TestIsSign(t *testing.T) {
	mapOfBools := map[bool]int{}
	mapOfBools[isSign("+")]++
	mapOfBools[isSign("-")]++
	mapOfBools[isSign("=")]++
	mapOfBools[isSign("/")]++
	if mapOfBools[false] != 1 {
		t.Fatalf("Expected determinant should be one, got %d.", mapOfBools[false])
	}
	if mapOfBools[true] != 3 {
		t.Fatalf("Expected determinant should be one, got %d.", mapOfBools[true])
	}
}

func TestReplaceEquation(t *testing.T) {
	newEquation := replaceEquation(testPolynom)
	if newEquation != "x^3;+2x^2;+x;+1;=;0" {
		t.Fatalf("Expected x^3;+2x^2;+x;+1;=;0, got %s", newEquation)
	}

}

func TestSplitEquation(t *testing.T) {
	variables, constants := splitEquation(replaceEquation(testPolynom))
	if len(variables) != 3 {
		t.Fatalf("Expected to have 3 variables, got %d from: %s", len(variables), variables)
	}
	if len(constants) != 2 {
		t.Fatalf("Expected to have 2 constants, got %d from: %s", len(constants), constants)
	}

}

func TestChangeSign(t *testing.T) {
	if changeSign("-2x", true) != "+2x" {
		t.Fatalf("Expected '+2x' as result, got %s", changeSign("-2x", true))
	}
	if changeSign("+3x", true) != "-3x" {
		t.Fatalf("Expected '-3x' as result, got %s", changeSign("+3x", true))
	}
	if changeSign("x", true) != "-x" {
		t.Fatalf("Expected '-x' as result, got %s", changeSign("x", true))
	}
	if changeSign("1", true) != "-1" {
		t.Fatalf("Expected '-1' as result, got %s", changeSign("1", true))
	}
}

func TestSeparatePowers(t *testing.T) {
	x3, x2, x1 := separatePowers([]string{"x^3", "x^2", "x"})
	if len(x3) != 1 {
		t.Fatalf("Expected to have 1 element at third degree, got %d from: %s", len(x3), x3)
	}
	if len(x2) != 1 {
		t.Fatalf("Expected to have 1 element at second degree, got %d from: %s", len(x2), x2)
	}
	if len(x1) != 1 {
		t.Fatalf("Expected to have 1 element at first degree, got %d from: %s", len(x1), x1)
	}

}

func TestSumVariableValues(t *testing.T) {
	sum := sumVariableValues([]string{"x", "4x", "-2x"})
	if sum != 3 {
		t.Fatalf("Expected to have 3 as result, got %f", sum)
	}
	sum = sumVariableValues([]string{"x^2", "4x^2", "-2x^2"})
	if sum != 3 {
		t.Fatalf("Expected to have 3 as result, got %f", sum)
	}

}

func TestSumConstantValues(t *testing.T) {
	if sumConstantValues([]string{"1", "4", "-2"}) != 3 {
		t.Fatalf("Expected to have 3 as result, got %f", sumConstantValues([]string{"1", "4", "-2"}))
	}

}

func TestSolveLinearEquation(t *testing.T) {
	myPolynom := createSamplePolynom(0, 0, 2, 4)
	if solveLinearEquation(myPolynom) != -2 {
		t.Fatalf("Expected to have 2 as result, got %f", solveLinearEquation(myPolynom))
	}

}

func TestSolveQuadraticEquation(t *testing.T) {
	myPolynom := createSamplePolynom(0, 2, 4, 1)
	solution := complexToReal(solveQuadraticEquation(myPolynom))
	x1 := solution.realSolutions[0]
	x2 := solution.realSolutions[1]
	if len(solution.realSolutions) != 2 {
		t.Fatalf("Expected to have 2 real solutions, got %d", len(solution.realSolutions))
	}
	if len(solution.complexSolutions) > 0 {
		t.Fatalf("Expected to have 0 complex solutions, got %d ", len(solution.complexSolutions))
	}
	p := func(x float64) float64 {
		return 2*x*x + 4*x + 1
	}
	if (p(x1) > 0.5 || p(x1) < -0.5) || (p(x2) > 0.5 || p(x2) < -0.5) {
		t.Fatalf("The two results aren't completely or at all correct: %f, %f", p(x1), p(x2))
	}

}

func TestSolveCubicEquation(t *testing.T) {
	myPolynom := createSamplePolynom(1, -3, 3, -1)
	x0, x1, x2 := solveCubicEquation(myPolynom)
	if real(x0) != 1 {
		t.Fatalf("Expected x0 == 1, got %f", real(x0))
	}
	if real(x1) != 1 {
		t.Fatalf("Expected x1 == 1, got %f", real(x1))
	}
	if real(x2) != 1 {
		t.Fatalf("Expected x2 == 1, got %f", real(x2))
	}

}

func TestComplexToReal(t *testing.T) {
	mySolution := EquationSolution{
		realSolutions:    []float64{1, 2, 3},
		complexSolutions: []complex128{complex(1, 0), complex(2, 2)},
	}
	updatedSolution := complexToReal(mySolution)
	if len(updatedSolution.complexSolutions) > 1 {
		t.Fatalf("Expected to have 1 complex solution, got %d", len(updatedSolution.complexSolutions))
	}
	if len(updatedSolution.realSolutions) < 4 {
		t.Fatalf("Expected to have 4 real solutions, got %d", len(updatedSolution.realSolutions))
	}
}

func TestDepressedCubic(t *testing.T) {
	polynom := Polynom{a: 1, b: -5, c: -4, d: 1}
	a := polynom.a
	b := polynom.b
	c := polynom.c
	d := polynom.d
	x1, x2, x3 := depressedCubic(polynom)
	errorMargin := 0.25
	eq := func(polynom Polynom, x float64) float64 {
		return a*(x*x*x) + b*(x*x) + c*x + d
	}
	eq1 := eq(polynom, x1)
	eq2 := eq(polynom, x2)
	eq3 := eq(polynom, x3)
	if eq1 > errorMargin || eq1 < -errorMargin {
		t.Fatalf("Expected eq(polynom, x1) to be really close to zero, got %f", eq1)
	}
	if eq2 > errorMargin || eq2 < -errorMargin {
		t.Fatalf("Expected eq(polynom, x2) to be really close to zero, got %f", eq2)
	}
	if eq3 > errorMargin || eq3 < -errorMargin {
		t.Fatalf("Expected eq(polynom, x3) to be really close to zero, got %f", eq3)
	}
}
