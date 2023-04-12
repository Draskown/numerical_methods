package lr3

// Solves a Cauchy problem of the equation using Runge-Kutte's formula
func RungeKutte() ([]float64, []float64) {	
	// Calculate the number of steps
    n := int((X0max - X0min) / H)

    // Initialize arrays to hold x and y values
    x := make([]float64, n+1)
    y := make([]float64, n+1)

    // Set initial values
    x[0] = X0min
    y[0] = Y0

    // Iterate over the number of steps and calculate the next value of y
    for i := 0; i < n; i++ {
        k1 := H * F(x[i], y[i])
        k2 := H * F(x[i]+H/2, y[i]+k1/2)
        k3 := H * F(x[i]+H/2, y[i]+k2/2)
        k4 := H * F(x[i]+H, y[i]+k3)
        y[i+1] = y[i] + (k1+2*k2+2*k3+k4)/6
        x[i+1] = x[i] + H
    }

    return x, y
}

// Solves a Cauchy problem of the system of equations using Runge-Kutte's formula
func RungeKutteSystem() ([]float64, []float64, []float64){
	// Calculate the number of steps
	n := int((X0SystemMax - X0SystemMin) / H)

	// Initialize arrays to hold x, y and z values
	x := make([]float64, n+1)
	y := make([]float64, n+1)
	z := make([]float64, n+1)

	// Set initial values
	x[0] = X0SystemMin
	y[0] = Y0System
	z[0] = Z0System

	// Iterate over the number of steps and calculate the next value of y
    for i := 0; i < n; i++ {
        k1y := H * FSystemY(x[i], y[i], z[i])
		k1z := H * FSystemZ(y[i], z[i])
		k2y := H * FSystemY(x[i]+H/2, y[i]+k1y/2, z[i]+k1z/2)
		k2z := H * FSystemZ(y[i]+k1y/2, z[i]+k1z/2)
		k3y := H * FSystemY(x[i]+H/2, y[i]+k2y/2, z[i]+k2z/2)
		k3z := H * FSystemZ(y[i]+k2y/2, z[i]+k2z/2)
		k4y := H * FSystemY(x[i]+H, y[i]+k3y, z[i]+k3z)
		k4z := H * FSystemZ(y[i]+k3y, z[i]+k3z)

		y[i+1] = y[i] + (k1y+2*k2y+2*k3y+k4y) / 6
		z[i+1] = z[i] + (k1z+2*k2z+2*k3z+k4z) / 6
		x[i+1] = x[i] + H
    }

    return x, y, z
}