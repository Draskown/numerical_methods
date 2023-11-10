package lr3

func F(x, y float64) float64 {
	return x*x + y*y
}

var Y0 float64 = 0.4
var X0min, X0max float64 = 0.0, 1.0

func FSystemY(x, y, z float64) float64 {
	return x*y + z
}

func FSystemZ(y, z float64) float64 {
	return y - z
}

var X0SystemMin, X0SystemMax float64 = 0, 1
var Y0System float64 = 1
var Z0System float64 = 0

var H float64 = 0.1