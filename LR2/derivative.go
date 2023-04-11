package lr2

// Calculates first and second derivatives of the inpus
func CalculateDerivatives() ([]float64, []float64) {

	var dx []float64
	var ddx []float64

	r := len(X)-2

	// First derivative
    for i := 0; i < 4; i++ {
        var d1 float64
        if i == 0 {
            d1 = (Y[1] - Y[0]) / (X[1] - X[0])
        } else if i == r {
            d1 = (Y[r] - Y[r-1]) / (X[r] - X[r-1])
        } else {
            d1 = (Y[i+1] - Y[i-1]) / (X[i+1] - X[i-1])
        }

		dx = append(dx, d1)
    }
    
    // Second derivative
    for i := 0; i < 4; i++ {
        var d2 float64
        if i == 0 {
            d2 = (Y[2] - 2*Y[1] + Y[0]) / ((X[1] - X[0])*(X[2] - X[0]))
        } else if i == r {
            d2 = (Y[r] - 2*Y[r-1] + Y[r-2]) / ((X[r] - X[r-1])*(X[r] - X[r-2]))
        } else {
            d2 = (Y[i+1] - 2*Y[i] + Y[i-1]) / ((X[i+1] - X[i])*(X[i] - X[i-1]))
        }

		ddx = append(ddx, d2)
    }

	return dx, ddx
}