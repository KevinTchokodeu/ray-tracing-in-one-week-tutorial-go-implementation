package main

// Importing packages
import (
	"os"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/trace"
)

// Main function
func main() {
	blue := trace.NewLambert(trace.NewColor(0.1, 0.2, 0.5))
	yellow := trace.NewLambert(trace.NewColor(0.8, 0.8, 0.0))
	bronze := trace.NewMetal(trace.NewColor(0.8, 0.6, 0.2), 0)
	glass := trace.NewDielectric(1.5)

	l := trace.NewList(
		trace.NewSphere(geom.NewVec(0, 0, -1), 0.5, blue),
		trace.NewSphere(geom.NewVec(0, -100.5, -1), 100, yellow),
		trace.NewSphere(geom.NewVec(1, 0, -1), 0.5, bronze),
		trace.NewSphere(geom.NewVec(-1, 0, -1), 0.5, glass),
		trace.NewSphere(geom.NewVec(-1, 0, -1), -0.45, glass),
	)
	f := trace.NewWindow(200, 100)
	if err := f.WritePPM(os.Stdout, l, 100); err != nil {
		panic(err)
	}
}
