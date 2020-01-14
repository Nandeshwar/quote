package api

import (
	"quote/pkg/quote"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApi(t *testing.T) {

	Convey("api test", t, func() {

		Convey("getNonReadImageTest", func() {
			Convey("getUnique files in ", func() {

				allImagesPath := []string{
					"image/competitionWithMySelf.jpg",
					"image/becomegood.jpg",
					"image/thanks-to-obstacles.jpg",
					"image/you-are-responsible.jpg",
					"image/lift-up.jpg",
				} //var imagePath string

				var imageRead []string

				for i := 0; i < len(allImagesPath); i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}

				for i := 0; i < len(allImagesPath); i++ {
					So(imageRead, ShouldContain, allImagesPath[i])
				}
				So(len(imageRead), ShouldEqual, len(allImagesPath))

				for i := 1; i <= len(allImagesPath)*2; i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}

				So(len(imageRead), ShouldEqual, len(allImagesPath))
				for i := 0; i < len(allImagesPath); i++ {
					So(imageRead, ShouldContain, allImagesPath[i])
				}

			})
		})

	})
}
