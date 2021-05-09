package data_plotter

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"visualizer/util"
)

const BytesInMegabyte float64 = 1024

func PlotMemoryData(data MemoryDataPoints, name string) (*plot.Plot, error) {
	p := util.NewDefPlot(name)
	p.X.Tick.Marker = util.NewSteppedPlotTicker(5)
	p.Y.Tick.Marker = util.NewSteppedPlotTicker(100)

	zeroTime := data[0].Timestamp

	cpuTime := make(plotter.XYs, len(data))
	cpuUsage := make(plotter.XYs, len(data))
	rss := make(plotter.XYs, len(data))

	for i, point := range data {
		var time = (point.Timestamp - zeroTime) / 1000

		cpuTime[i].X = time
		cpuTime[i].Y = point.CPUTime

		cpuUsage[i].X = time
		cpuUsage[i].Y = point.CPUUsage

		rss[i].X = time
		rss[i].Y = point.RSS / BytesInMegabyte
	}

	if err := plotutil.AddLinePoints(p,
		"cpu_time[sec]", cpuTime,
		"cpu_usage[%]", cpuUsage,
		"rss[mb]", rss,
	); err != nil {
		return nil, fmt.Errorf("failed to add lines to plot: %s", err)
	}
	return p, nil
}
