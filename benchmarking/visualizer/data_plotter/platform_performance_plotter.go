package data_plotter

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"math"
	"visualizer/util"
)

func PlotPlatformPerformance(data ResultDataPointMap, platform string) (*plot.Plot, error) {
	p := util.NewDefPlot(fmt.Sprintf("%s/response", platform))
	p.X.Tick.Marker = util.NewSteppedPlotTicker(5)
	p.Y.Tick.Marker = util.NewSteppedPlotTicker(10)
	p.Y.Label.Text = "Response Time(ms)"
	p.Y.Min = 0

	quarkus := make(plotter.XYs, len(data[FrameworkQuarkus][platform].ClientDataPoints))
	micronaut := make(plotter.XYs, len(data[FrameworkMicronaut][platform].ClientDataPoints))
	helidon := make(plotter.XYs, len(data[FrameworkHelidon][platform].ClientDataPoints))
	springboot := make(plotter.XYs, len(data[FrameworkSpringBoot][platform].ClientDataPoints))

	appendFramework := func(ys plotter.XYs, data ClientDataPoints) {
		zeroTime := data[0].Timestamp
		for i, point := range data {
			var time = math.Max(0, (point.Timestamp-zeroTime)/1000)

			ys[i].X = time
			ys[i].Y = point.Latency
		}
	}

	appendFramework(quarkus, data[FrameworkQuarkus][platform].ClientDataPoints)
	appendFramework(micronaut, data[FrameworkMicronaut][platform].ClientDataPoints)
	appendFramework(helidon, data[FrameworkHelidon][platform].ClientDataPoints)
	appendFramework(springboot, data[FrameworkSpringBoot][platform].ClientDataPoints)

	if err := plotutil.AddLinePoints(p,
		"quarkus[ms]", quarkus,
		"micronaut[ms]", micronaut,
		"helidon[ms]", helidon,
		"springboot[ms]", springboot,
	); err != nil {
		return nil, fmt.Errorf("failed to add lines to plot: %s", err)
	}
	return p, nil
}
