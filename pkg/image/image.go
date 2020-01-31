package image

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// TODO: scrolling, such as when images are larger than the window.
func decode(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s: %v", filename, err)
	}
	return m, nil
}

func DisplayImage(file string) {
	log.SetFlags(0)
	driver.Main(func(s screen.Screen) {

		// TODO: view multiple images.
		src, err := decode(file)
		if err != nil {
			log.Fatal(err)
		}
		w := widget.NewSheet(widget.NewImage(src, src.Bounds()))
		if err := widget.RunWindow(s, w, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Title: "ImageView Shiny Example",
			},
		}); err != nil {
			log.Fatal(err)
		}
	})
}

func GetImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}
