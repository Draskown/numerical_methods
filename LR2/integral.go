package lr2

// Calculates integrals for the input function
func CalculateIntegrals() (float64, float64, float64, float64) {
	// Length of the array
	n := len(X)

	// Integration step
	h := (X[n-1] - X[0]) / float64(n-1)

	// Init the needed values for formulas
	s := 0.0
	sum := 0.0
	sumEven := 0.0
	sumOdd := 0.0

	for i := 0; i < n-1; i++ {
		// Sum the squares of the intervals
		s += Y[i] * h

		// Calculating the sum of the Ys
		sum += Y[i]

		// Count the even elements
		if i%2 == 0 && i != 0 {
			sumEven += Y[i]
		} else {
			// Count the odd elements
			sumOdd += Y[i]
		}
	}

	// Apply Newton-Kotes formula
	nk := h / 2 * (Y[0] + 2*(Y[1]+Y[2]+Y[3]) + Y[n-1])

	// Apply trapeze formula
	t := h * (0.5*Y[0] + sum + 0.5*Y[n-1])

	// Apply the Simons formula
	sim := (h / float64(n-2)) * (Y[0] + 4*sumEven + 2*sumOdd + Y[n-1])

	return nk, s, t, sim
}