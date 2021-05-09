package plot_generators

import (
	"bufio"
	"bytes"
	"fmt"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"visualizer/data_plotter"
	"visualizer/util"
)


const (
	PlotHeight = 120 * vg.Millimeter
	PlotWidth  = 235 * vg.Millimeter
)

const (
	ProcessMemoryFileName  string = "memory.csv"
	ClientResponseFileName string = "last_result_aggregate.csv"
)

func ReadMemoryPoints(path string) ([]data_plotter.MemoryDataPoint, error) {
	memoryCSVData, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, ProcessMemoryFileName))
	if err != nil {
		return nil, err
	}

	memoryCSVReader := bufio.NewReader(bytes.NewReader(memoryCSVData))
	memoryCSVIndex, memoryCSVLength, err := util.ParseCSVIndexTableAndLength(memoryCSVReader, memoryCSVData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse memory csv: %s", err)
	}

	memoryData := make([]data_plotter.MemoryDataPoint, memoryCSVLength)
	if err := util.ParseCSVFile(memoryCSVReader, memoryCSVLength, func(line []string, idx uint64) error {
		if dataPoint, err := data_plotter.ParseMemoryDataPoint(
			line[memoryCSVIndex["TIMESTAMP"]],
			line[memoryCSVIndex["TIME"]],
			line[memoryCSVIndex["%CPU"]],
			line[memoryCSVIndex["RSS"]],
		); err != nil {
			return fmt.Errorf("failed to parse memory data point: %s", err)
		} else {
			memoryData[idx] = dataPoint
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to parse memory csv: %s", err)
	}
	return memoryData, nil
}

func ReadClientPoints(path string) ([]data_plotter.ClientDataPoint, error) {
	clientCSVData, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, ClientResponseFileName))
	if err != nil {
		return nil, err
	}

	clientCSVReader := bufio.NewReader(bytes.NewReader(clientCSVData))
	clientCSVIndex, clientCSVLength, err := util.ParseCSVIndexTableAndLength(clientCSVReader, clientCSVData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client csv: %s", err)
	}

	clientData := make([]data_plotter.ClientDataPoint, clientCSVLength)
	if err := util.ParseCSVFile(clientCSVReader, clientCSVLength, func(line []string, idx uint64) error {
		if dataPoint, err := data_plotter.ParseClientDataPoint(
			line[clientCSVIndex["timeStamp"]],
			line[clientCSVIndex["elapsed"]],
			line[clientCSVIndex["Latency"]],
			line[clientCSVIndex["Connect"]],
		); err != nil {
			return fmt.Errorf("failed to parse client data point: %s", err)
		} else {
			clientData[idx] = dataPoint
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to parse client csv: %s", err)
	}
	return clientData, nil
}

func WritePlatformPerformancePoint(path string, data data_plotter.ResultDataPointMap) error {
	if len(data) < 1 {
		return fmt.Errorf("cannot compare 0 platforms")
	}
	for framework := range data {
		for name := range data[framework] {
			p, err := data_plotter.PlotPlatformPerformance(data, name)
			if err != nil {
				return fmt.Errorf("failed to plot performance: %s", err)
			}
			if err := p.Save(PlotWidth, PlotHeight, fmt.Sprintf("%s/platform_performance_%s.png", path, name)); err != nil {
				return fmt.Errorf("failed to save plot: %s", err)
			}
		}
		break
	}
	return nil
}
