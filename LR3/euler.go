package lr3

// Solves a differential equation using Euler's formula
func Euler() ([]float64, []float64) {
	// Define the initial conditions
	y0 := 0.4

	// Define the step size and the number of steps
	h := 0.1

	x := X0min
	y := y0

	// Define the arrays of the output data
	var xSlice []float64
	var ySlice []float64

	// Add the first values
	xSlice = append(xSlice, x)
	ySlice = append(ySlice, y)

	// Solve the equation using Euler;s formula
	for i := 0.0; i < X0max-h; i += h {
		y += h * F(x, y)
		x += h

		xSlice = append(xSlice, x)
		ySlice = append(ySlice, y)
	}

	return xSlice, ySlice
}