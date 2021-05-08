package main

// Importing packages
import (
	"fmt"
	"os"
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
			r := float64(i) / float64(img_width-1)
			g := float64(j) / float64(img_height-1)
			b := 0.25

			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)

			fmt.Println(ir, ig, ib)
		}
	}
	println("\nDone\n")
}
