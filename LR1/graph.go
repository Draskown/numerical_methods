package lr1

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Defines a function to be interpolated
func f(x float64) float64 {
	return math.Pow(math.Log(x), 13/4)
}

// Define global values to hold input data
var x = []float64{2, 3, 4}
var y = []float64{f(x[0]), f(x[1]), f(x[2])}
var a = 2.5


// Calculates the polynomial at the given point
func calculatePolynomial(t float64, x []float64, y []float64) float64 {
    n := len(x)
    
    var s float64 = 0
    var p float64

    for i := 0; i < n; i++ {
        p = 1

        for j := 0; j < n; j++ {
            if i != j {
                p *= (t - x[j]) / (x[i] - x[j])
            }
        }

        s += y[i]*p
    }

    return s
}


// Creates a plot of the interpolated func
func CreatePolynomialGraph() (float64, float64){
    // Define the dataPol points for the interpolated function
    dataFunc := make(plotter.XYs, 0)    
    dataPol := make(plotter.XYs, 0)

    for i := 10; i <= 40; i++ {
        dX := float64(i) / 10
        dY := calculatePolynomial(dX, x, y)
        dataPol = append(dataPol, plotter.XY{X: dX, Y: dY})
        dataFunc = append(dataFunc, plotter.XY{X: dX, Y: f(dX)})
    }

    // Create the plot and set the title
    p := plot.New()
	
    p.Title.Text = "Interpolated Function Graph for ln(x)^13/4"

    // Set the X and Y axis labels
    p.X.Label.Text = "x"
    p.Y.Label.Text = "y"

    // Add the data points to the plot
    linePol, err := plotter.NewLine(dataPol)
    if err != nil {
        panic(err)
    }
    linePol.LineStyle.Color = plotutil.SoftColors[0]
    p.Add(linePol)
    
    funcLine, err := plotter.NewLine(dataFunc)
    if err != nil {
        panic(err)
    }
    funcLine.Color = plotutil.SoftColors[1]
    p.Add(funcLine)
    
    // Add a legend to the plot
    p.Legend.Add("Function", funcLine)
    p.Legend.Add("Polynomial", linePol)
    
    // Save the plot to a PNG file
    if err := p.Save(4*vg.Inch, 4*vg.Inch, "LR1/polynomial.png"); err != nil {
        panic(err)
    }

    polynomialA := calculatePolynomial(a, x, y)
    error := math.Abs(f(a) - polynomialA)

    return polynomialA, error
}


// Calculates the approximation of the function
func calculateApproximation(t float64, x []float64, y []float64) (float64, float64) {
    sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
    for i := 0; i < len(x); i++ {
        sumX += x[i]
        sumY += y[i]
        sumXY += x[i] * y[i]
        sumX2 += x[i] * x[i]
    }

    m := (float64(len(x))*sumXY - sumX*sumY) / (float64(len(x))*sumX2 - sumX*sumX)
    b := (sumY - m*sumX) / float64(len(x))
    
    return m, b
}


// Creates a plot of the approximated func
func CreateApproximationGraph() float64 {
    // Linear least-squares approximation
    m, b := calculateApproximation(a, x, y)

    res := math.Pow(math.Log(a), 13.0/14.0)*m + b

    // Create data points for plotting the function and the approximation
    nPoints := 50
    minX, maxX := 1.5, 4.5
    stepX := (maxX - minX) / float64(nPoints-1)
    ptsFunc := make(plotter.XYs, nPoints)
    ptsApprox := make(plotter.XYs, nPoints)
    for i := 0; i < nPoints; i++ {
        dX := minX + float64(i)*stepX
        ptsFunc[i].X = dX
        ptsFunc[i].Y = f(dX)
        ptsApprox[i].X = dX
        ptsApprox[i].Y = m*math.Log(dX) + b
    }

    // Create a new plot
    p := plot.New()
    
    p.Title.Text = "Linear least-squares approximation for ln(x)^13/4"
    p.X.Label.Text = "x"
    p.Y.Label.Text = "y"

    // Add data points to the plot
    funcLine, err := plotter.NewLine(ptsFunc)
    if err != nil {
        panic(err)
    }
    funcLine.LineStyle.Width = vg.Points(1)
    funcLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
    funcLine.Color = plotutil.SoftColors[0]
    p.Add(funcLine)

    approxLine, err := plotter.NewLine(ptsApprox)
    if err != nil {
        panic(err)
    }
    approxLine.LineStyle.Width = vg.Points(1)
    approxLine.Color = plotutil.SoftColors[1]
    p.Add(approxLine)

    // Add a legend to the plot
    p.Legend.Add("Function", funcLine)
    p.Legend.Add("Approximation", approxLine)
    
    if err := p.Save(4*vg.Inch, 4*vg.Inch, "LR1/approximation.png"); err != nil {
        panic(err)
    }

    return res
}