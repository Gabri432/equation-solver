package equationsolver

// Creates a polynom in the form: ax^3+bx^2+cx+d
type Polynom struct {
	A float64
	B float64
	C float64
	D float64
}

type EquationSolution struct {
	RealSolutions    []float64    // set of real solutions
	ComplexSolutions []complex128 // set of complex solutions
	ErrorDescription string       // error message
}

const (
	NO_EQUAL_SIGN_ERROR        = "Error: No equal sign detected."
	DOUBLE_MULTIPLY_SIGN_ERROR = "Error: Double multiply sign detected."
	DOUBLE_POWER_SIGN_ERROR    = "Error: Double power sign detected."
	DOUBLE_DIVISION_SIGN_ERROR = "Error: Double division sign detected."
	DOUBLE_VARIABLE_SIGN_ERROR = "Error: Double variable sign detected."
)
