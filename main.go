package main

import (
	"bytes"
	"github.com/wcharczuk/go-chart/v2"
	"math"
	"os"
)

// Some data with jumpy values
var roughValues = []float64{0, 1, 0, 1, 10, 0, 1, 0, 8, 1, 3, 1, 1, 3, 9, 0, 1, 0, 1, 10, 0, 1, 0, 8, 1, 3, 1, 1, 3, 9, 0, 1, 0, 1, 10, 0, 1, 0, 8, 1, 3, 1, 1, 3, 9, 0, 1, 0, 1, 10, 0, 1, 0, 8, 1, 3, 1, 1, 3, 9, 2, 1, 2, 3}

func main() {
	// Copy values of `roughValues` into `smoothValues` to see a before-after view.
	smoothValues := make([]float64, len(roughValues))
	xValues := make([]float64, len(roughValues)) // This will be plotted on the X axis

	for i := range roughValues {
		smoothValues[i] = roughValues[i]
		xValues[i] = float64(i)
	}

	windowSize := 4 // Try changing the window size to see how the curve changes!
	smoothData(smoothValues, windowSize)

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: roughValues,
			},
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: smoothValues,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output.png", buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func smoothData(data []float64, windowSize int) {
	for i := 0; i < len(data); i += windowSize {
		_windowSize := windowSize
		if i+_windowSize > len(data) {
			// Clamp window size to prevent array out of bounds access
			_windowSize = len(data)-i
		}

		smooth(data[i:_windowSize+i])
	}
}

func smooth(window []float64) {
	windowSize := len(window)
	trueWindowSize := windowSize - 1

	maxima := -99999.0
	minima := 99999.0
	argMax := -1

	for i, val := range window {
		if val > maxima {
			maxima = val
			argMax = i
		}
		if val < minima {
			minima = val
		}
	}

	// This allows us to handle negative peaks (ie. troughs)
	if math.Abs(maxima) < math.Abs(minima) {
		_maxima := maxima
		maxima = minima
		minima = _maxima
	}

	// If you have no idea what's happening here. Read the readme!
	poi := -1.0

	if trueWindowSize-argMax > argMax-0 {
		poi = float64(trueWindowSize)
	} else if trueWindowSize-argMax < argMax-0 {
		poi = 0.0
	} else {
		poi = 0.0
	}

	// You could multiply twice to square, but math.Pow() is clearer

	slope := (minima - maxima) / math.Pow(poi-float64(argMax), 2)

	// Evaluate curve at points
	for i := range window {
		window[i] = slope*math.Pow(float64(i-argMax), 2) + maxima
	}
}
