package data_plotter

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"visualizer/latex"
	"visualizer/util"
)

type PointAverageFetcher = func(r ResultDataPoints) float64

func PlotAverageData(data ResultDataPointMap, name string, yName string, fetcher PointAverageFetcher) (*plot.Plot, latex.AverageDataPoints, error) {
	p := util.NewDefPlot(name)
	p.Title.Text = name
	p.Y.Label.Text = yName
	p.X.Label.Text = "Platform"

	barWidth := vg.Millimeter * 10
	halfWidth := barWidth / 2

	quarkusDataPoints := getAverageValuesFor(data[FrameworkQuarkus], fetcher)
	quarkusBars, err := plotter.NewBarChart(toValues(quarkusDataPoints), barWidth)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create bars for quarkus: %s", err)
	}

	micronautDataPoints := getAverageValuesFor(data[FrameworkMicronaut], fetcher)
	micronautBars, err := plotter.NewBarChart(toValues(micronautDataPoints), barWidth)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create bars for micronaut: %s", err)
	}

	helidonDataPoints := getAverageValuesFor(data[FrameworkHelidon], fetcher)
	helidonBars, err := plotter.NewBarChart(toValues(helidonDataPoints), barWidth)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create bars for helidon: %s", err)
	}

	springbootDataPoints := getAverageValuesFor(data[FrameworkSpringBoot], fetcher)
	springbootBars, err := plotter.NewBarChart(toValues(springbootDataPoints), barWidth)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create bars for spingboot: %s", err)
	}

	quarkusBars.Color = plotutil.Color(0)
	quarkusBars.Offset = -barWidth - halfWidth
	micronautBars.Color = plotutil.Color(1)
	micronautBars.Offset = -halfWidth
	helidonBars.Color = plotutil.Color(2)
	helidonBars.Offset = halfWidth
	springbootBars.Color = plotutil.Color(3)
	springbootBars.Offset = barWidth + halfWidth

	p.Add(quarkusBars, micronautBars, helidonBars, springbootBars)
	p.Legend.Add(FrameworkQuarkus, quarkusBars)
	p.Legend.Add(FrameworkMicronaut, micronautBars)
	p.Legend.Add(FrameworkHelidon, helidonBars)
	p.Legend.Add(FrameworkSpringBoot, springbootBars)
	p.Legend.Left = true
	p.NominalX(PlatformHotSpot, PlatformOpenJ9, PlatformNative)

	fmt.Printf("Generated averaged plot for %s\n", name)

	points := make(latex.AverageDataPoints)
	points[FrameworkQuarkus] = quarkusDataPoints
	points[FrameworkMicronaut] = micronautDataPoints
	points[FrameworkHelidon] = helidonDataPoints
	points[FrameworkSpringBoot] = springbootDataPoints
	return p, points, nil
}

func getAverageValuesFor(frameworkData map[Platform]ResultDataPoints, fetcher PointAverageFetcher) map[string]float64 {
	return map[string]float64{
		PlatformHotSpot: fetcher(frameworkData[PlatformHotSpot]),
		PlatformOpenJ9:  fetcher(frameworkData[PlatformOpenJ9]),
		PlatformNative:  fetcher(frameworkData[PlatformNative]),
	}
}

func toValues(input map[string]float64) plotter.Values {
	return plotter.Values{
		input[PlatformHotSpot],
		input[PlatformOpenJ9],
		input[PlatformNative],
	}
}
