package main

// Importing packages
import (
	"fmt"
	"os"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/img"
)

// Main function
func main() {

	// Image
	img_width := 255
	img_height := 255

	// Render
	fmt.Println("P3", img_width, img_height, "255")

	lowerLeft := geom.NewVec(-2, -1, -1)
	horizontal := geom.NewVec(4, 0, 0)
	vertical := geom.NewVec(0, 2, 0)
	origin := geom.NewVec(0, 0, 0)

	for j := img_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanline remaining: %d", j)
		for i := 0; i < img_width; i++ {
			u := float64(i) / float64(img_width)
			v := float64(j) / float64(img_height)
			r := geom.NewRay(
				origin,
				lowerLeft.Plus((horizontal.Scaled(u)).Plus(vertical.Scaled(v))).ToUnit(),
			)
			col := color(r)
			ir := int(255.99 * col.R())
			ig := int(255.99 * col.G())
			ib := int(255.99 * col.B())
			fmt.Println(ir, ig, ib)
		}
	}
	println("\nDone\n")
}

func color(r geom.Ray) img.Color {
	t := 0.5 * (r.Dir.Y() + 1.0)
	white := img.NewColor(1, 1, 1).Scaled(1 - t)
	blue := img.NewColor(0.5, 0.7, 1).Scaled(t)
	return white.Plus(blue)
}
