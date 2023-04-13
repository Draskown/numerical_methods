package lr3

func Adams() ([]float64, []float64) {
	// Calculate the number of steps
	n := int((X0max - X0min) / H)

	// Initialize arrays to hold x and y values
	x := make([]float64, n+1)
	y := make([]float64, n+1)

	// Use the Runge-Kutta method to compute the first four values of y
	xRK, yRK := RungeKutte()

	for i := 0; i < 4; i++ {
		x[i] = xRK[i]
		y[i] = yRK[i]
	}

	// Use the Adams-Bashforth method to compute the remaining values of y
	for i := 4; i < n+1; i++ {
		y[i] = y[i-1] + H*(55*F(x[i-1], y[i-1])-59*F(x[i-2], y[i-2])+37*F(x[i-3], y[i-3])-9*F(x[i-4], y[i-4]))/24
		x[i] = X0min + float64(i)*H
	}

	return x, y
}

func AdamsSystem() ([]float64, []float64, []float64) {
	// Calculate the number of steps
	n := int((X0SystemMax - X0SystemMin) / H)

	// Initialize arrays to hold x and y values
	x := make([]float64, n+1)
	y := make([]float64, n+1)
	z := make([]float64, n+1)

	// Use the Runge-Kutta method to compute the first four values of y and z
	xRK, yRK, zRK := RungeKutteSystem()

	for i := 0; i < 4; i++ {
		x[i] = xRK[i]
		y[i] = yRK[i]
		z[i] = zRK[i]
	}

	// Use the Adams-Bashforth method to compute the remaining values of y and z
	for i := 4; i < n+1; i++ {
		// Predictor
		yPred := y[i-1] + (H/24)*(55*FSystemY(x[i-1], y[i-1], z[i-1])-59*FSystemY(x[i-2], y[i-2], z[i-2])+37*FSystemY(x[i-3], y[i-3], z[i-3])-9*FSystemY(x[i-4], y[i-4], z[i-4]))
		zPred := z[i-1] + (H/24)*(55*FSystemZ(y[i-1], z[i-1])-59*FSystemZ(y[i-2], z[i-2])+37*FSystemZ(y[i-3], z[i-3])-9*FSystemZ(y[i-4], z[i-4]))
		// Corrector
		y[i] = y[i-1] + (H/24)*(9*FSystemY(x[i], yPred, zPred)+19*FSystemY(x[i-1], y[i-1], z[i-1])-5*FSystemY(x[i-2], y[i-2], z[i-2])+FSystemY(x[i-3], y[i-3], z[i-3]))
		z[i] = z[i-1] + (H/24)*(9*FSystemZ(y[i], zPred)-19*FSystemZ(y[i-1], z[i-1])+5*FSystemZ(y[i-2], z[i-2])+FSystemZ(y[i-3], z[i-3]))
		x[i] = X0SystemMin + float64(i)*H
	}

	return x, y, z
}