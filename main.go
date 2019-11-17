package collage

import (
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/png"
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

func MapImages(images map[int]string) (map[int]*PixelImage, error) {
	files := make(map[int]*PixelImage)
	sqr := int(math.Round(math.Sqrt(float64(len(images)))))
	keys := sortedKeys(images)
	h := 0
	w := 0
	for _, k := range keys {
		files[k] = new(PixelImage)
		path := images[k]
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		files[k].img, _, err = image.Decode(file)
		if err != nil {
			return nil, err
		}
		x := k - 1
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

		files[k].rectangle = image.Rectangle{files[k].sp, files[k].sp.Add(files[k].img.Bounds().Size())}

	}
	return files, nil
}

func MakeNewCollage(images map[int]*PixelImage, output string, quality int) error {
	// create the sizing for the collage
	r := image.Rectangle{image.Point{0, 0}, images[len(images)-1].rectangle.Max}

	// create the new image
	rgba := image.NewRGBA(r)

	// iterate the images and add them to the collage
	for _, PxI := range images {
		draw.Draw(rgba, PxI.rectangle, PxI.img, image.Point{0, 0}, draw.Src)
	}

	out, err := os.Create(output)
	if err != nil {
		return err
	}

	var opt jpeg.Options
	opt.Quality = quality

	jpeg.Encode(out, rgba, &opt)

	return nil

}
