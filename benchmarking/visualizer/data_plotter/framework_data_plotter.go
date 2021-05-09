package data_plotter

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"math"
	"visualizer/util"
)

func PlotFrameworkDataClientResponseTime(data FrameworkResultDataPoints, name string) (*plot.Plot, error) {
	p := util.NewDefPlot(name)
	p.Y.Label.Text = "Response Time[ms]"
	p.X.Tick.Marker = util.NewSteppedPlotTicker(5)
	p.Y.Tick.Marker = util.NewSteppedPlotTicker(10)
	p.Y.Min = 0

	if err := plotutil.AddLinePoints(p,
		PlatformHotSpot, createLinePoints(data[PlatformHotSpot].ClientDataPoints),
		PlatformOpenJ9, createLinePoints(data[PlatformOpenJ9].ClientDataPoints),
		PlatformNative, createLinePoints(data[PlatformNative].ClientDataPoints),
	); err != nil {
		return nil, fmt.Errorf("failed to add plot lines %s", err)
	}
	return p, nil
}

func createLinePoints(data ClientDataPoints) plotter.XYs {
	zeroTime := data[0].Timestamp
	latency := make(plotter.XYs, len(data))
	for i, point := range data {
		var time = math.Max(0, (point.Timestamp-zeroTime)/1000)

		latency[i].X = time
		latency[i].Y = math.Min(Max, point.Latency)
	}
	return latency
}

type MemoryPointValueFetcher = func(point MemoryDataPoint) float64

func PlotFrameworkDataResource(data FrameworkResultDataPoints, name string, yValueName string, fetcher MemoryPointValueFetcher) (*plot.Plot, error) {
	p := util.NewDefPlot(name)
	p.Y.Label.Text = yValueName

	if err := plotutil.AddLinePoints(p,
		PlatformHotSpot, createResourceLinePoints(data[PlatformHotSpot].MemoryDataPoints, fetcher),
		PlatformOpenJ9, createResourceLinePoints(data[PlatformOpenJ9].MemoryDataPoints, fetcher),
		PlatformNative, createResourceLinePoints(data[PlatformNative].MemoryDataPoints, fetcher),
	); err != nil {
		return nil, fmt.Errorf("failed to add plot lines %s", err)
	}
	return p, nil
}

func createResourceLinePoints(data MemoryDataPoints, fetcher MemoryPointValueFetcher) plotter.XYs {
	zeroTime := data[0].Timestamp
	resource := make(plotter.XYs, len(data))
	for i, point := range data {
		var time = math.Max(0, (point.Timestamp-zeroTime)/1000)

		resource[i].X = time
		resource[i].Y = fetcher(point)
	}
	return resource
}
