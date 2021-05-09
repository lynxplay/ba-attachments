package latex

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type AverageDataPoints = map[string]map[string]float64

const LatexTemplate = `\begin{table}[h]
    \centering
    \begin{tabular}{l|ccc}
        \hline
        & HotSpot & OpenJ9 & Native \\
        \hline
        Quarkus & <<.Values.quarkus.hotspot>>[<<.Unit>>] & <<.Values.quarkus.openj9>>[<<.Unit>>] & <<.Values.quarkus.native>>[<<.Unit>>] \\
        Micronaut & <<.Values.micronaut.hotspot>>[<<.Unit>>] & <<.Values.micronaut.openj9>>[<<.Unit>>] & <<.Values.micronaut.native>>[<<.Unit>>] \\
        Helidon & <<.Values.helidon.hotspot>>[<<.Unit>>] & <<.Values.helidon.openj9>>[<<.Unit>>] & <<.Values.helidon.native>>[<<.Unit>>] \\
        SpringBoot & <<.Values.springboot.hotspot>>[<<.Unit>>] & <<.Values.springboot.openj9>>[<<.Unit>>] & <<.Values.springboot.native>>[<<.Unit>>] \\
        \hline
    \end{tabular}
    \caption{<<.Title>>}
\end{table}
`

func WriteLatexAverages(input AverageDataPoints, filePath string, title string, id string, unit string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = e
		}
	}()

	if tmpl, err := template.New("latex-table").Delims("<<", ">>").Parse(LatexTemplate); err != nil {
		return err
	} else {
		if err := tmpl.Execute(file, struct {
			Values AverageDataPoints
			Title  string
			Id     string
			Unit   string
		}{input, fmt.Sprintf("Rohdaten zum Diagramm: %s", title), strings.ReplaceAll(id, "_", "-"), unit}); err != nil {
			return err
		}
	}
	return err
}
