package cvhelpers

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

// Reformat into NCWH (using an openCV library function)
func NHWCToNCHW(input gocv.Mat) gocv.Mat {
	return gocv.BlobFromImage(input, 1.0, image.Pt(0.0, 0.0), gocv.NewScalar(0.0, 0.0, 0.0, 0.0), false, false)
}

// Reformat into NHWC
func NCHWToNHWC(input gocv.Mat, w, h, c int) (*gocv.Mat, error) {
	inputRawData, err := input.DataPtrFloat32()
	if err != nil {
		return nil, fmt.Errorf("could not get underlying data (%s)", err)
	}
	output := gocv.NewMatWithSize(h, w, gocv.MatTypeCV32FC3)
	outputRawData, err := output.DataPtrFloat32()
	if err != nil {
		return nil, fmt.Errorf("could not get underlying data (%s)", err)
	}
	for inputIndex, dataPoint := range inputRawData {
		channel := inputIndex / (h * w)
		height := (inputIndex % (h * w)) / w
		width := (inputIndex % (h * w)) % w
		outputIndex := height*w*c + width*c + channel
		outputRawData[outputIndex] = dataPoint
	}
	return &output, nil
}
