package plot_generators

import (
	"fmt"
	path2 "path"
	"visualizer/data_plotter"
)

func WriteFrameworkComparePlots(path string, data data_plotter.FrameworkResultDataPoints) error {
	frameworkName := path2.Base(path)

	if p, err := data_plotter.PlotFrameworkDataClientResponseTime(data, fmt.Sprintf("%s_response_times", frameworkName)); err != nil {
		return fmt.Errorf("failed to plot framework response time: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/response_time.png", path)); err != nil {
			return fmt.Errorf("failed to save %s/response_time: %s", path, err)
		}
		fmt.Printf("Generated %s response_time comparison!\n", frameworkName)
	}

	if p, err := data_plotter.PlotFrameworkDataResource(data, fmt.Sprintf("%s_rss_usage", frameworkName), "rss[mb]", func(point data_plotter.MemoryDataPoint) float64 {
		return point.RSS / data_plotter.BytesInMegabyte
	}); err != nil {
		return fmt.Errorf("failed to plot framework rss usage: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/rss_usage.png", path)); err != nil {
			return fmt.Errorf("failed to save %s/rss_usage: %s", path, err)
		}
		fmt.Printf("Generated %s rss_usage comparison!\n", frameworkName)
	}

	if p, err := data_plotter.PlotFrameworkDataResource(data, fmt.Sprintf("%s_cpu_time", frameworkName), "time[s]", func(point data_plotter.MemoryDataPoint) float64 {
		return point.CPUTime
	}); err != nil {
		return fmt.Errorf("failed to plot framework cpu_time: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/cpu_time.png", path)); err != nil {
			return fmt.Errorf("failed to save %s/cpu_time: %s", path, err)
		}
		fmt.Printf("Generated %s cpu_time comparison!\n", frameworkName)
	}

	if p, err := data_plotter.PlotFrameworkDataResource(data, fmt.Sprintf("%s_cpu_usage", frameworkName), "percentage", func(point data_plotter.MemoryDataPoint) float64 {
		return point.CPUUsage
	}); err != nil {
		return fmt.Errorf("failed to plot framework cpu_usage: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/cpu_usage.png", path)); err != nil {
			return fmt.Errorf("failed to save %s/cpu_usage: %s", path, err)
		}
		fmt.Printf("Generated %s cpu_usage comparison!\n", frameworkName)
	}
	return nil
}
