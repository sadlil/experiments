package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	graph := chart.Chart{
		Title: "test-chart",
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)

	fmt.Println(err)

	ioutil.WriteFile("a.png", buffer.Bytes(), os.ModePerm)
}
