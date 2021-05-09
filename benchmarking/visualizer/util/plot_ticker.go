package util

import (
	"fmt"
	"gonum.org/v1/plot"
	"math"
)

type SteppedPlotTicker struct {
	StepSize float64
}

func NewSteppedPlotTicker(stepSize float64) SteppedPlotTicker {
	return SteppedPlotTicker{StepSize: stepSize}
}

func (m SteppedPlotTicker) Ticks(min, max float64) []plot.Tick {
	minEven := math.Floor(min/m.StepSize) * m.StepSize
	result := make([]plot.Tick, 0)
	if minEven != min {
		result = append(result, plot.Tick{
			Value: min,
			Label: fmt.Sprintf("%2.f", min),
		})
	}

	i := minEven
	for ; i < max-m.StepSize*0.5; i += m.StepSize {
		result = append(result, plot.Tick{
			Value: i,
			Label: fmt.Sprintf("%2.f", i),
		})
	}

	result = append(result, plot.Tick{
		Value: max,
		Label: fmt.Sprintf("%2.f", max),
	})

	return result
}
