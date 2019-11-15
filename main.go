package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"math"
	"os"
	"sort"
)

type PixelImage struct {
	img       image.Image
	sp        image.Point
	rectangle image.Rectangle
}

func sortedKeys(m map[int]string) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func mapImages(images map[int]string) map[int]*PixelImage {
	files := make(map[int]*PixelImage)
	sqr := int(math.Round(math.Sqrt(float64(len(images)))))
	fmt.Println(len(images), math.Sqrt(float64(len(images))), float64(len(images)))
	keys := sortedKeys(images)
	h := 0
	w := 0
	for _, k := range keys {
		files[k] = new(PixelImage)
		path := images[k]
		fmt.Println(k, path, h)
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		files[k].img, _, err = image.Decode(file)
		if err != nil {
			fmt.Println(err)
		}
		x := k - 1
		fmt.Println("sqr", sqr, "- x", x, "- k", k, "- %", k%sqr)
		if k == 0 {
			files[0].sp = image.Point{0, 0}
		} else {
			if k%sqr != 0 {
				w += files[x].img.Bounds().Dy()
			}
			files[k].sp = image.Point{w, h}
		}

		if k%sqr == sqr-1 {
			h += files[x].img.Bounds().Dy()
			w = 0
		}
		fmt.Println(files[k].sp, h)

		files[k].rectangle = image.Rectangle{files[k].sp, files[k].sp.Add(files[k].img.Bounds().Size())}

	}
	return files
}

func main() {
	var images = map[int]string{
		0:  "/home/enju/pic/np.jpg",
		1:  "/home/enju/pic/np2.jpg",
		2:  "/home/enju/pic/np3.jpg",
		3:  "/home/enju/pic/np4.jpg",
		4:  "/home/enju/pic/np5.jpg",
		5:  "/home/enju/pic/np6.jpg",
		6:  "/home/enju/pic/np7.jpg",
		7:  "/home/enju/pic/np8.jpg",
		8:  "/home/enju/pic/np9.jpg",
		9:  "/home/enju/pic/np10.jpg",
		10: "/home/enju/pic/np11.jpg",
		11: "/home/enju/pic/np12.jpg",
		12: "/home/enju/pic/np13.jpg",
		13: "/home/enju/pic/np14.jpg",
		14: "/home/enju/pic/np15.jpg",
		15: "/home/enju/pic/np16.jpg",
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
