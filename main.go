package main

// Importing packages
import (
	"os"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/trace"
)

// Main function
func main() {
	l := trace.NewList(
		trace.NewSphere(geom.NewVec(0, 0, -1), 0.5),
		trace.NewSphere(geom.NewVec(0, -100.5, -1), 100),
	)
	f := trace.NewFrame(200, 100)
	if err := f.WritePPM(os.Stdout, l); err != nil {
		panic(err)
	}
}
