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
	if changeSign("-2x") != "+2x" {
		t.Fatalf("Expected '+2x' as result, got %s", changeSign("-2x"))
	}
	if changeSign("+3x") != "-3x" {
		t.Fatalf("Expected '-3x' as result, got %s", changeSign("+3x"))
	}
	if changeSign("x") != "x" {
		t.Fatalf("Expected 'x' as result, got %s", changeSign("x"))
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
	if sumVariableValues([]string{"x", "4x", "-2x"}) != 3 {
		t.Fatalf("Expected to have 3 as result, got %f", sumVariableValues([]string{"x", "4x", "-2x"}))
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
	x1, x2 := solveQuadraticEquation(myPolynom)
	if x1 > x2 {
		t.Fatal("Expected x2 to be bigger than x1, but it is not.")
	}
	if x1 < -1.8 || x1 > -1.6 {
		t.Fatalf("Expected to have 2 as result, got %f", solveLinearEquation(myPolynom))
	}
	if x2 < -0.3 || x1 > 0 {
		t.Fatalf("Expected to have 2 as result, got %f", solveLinearEquation(myPolynom))
	}

}
