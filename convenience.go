package cvhelpers

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func PrintNCHWRegion(img gocv.Mat, w, h, c int, region image.Rectangle) error {
	inputRawData, err := img.DataPtrFloat32()
	if err != nil {
		return fmt.Errorf("could not get underlying data (%s)", err)
	}

	fmt.Println()

	for inputIndex, dataPoint := range inputRawData {
		// Figure out positions
		channel := inputIndex / (h * w)
		height := (inputIndex % (h * w)) / w
		width := (inputIndex % (h * w)) % w

		// Only write things if within region
		if height >= region.Min.Y && height < region.Max.Y &&
			width >= region.Min.X && width < region.Max.X {
			localWidth := width - region.Min.X

			// Header for each channel
			if height == region.Min.Y && width == region.Min.X {
				if channel == 0 {
					fmt.Printf(" BLUE |")
				} else if channel == 1 {
					fmt.Printf("GREEN |")
				} else if channel == 2 {
					fmt.Printf("  RED |")
				}

				for i := region.Min.X; i < region.Max.X; i++ {
					fmt.Printf("  %4d  |", i)
				}
				fmt.Println()
				fmt.Printf("-------")
				for i := region.Min.X; i < region.Max.X; i++ {
					fmt.Printf("---------")
				}
				fmt.Println()
			}

			// Prefix for each line
			if localWidth == 0 {
				fmt.Printf("%5d |", height)
			}

			// Actual value
			fmt.Printf("%8.3f ", dataPoint)

			// Newline after each row
			if localWidth == region.Max.X-region.Min.X-1 {
				fmt.Println()
			}

			// Extra space between channels
			if height == region.Max.Y-1 && width == region.Max.X-1 {
				fmt.Println()
			}
		}
	}
	return nil
}
