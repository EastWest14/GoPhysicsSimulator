//converter converts bitmaps into an encoded video.
package converter

import (
	//"os"
	"log"
	"os/exec"
)

func createSampleVideo() {
	cmd := exec.Command("ffmpeg", "-f", "image2", "-r", "1", "-i", "test_images/image%05d.png", "-vcodec", "mpeg4", "-y", "new_movie.mp4")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
