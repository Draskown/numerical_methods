package lr3

// Caclulates the values for the prediction at the next point y
func milne(x0, y0, h float64) float64{
	f := F(x0, y0)
	f1 := F(x0+h, y0+h*f)
	f2 := F(x0+2*h, y0+2*h*f1)
	y1 := y0 + h/3.0*(f+f2+4*f1)
	return y1
}

// Solves a Cauchy problem of the equation using Milne's formula
func Milne() ([]float64, []float64) {
	// Calculate the number of steps
    n := int((X0max - X0min) / H)

	// Initialize arrays to hold x and y values
	x := make([]float64, n+1)
	y := make([]float64, n+1)
 
	// Set initial values
	x[0] = X0min
	y[0] = Y0

	// Predict the values of Y
	for i := 0; i < n; i++ {
		y1 := milne(x[i]-H, y[i], H)
		y2 := milne(x[i]-2*H, y[i], 2*H)
		y3 := milne(x[i], y2, H)
		y[i+1] = y1 + (y1-y3)/15.0
		x[i+1] = x[i]+H
	}

	return x, y
}

// Caclulates the values for the prediction at the next points y and z
func milneSystem(x0, y0, z0, h float64) (float64, float64){
	f1 := FSystemY(x0, y0, z0)
	f2 := FSystemZ(y0, z0)
	f3 := FSystemY(x0+h, y0+h*f1, z0+h*f2)
	f4 := FSystemZ(y0+h*f1, z0+h*f2)
	y1 := y0 + h/3.0*(f1+f3+4*f2)
	z1 := z0 + h/3.0*(f2+f4+4*f3)
	return y1, z1
}

// Solves a Cauchy problem of the system of equations using Milne's formula
func MilneSystem() ([]float64, []float64, []float64) {
	// Calculate the number of steps
    n := int((X0SystemMax - X0SystemMin) / H)

	// Initialize arrays to hold x and y values
	x := make([]float64, n+1)
	y := make([]float64, n+1)
	z := make([]float64, n+1)
 
	// Set initial values
	x[0] = X0SystemMin
	y[0] = Y0System
	z[0] = Z0System

	// Predict the values of y and z
	for i := 0; i < n; i++ {
		y1, z1 := milneSystem(x[i]-H, y[i], z[i], H)
		y2, z2 := milneSystem(x[i]-2*H, y[i], z[i], 2*H)
		y3, z3 := milneSystem(x[i], y2, z2, H)
		x[i+1] = x[i]+H
		y[i+1] = y1 + (y1-y3)/15.0
		z[i+1] = z1 + (z1-z3)/15.0
	}

	return x, y, z
}