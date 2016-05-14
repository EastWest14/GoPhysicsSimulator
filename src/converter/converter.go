//converter converts bitmaps into an encoded video.
package converter

import (
	"os"
	"fmt"
	"math"
	"image"
	"image/color"
	"image/png"
	"log"
	"os/exec"
)

type Circle struct {
    X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
    var dx, dy float64 = c.X - x, c.Y - y
    d := math.Sqrt(dx*dx+dy*dy) / c.R
    if d > 1 {
        return 0
    } else {
        return 255
    }
}

func createSampleImage(name string, serialNum int) {
	var w, h int = 1000, 1000
    var hw, hh float64 = float64(w / 2), float64(h / 2)
    r := 300.0*(1 - float64(serialNum)/100.0)
    θ := 2 * math.Pi / 3
    cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
    cg := &Circle{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 60}
    cb := &Circle{hw - r*math.Sin(-θ), hh - r*math.Cos(-θ), 60}

    m := image.NewRGBA(image.Rect(0, 0, w, h))
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {
            c := color.RGBA{
                cr.Brightness(float64(x), float64(y)),
                cg.Brightness(float64(x), float64(y)),
                cb.Brightness(float64(x), float64(y)),
                255,
            }
            m.Set(x, y, c)
        }
    }

    f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    png.Encode(f, m)
}

func createImageSeries(name string, num int) {
	for i := 1; i <= num; i++ {
		fullName := name + fmt.Sprintf("%05d", i) + ".png"
		createSampleImage(fullName, i)
	}
}

func createVideo(name string) {
	cmd := exec.Command("ffmpeg", "-f", "image2", "-r", "24", "-i", name + "%05d.png", "-vcodec", "mpeg4", "-y", name + ".mp4")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func createSampleVideo() {
	cmd := exec.Command("ffmpeg", "-f", "image2", "-r", "1", "-i", "test_images/image%05d.png", "-vcodec", "mpeg4", "-y", "new_movie.mp4")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func createCircleVideo(name string) {
	createImageSeries(name, 100)
	createVideo(name)
}
