package main

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	// Path to the PDF file
	pdfPath := "/home/admin1/Documents/golang-by-example/UploadFile/Golang.pdf"

	// Open the PDF file
	file, err := os.Open(pdfPath)
	if err != nil {
		fmt.Println("Error opening PDF file:", err)
		return
	}
	defer file.Close()

	// Get the page dimensions
	pageDims, err := api.PageDimsFile(pdfPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
        var width_mm  float64
		var height_mm float64
	// Print the paper size for each page
	for _, dims := range pageDims {
		// Convert dimensions from points to millimeters
		width_mm = dims.Width * 0.352778
		height_mm = dims.Height * 0.352778
	    
	}
	fmt.Printf("%.0f × %.0f mm\n", width_mm, height_mm)
}
