package data_plotter

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"math"
	"visualizer/util"
)

const Max = 250

func PlotClientData(data ClientDataPoints, name string) (*plot.Plot, error) {
	p := util.NewDefPlot(name)
	p.X.Tick.Marker = util.NewSteppedPlotTicker(5)
	p.Y.Tick.Marker = util.NewSteppedPlotTicker(10)

	zeroTime := data[0].Timestamp

	elapsed := make(plotter.XYs, len(data))
	latency := make(plotter.XYs, len(data))
	connectDuration := make(plotter.XYs, len(data))

	for i, point := range data {
		var time = math.Max(0, (point.Timestamp-zeroTime)/1000)

		elapsed[i].X = time
		elapsed[i].Y = math.Min(Max, point.Elapsed)

		latency[i].X = time
		latency[i].Y = math.Min(Max, point.Latency)

		connectDuration[i].X = time
		connectDuration[i].Y = math.Min(Max, point.ConnectDuration)
	}

	if err := plotutil.AddLinePoints(p,
		"elapsed[ms]", elapsed,
		"latency[ms]", latency,
		"connectDuration[ms]", connectDuration,
	); err != nil {
		return nil, fmt.Errorf("failed to add lines to plot: %s", err)
	}
	return p, nil
}
