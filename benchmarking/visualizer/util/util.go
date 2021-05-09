package util

import (
	"bufio"
	"bytes"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"io"
	"log"
	"os"
	"strings"
)

func CloseVerbose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatalf("Failed to close resource: %s", err)
	}
}

func PathToName(path string, depth int) string {
	source := strings.Split(path, string(os.PathSeparator))

	var skip = len(source) - depth
	if skip < 0 {
		skip = 0
	}
	return strings.Join(source[skip:], "/")
}

func NewDefPlot(name string) *plot.Plot {
	p := plot.New()
	p.Title.Text = name
	p.Add(plotter.NewGrid())

	p.X.Label.Text = "Time since begin of experiment(s)"
	p.X.Tick.Label.Font.Variant = "Mono"
	p.Y.Tick.Label.Font.Variant = "Mono"
	p.Legend.Top = true
	return p
}

func ParseCSVFile(reader *bufio.Reader, length uint64, consumer func(line []string, idx uint64) error) error {
	var idx uint64 = 0
	var lastLine []string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
		}
		if idx == length {
			return nil
		}
		lastLine = strings.Split(string(line), ",")
		if err := consumer(lastLine, idx); err != nil {
			return fmt.Errorf("parsing failed due to line consumer: %s", err)
		}
		idx++
	}
}

func ParseCSVIndexTableAndLength(tableReader *bufio.Reader, source []byte) (map[string]int, uint64, error) {
	length, err := ParseReaderLineLength(bytes.NewReader(source))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to parse line length: %s", err)
	}

	csvIndexTable, err2 := ParseCSVIndexTable(tableReader)
	if err2 != nil {
		return nil, 0, err2
	}
	return csvIndexTable, length - 1, nil
}

func ParseCSVIndexTable(tableReader *bufio.Reader) (map[string]int, error) {
	csvIndexTable := make(map[string]int)
	if firstLine, _, err := tableReader.ReadLine(); err != nil {
		return nil, fmt.Errorf("failed to parse csv header: %s", err)
	} else {
		for i, collum := range strings.Split(string(firstLine), ",") {
			csvIndexTable[collum] = i
		}
	}
	return csvIndexTable, nil
}

func ParseReaderLineLength(reader io.Reader) (uint64, error) { // see https://stackoverflow.com/a/52153000
	var count uint64
	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
