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
				}

				var imageRead []string

				for i := 0; i < len(allImagesPath); i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, len(allImagesPath))

				for i := 0; i < len(allImagesPath); i++ {
					So(imageRead, ShouldContain, allImagesPath[i])
				}

				for i := 1; i <= len(allImagesPath)*2; i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, len(allImagesPath))

				for i := 0; i < len(allImagesPath); i++ {
					So(imageRead, ShouldContain, allImagesPath[i])
				}

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+1; i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 1)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+2; i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 2)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+3; i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 3)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+4; i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 4)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+5; i++ {
					imageRead2, _ := getNonReadImage(len(allImagesPath), imageRead, quote.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 5)

			})
		})

	})
}
