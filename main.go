package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"math"
	"os"
)

type PixelImage struct {
	img       image.Image
	sp        image.Point
	rectangle image.Rectangle
}

func mapImages(images map[int]string) map[int]*PixelImage {
	files := make(map[int]*PixelImage)
	sqr := int(math.Round(math.Cbrt(float64(len(images)))))
	for k, path := range images {
		files[k] = new(PixelImage)
		fmt.Println(k, path)
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		files[k].img, _, err = image.Decode(file)
		if err != nil {
			fmt.Println(err)
		}
		x := k - 1
		if k == 0 {
			files[0].sp = image.Point{0, 0}
		} else {
			if k >= sqr {
				if k%sqr == 0 {
					files[k].sp = image.Point{0, files[x].img.Bounds().Dy()}
				} else {
					files[k].sp = image.Point{files[x].img.Bounds().Dx(), files[x].img.Bounds().Dy()}
				}
			} else {
				files[k].sp = image.Point{0, files[x].img.Bounds().Dy()}
			}
		}

		files[k].rectangle = image.Rectangle{files[k].sp, files[k].sp.Add(files[k].img.Bounds().Size())}

	}
	return files
}

func main() {
	var images = map[int]string{
		0: "/home/enju/pic/np.jpg",
		1: "/home/enju/pic/np.jpg",
		2: "/home/enju/pic/np.jpg",
		3: "/home/enju/pic/np.jpg",
		4: "/home/enju/pic/np.jpg",
		5: "/home/enju/pic/np.jpg",
		6: "/home/enju/pic/np.jpg",
		7: "/home/enju/pic/np.jpg",
		8: "/home/enju/pic/np.jpg",
	}

	files := mapImages(images)
	fmt.Println(files)

	r := image.Rectangle{image.Point{0, 0}, files[len(files)-1].rectangle.Max}

	rgba := image.NewRGBA(r)

	for _, PxI := range files {
		draw.Draw(rgba, PxI.rectangle, PxI.img, image.Point{0, 0}, draw.Src)
	}

	out, err := os.Create("./image.jpg")
	if err != nil {
		fmt.Println(err)
	}

	var opt jpeg.Options
	opt.Quality = 80

	jpeg.Encode(out, rgba, &opt)
}
