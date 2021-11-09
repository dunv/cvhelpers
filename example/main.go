package main

import (
	"errors"
	"image"

	"github.com/dunv/cvhelpers"
	"gocv.io/x/gocv"
)

func main() {
	// Read
	input := gocv.IMRead("./example.jpg", gocv.IMReadColor)
	defer input.Close()

	// convert to NCHW
	nchw := cvhelpers.NHWCToNCHW(input)
	defer nchw.Close()

	// convert to NHCW
	nhwc, err := cvhelpers.NCHWToNHWC(nchw, input.Cols(), input.Rows(), input.Channels())
	if err != nil {
		panic(err)
	}
	defer nhwc.Close()

	// print region (used for debugging preprocessing)
	err = cvhelpers.PrintNCHWRegion(
		nchw,
		input.Cols(), input.Rows(), input.Channels(),
		image.Rect(200, 600, 210, 610),
	)
	if err != nil {
		panic(err)
	}

	if ok := gocv.IMWrite("./example_after_transormation.jpg", *nhwc); !ok {
		panic(errors.New("could not write"))
	}

}
