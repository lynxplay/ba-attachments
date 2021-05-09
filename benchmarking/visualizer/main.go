package main

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"visualizer/data_mutator"
	"visualizer/data_plotter"
	"visualizer/plot_generators"
	"visualizer/util"
)

/// main may easily be called using `find ./../results -name "hotspot" -or -name "openj9" -or -name "native" | xargs go run *.go`
/// to create all images
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Could not plot graphs for %d directories. (expected 1)\n", len(args))
		return
	}

	InitPlotter()

	root := args[0]
	folders, err := parseResultFolders(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	mode := os.Getenv("BA_MODE")
	if mode == "" {
		mode = "GEN"
	}
	fmt.Printf("Running in %s mode.\n", mode)

	// Mutate data
	if mode == "OPT" {
		MutateData(folders, root)
		return
	}

	visualFolder := path.Join(root, ".visuals")
	if _, err := os.Stat(visualFolder); os.IsNotExist(err) {
		if err := os.Mkdir(visualFolder, 0777); err != nil {
			fmt.Printf("Failed to generate visuals folder: %s\n", err)
			return
		}
	}

	memoryPoints := make(data_plotter.ResultDataPointMap)
	for framework, jvms := range folders {
		memoryPoints[framework] = make(map[string]data_plotter.ResultDataPoints) // init map

		for _, jvm := range jvms {
			finalJvm := jvm
			finalFramework := framework
			sourcePath := path.Join(root, finalFramework, finalJvm)
			outputPath := path.Join(visualFolder, framework, jvm)
			err = os.MkdirAll(outputPath, 0777) // make output dir
			if err != nil {
				fmt.Printf("failed to create output directory for %s/%s: %s\n", framework, jvm, err)
				return
			}

			memoryDataPoints, err := plot_generators.ReadMemoryPoints(sourcePath) // read in memory points
			if err != nil {
				fmt.Printf("Failed to generate memory points %s/%s: %s\n", finalFramework, finalJvm, err)
				return
			}

			if p, err := data_plotter.PlotMemoryData(memoryDataPoints, util.PathToName(sourcePath, 2)); err != nil { // Plot and write memory points
				fmt.Printf("failed to plot memory: %s\n", err)
				return
			} else {
				if err := p.Save(plot_generators.PlotWidth, plot_generators.PlotHeight, fmt.Sprintf("%s/memory_plot.png", outputPath)); err != nil {
					fmt.Printf("failed to save plot: %s", err)
				}
			}
			fmt.Printf("Generated memory plot for %s/%s!\n", finalFramework, finalJvm)

			clientDataPoints, err := plot_generators.ReadClientPoints(path.Join(root, finalFramework, finalJvm)) // Read in client points
			if err != nil {
				fmt.Printf("Failed to generate client points %s/%s: %s\n", finalFramework, finalJvm, err)
				return
			}
			if p, err := data_plotter.PlotClientData(clientDataPoints, util.PathToName(sourcePath, 2)); err != nil { // Plot and write client points
				fmt.Printf("Failed to plot client points: %s\n", err)
				return
			} else {
				if err := p.Save(plot_generators.PlotWidth, plot_generators.PlotHeight, fmt.Sprintf("%s/client_plot.png", outputPath)); err != nil {
					fmt.Printf("Failed to save client plot: %s", err)
					return
				}
			}
			fmt.Printf("Generated client plot for %s/%s!\n", finalFramework, finalJvm)

			memoryPoints[finalFramework][finalJvm] = data_plotter.ResultDataPoints{
				MemoryDataPoints: memoryDataPoints,
				ClientDataPoints: clientDataPoints,
			}
		}

		if err := plot_generators.WriteFrameworkComparePlots(fmt.Sprintf("%s/%s", visualFolder, framework), memoryPoints[framework]); err != nil {
			fmt.Printf("failed to generate framework comaparison: %s\n", err)
			return
		}
	}

	if err := plot_generators.WritePlatformPerformancePoint(visualFolder, memoryPoints); err != nil {
		fmt.Printf("Failed to generate platform performance: %s\n", err)
		return
	}

	// Averages
	if err := plot_generators.WriteAverageData(visualFolder, memoryPoints); err != nil {
		fmt.Printf("Failed to generate average plots: %s\n", err)
		return
	}

	fmt.Printf("Done!\n")
}

func MutateData(folders map[string][]string, root string) {
	for framework, jvms := range folders {
		for _, jvm := range jvms {
			data_mutator.AverageRaw(path.Join(root, framework, jvm))
			fmt.Printf("Optimized %s/%s/%s\n", root, framework, jvm)
		}
	}
}

func parseResultFolders(root string) (map[string][]string, error) {
	folders := make(map[string][]string, 0)
	rootDir, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("failed to parse root directory: %s", err)
	}
	for _, frameworkDir := range rootDir {
		if strings.HasPrefix(frameworkDir.Name(), ".") {
			continue // Ignore hidden folders (such as .visuals)
		}
		jvmFiles, err := ioutil.ReadDir(fmt.Sprintf("%s%c%s", root, os.PathSeparator, frameworkDir.Name()))
		if err != nil {
			continue // Ignore files
		}
		for _, jvmFolder := range jvmFiles {
			if !jvmFolder.IsDir() {
				break // Ignore files
			}
			folders[frameworkDir.Name()] = append(folders[frameworkDir.Name()], jvmFolder.Name())
		}
	}
	return folders, nil
}

func InitPlotter() {
	plotter.DefaultLineStyle.Width = 2 * vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = 0
	plotutil.DefaultGlyphShapes = plotutil.DefaultGlyphShapes[0:1]
	plotutil.DefaultDashes = plotutil.DefaultDashes[0:1]
	plotutil.DefaultColors = append(plotutil.DefaultColors, plotutil.DarkColors...)
}
