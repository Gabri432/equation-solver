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
