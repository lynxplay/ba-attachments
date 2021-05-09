package plot_generators

import (
	"fmt"
	"os"
	"sort"
	"visualizer/data_plotter"
	"visualizer/latex"
)

func WriteAverageData(path string, data data_plotter.ResultDataPointMap) error {
	if err := os.MkdirAll(fmt.Sprintf("%s/.tex", path), 0777); err != nil {
		return err
	}

	if p, data, err := data_plotter.PlotAverageData(data, "used_cpu_time_total", "time[s]", func(r data_plotter.ResultDataPoints) float64 {
		cpuTimes := make([]float64, len(r.MemoryDataPoints))
		for i, point := range r.MemoryDataPoints {
			cpuTimes[i] = point.CPUTime
		}
		sort.Float64s(cpuTimes)
		return cpuTimes[len(cpuTimes)-1]
	}); err != nil {
		return fmt.Errorf("failed to plot used_cpu_time_total: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/used_cpu_time_total.png", path)); err != nil {
			return fmt.Errorf("failed to save used_cpu_time_total: %s", err)
		}
		if err := latex.WriteLatexAverages(
			data,
			fmt.Sprintf("%s/.tex/used_cpu_time_total.tex", path),
			"Insgesamt benötigte CPU Zeit",
			"used_cpu_time_total",
			"s",
		); err != nil {
			return err
		}
	}

	if p, data, err := data_plotter.PlotAverageData(data, "used_cpu_time_at_5", "time[s]", func(r data_plotter.ResultDataPoints) float64 {
		timeAtZero := r.MemoryDataPoints[0].Timestamp
		for _, point := range r.MemoryDataPoints {
			if timeAtZero+float64(5000) <= point.Timestamp {
				return point.CPUTime
			}
		}
		panic("No data point at 5 seconds??!")
	}); err != nil {
		return fmt.Errorf("failed to plot used_cpu_time_at_5: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/used_cpu_time_at_5.png", path)); err != nil {
			return fmt.Errorf("failed to save used_cpu_time_at_5 plot: %s", err)
		}
		if err := latex.WriteLatexAverages(
			data,
			fmt.Sprintf("%s/.tex/used_cpu_time_at_5.tex", path),
			"Benötigte CPU Zeit nach 5 Sekunden des Experimentes",
			"used_cpu_time_at_5",
			"s",
		); err != nil {
			return err
		}
	}

	if p, data, err := data_plotter.PlotAverageData(data, "average_response_time", "time[ms](median)", func(r data_plotter.ResultDataPoints) float64 {
		resTime := make([]float64, len(r.ClientDataPoints))
		for i, point := range r.ClientDataPoints {
			resTime[i] = point.Latency
		}
		sort.Float64s(resTime)
		return resTime[(len(resTime)-1)/2]
	}); err != nil {
		return fmt.Errorf("failed to plot average response time: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/average_response_time.png", path)); err != nil {
			return fmt.Errorf("failed to save average_response_time plot: %s", err)
		}
		if err := latex.WriteLatexAverages(
			data,
			fmt.Sprintf("%s/.tex/average_response_time.tex", path),
			"Durchschnittliche Antwortgeschwindigkeit",
			"average_response_time",
			"ms",
		); err != nil {
			return err
		}
	}

	if p, data, err := data_plotter.PlotAverageData(data, "total_response_time", "time[s]", func(r data_plotter.ResultDataPoints) float64 {
		return (r.ClientDataPoints[len(r.ClientDataPoints)-1].Timestamp - r.ClientDataPoints[0].Timestamp) / 1000
	}); err != nil {
		return fmt.Errorf("failed to plot average response time: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/total_response_time.png", path)); err != nil {
			return fmt.Errorf("failed to save total_response_time plot: %s", err)
		}
		if err := latex.WriteLatexAverages(
			data,
			fmt.Sprintf("%s/.tex/total_response_time.tex", path),
			"Insgesamt benötigte Zeit des Belastungstests",
			"total_response_time",
			"s",
		); err != nil {
			return err
		}
	}

	if p, data, err := data_plotter.PlotAverageData(data, "start_rss", "rss[mb]", func(r data_plotter.ResultDataPoints) float64 {
		return r.MemoryDataPoints[0].RSS / data_plotter.BytesInMegabyte
	}); err != nil {
		return fmt.Errorf("failed to plot start_rss: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/start_rss.png", path)); err != nil {
			return fmt.Errorf("failed to save start_rss plot: %s", err)
		}
		if err := latex.WriteLatexAverages(
			data,
			fmt.Sprintf("%s/.tex/start_rss.tex", path),
			"Arbeitsspeicher zum Start des Experiments",
			"start_rss",
			"mb",
		); err != nil {
			return err
		}
	}

	if p, data, err := data_plotter.PlotAverageData(data, "start_rss_at_5", "rss[mb]", func(r data_plotter.ResultDataPoints) float64 {
		timeAtZero := r.MemoryDataPoints[0].Timestamp
		for _, point := range r.MemoryDataPoints {
			if timeAtZero+float64(5000) <= point.Timestamp {
				return point.RSS / data_plotter.BytesInMegabyte
			}
		}
		panic("No data point at 5 seconds??!")
	}); err != nil {
		return fmt.Errorf("failed to plot start_rss_at_5: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/start_rss_at_5.png", path)); err != nil {
			return fmt.Errorf("failed to save start_rss_at_5 plot: %s", err)
		}
		if err := latex.WriteLatexAverages(
			data,
			fmt.Sprintf("%s/.tex/start_rss_at_5.tex", path),
			"Arbeitsspeicher nach 5 Sekunden des Experiments",
			"start_rss_at_5",
			"mb",
		); err != nil {
			return err
		}
	}

	if p, data, err := data_plotter.PlotAverageData(data, "end_rss", "rss[mb]", func(r data_plotter.ResultDataPoints) float64 {
		return r.MemoryDataPoints[len(r.MemoryDataPoints)-1].RSS / data_plotter.BytesInMegabyte
	}); err != nil {
		return fmt.Errorf("failed to plot end_rss: %s", err)
	} else {
		if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/end_rss.png", path)); err != nil {
			return fmt.Errorf("failed to save end_rss plot: %s", err)
		}
		if err := latex.WriteLatexAverages(
			data,
			fmt.Sprintf("%s/.tex/end_rss.tex", path),
			"Arbeitsspeicher nach Ende des Experiments",
			"end_rss",
			"mb",
		); err != nil {
			return err
		}
	}

	return nil
}
