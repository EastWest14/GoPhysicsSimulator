package converter

import (
	"testing"
	_ "os"
	//"fmt"
)

/*func TestCreateSampleVideo(t *testing.T) {
	createSampleVideo()
}*/

/*func TestSampleImage(t *testing.T) {
	const imageName = "sample"
	fullName := imageName + fmt.Sprintf("%05d", 1) + ".png"
	createSampleImage(fullName, 1)
	//os.Remove(fullName)
}*/

/*func TestCreateImageSeries(t *testing.T) {
	createImageSeries("seriesOne", 20)
}*/

func TestCreateFirstVideo(t *testing.T) {
	createCircleVideo("circle_first_video")
}