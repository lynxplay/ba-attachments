package data_mutator

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"strings"
	"visualizer/data_plotter"
	"visualizer/util"
)

var Divider = 2000

func AverageRaw(folder string) {
	source, err := os.Open(path.Join(folder, "raw_last_result_aggregate.csv"))
	if err != nil {
		panic(err)
	}
	defer source.Close()

	target, err := os.Create(path.Join(folder, "last_result_aggregate.csv"))
	if err != nil {
		panic(err)
	}
	defer target.Close()

	reader := bufio.NewReader(source)
	table, err := util.ParseCSVIndexTable(reader)
	if err != nil {
		panic(err)
	}

	target.Write([]byte(fmt.Sprintf(
		"%s,%s,%s,%s\n",
		"timeStamp",
		"elapsed",
		"Latency",
		"Connect",
	)))
	buffer := make([]data_plotter.ClientDataPoint, Divider)
	idx := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			writeAverage(target, buffer, idx)
			return
		}
		if len(strings.TrimSpace(string(line))) == 0 {
			writeAverage(target, buffer, idx)
			return
		}
		csvValues := strings.Split(string(line), ",")
		if dataPoint, err := data_plotter.ParseClientDataPoint(
			csvValues[table["timeStamp"]],
			csvValues[table["elapsed"]],
			csvValues[table["Latency"]],
			csvValues[table["Connect"]],
		); err != nil {
			panic(err)
		} else {
			buffer[idx] = dataPoint
			idx++
		}

		if idx >= Divider {
			writeAverage(target, buffer, idx)
			buffer = make([]data_plotter.ClientDataPoint, Divider)
			idx = 0
		}
	}
}

func writeAverage(writer io.Writer, points []data_plotter.ClientDataPoint, idx int) {
	if idx < 1 {
		return
	}
	avrg := data_plotter.ClientDataPoint{}
	div := float64(idx)
	for _, point := range points {
		avrg.Timestamp += point.Timestamp / div
		avrg.Elapsed += point.Elapsed / div
		avrg.Latency += point.Latency / div
		avrg.ConnectDuration += point.ConnectDuration / div
	}

	result := fmt.Sprintf(
		"%s,%d,%d,%d\n",
		fmt.Sprintf("%0.f", avrg.Timestamp),
		int(math.Min(data_plotter.Max, avrg.Elapsed)),
		int(math.Min(data_plotter.Max, avrg.Latency)),
		int(math.Min(data_plotter.Max, avrg.ConnectDuration)),
	)

	_, err := writer.Write([]byte(result))
	if err != nil {
		panic(err)
	}
}
