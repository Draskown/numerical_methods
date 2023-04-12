package lr3

// Solves a Cauchy problem of the equation using Euler's formula
func Euler() ([]float64, []float64) {
	// Calculate the number of steps
    n := int((X0max - X0min) / H)

	// Define the arrays of the output data
	x := make([]float64, n+1)
	y := make([]float64, n+1)

	// Add the first values
	x[0] = X0SystemMin
	y[0] = Y0

	// Solve the equation using Euler's formula
	for i := 0; i < n; i++ {
		y[i+1] = y[i] + H * F(x[i], y[i])
		x[i+1] = x[i]+H
	}

	return x, y
}

// Solves a Cauchy problem of the system of equations using Euler's formula
func EulerSystem() ([]float64, []float64, []float64){
	n := int((X0SystemMax - X0SystemMin) / H)

	// Define the arrays of the output data
	x := make([]float64, n+1)
	y := make([]float64, n+1)
	z := make([]float64, n+1)

	// Add the first values
	x[0] = X0SystemMin
	y[0] = Y0System
	z[0] = Z0System

    // Iterate using Euler's method
    for i := 0; i < n; i++ {
        // Calculate the next values of y and z
        y[i+1] = y[i] + H*FSystemY(x[i], y[i], z[i])
        z[i+1] = z[i] + H*FSystemZ(y[i], z[i])
		x[i+1] = x[i] + H
    }

	return x, y, z
}