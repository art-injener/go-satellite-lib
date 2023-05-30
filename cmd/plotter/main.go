package main

import (
	"fmt"
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/art-injener/go-satellite-lib/internal/dsp/generator"
)

func main() {
	frequency := 100.0 // Начальная частота сигнала
	amplitude := 0.5   // Начальная амплитуда сигнала
	signalBytes := generator.SimpleSignal(frequency, amplitude)

	// Преобразование сигнала в точки графика
	pts := make(plotter.XYs, len(signalBytes))
	for i, sample := range signalBytes {
		pts[i].X = float64(i) / float64(len(signalBytes))
		pts[i].Y = float64(sample) / 255
	}

	// Создание графика
	p := plot.New()

	p.Title.Text = "График радиосигнала"
	p.X.Label.Text = "Время (с)"
	p.Y.Label.Text = "Амплитуда"

	// Добавление линии на график
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	line.Color = color.RGBA{R: 255, A: 255}
	p.Add(line)

	// Сохранение графика в файл
	err = p.Save(10*vg.Inch, 4*vg.Inch, "graph.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("График успешно сохранен в файл graph.png")
}
