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
