package main

// Importing packages
import (
	"fmt"
	"os"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/img"
)

// Main function
func main() {

	// Image
	img_width := 255
	img_height := 255

	// Render
	fmt.Println("P3", img_width, img_height, "255")

	for j := img_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanline remaining: %d", j)
		for i := 0; i < img_width; i++ {
			col := img.NewColor(float64(i)/float64(img_width), float64(j)/float64(img_height), 0.2)

			ir := int(255.99 * col.R())
			ig := int(255.99 * col.G())
			ib := int(255.99 * col.B())
			fmt.Println(ir, ig, ib)
		}
	}
	println("\nDone\n")
}
