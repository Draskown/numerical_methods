package lr1

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var A = "Fuck"

func CreateGraph() {
	// Create a new plot with a title and axis labels
	p := plot.New()

	p.Title.Text = "ln(x)^13/4"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	// Define the function to be plotted
	f := func(x float64) float64 {
		return math.Pow(math.Log(x), 13.0/4.0)
	}

	// Create a set of data points for the function over the interval [1, 10]
	n := 100
	data := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		x := 1.0 + float64(i)*(10.0-1.0)/float64(n-1)
		data[i].X = x
		data[i].Y = f(x)
	}

	// Create a line plot for the data points
	line, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}

	// Add the line plot to the plot
	p.Add(line)

	// Set the range of the x and y axis
	p.X.Min = 1.0
	p.X.Max = 10.0
	p.Y.Min = 0.0

	// Save the plot to a PNG file
	err = p.Save(10*vg.Centimeter, 10*vg.Centimeter,"./LR1/polynom.png")
	if err != nil {
		panic(err)
	}
}
