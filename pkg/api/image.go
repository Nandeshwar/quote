package api

import (
	"fmt"
	"net/http"
	"strings"

	image2 "quote/pkg/image"
	"quote/pkg/service/quote"

	"github.com/logic-building/functional-go/fp"
)

func (s *Server) quotesDevotional(w http.ResponseWriter, r *http.Request) {
	allImageLen, allImages := quote.AllQuotesImage()

	imageReadList, imagePath := getNonReadImage("All Image", allImageLen, quote.AllImageRead, quote.QuoteForTheDayImage, allImages)
	quote.AllImageRead = imageReadList

	width, height := image2.GetImageDimension(imagePath)

	width, height = increaseImageSize(width, height, s.imageSize.DevotionalImageMinWidth, s.imageSize.DevotionalImageMinHeight, 100)
	width, height = reduceImageSize(width, height, s.imageSize.DevotionalImageMaxWidth, s.imageSize.DevotionalImageMaxHeight, 100)

	fmt.Fprintf(w, "<head><meta http-equiv='refresh' content='300' /> </head>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<a href='http://localhost:1922/%s' target='_blank'><img src='%s' alt='Nandeshwar' style='width:%vpx;height:%vpx;'> </a>", imagePath, imagePath, width, height))
	w.WriteHeader(http.StatusOK)
}

func (s *Server) quotesMotivational(w http.ResponseWriter, r *http.Request) {
	allImageLen, allImages := s.quoteService.AllMotivationalImage()
	imageReadList, imagePath := getNonReadImage("MotivationalImage", allImageLen, quote.MotivationalImageRead, s.quoteService.GetQuoteMotivationalImage, allImages)
	quote.MotivationalImageRead = imageReadList

	width, height := image2.GetImageDimension(imagePath)

	width, height = increaseImageSize(width, height, s.imageSize.MotivationalImageMinWidth, s.imageSize.MotivationalImageMinHeight, 100)
	width, height = reduceImageSize(width, height, s.imageSize.MotivationalImageMaxWidth, s.imageSize.MotivationalImageMaxHeight, 100)

	fmt.Fprintf(w, "<head>Quote for the day! <meta http-equiv='refresh' content='300' /> </head>")
	fmt.Fprintf(w, "<h1>Quote for the day!</h1>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<img src='%s' alt='Nandeshwar' style='width:%vpx;height:%vpx;'>", imagePath, width, height))
}
func findImage(searchText string) []string {
	_, allImages := quote.AllQuotesImage()

	find := func(image string) bool {
		if strings.Contains(strings.ToLower(image), strings.ToLower(searchText)) {
			return true
		}
		return false
	}

	foundImageList := fp.FilterStr(find, allImages)
	return foundImageList
}

func displayImage(foundImages []string, w http.ResponseWriter) {
	fmt.Fprintf(w, "<h1>Images:</h1>")
	fmt.Fprintf(w, "<h1>Click to view image:</h1>")
	for ind, imagePath := range foundImages {
		imagePathName := strings.Split(imagePath, "/")
		fmt.Fprintf(w, fmt.Sprintf("<a href='http://localhost:1922/%s'> %d. <img src='%s' alt='%s' style='width:%vpx;height:%vpx;'></a>", imagePath, ind+1, imagePath, imagePathName[1], 400, 25))
		fmt.Fprintf(w, "</br>")
	}
}

func getNonReadImage(apiName string, allImageLen int, imageRead []string, f func([]string) string, allImages []string) (imageRead2 []string, imagePath string) {

	for {
		imagePath = f(allImages)

		if len(imageRead) >= allImageLen {
			imageRead = nil
			fmt.Printf("\nImage Cycle End for api=%s", apiName)
			imageRead = append(imageRead, imagePath)
			fmt.Printf("\nNew Image Cycle Started for api=%s", apiName)
			fmt.Printf("\n%d/%d. Image for api %s: %s", len(imageRead), allImageLen, apiName, imagePath)
			imageRead2 = append(imageRead2, imageRead...)
			return imageRead2, imagePath
		}

		if !fp.ExistsStr(imagePath, imageRead) {
			imageRead = append(imageRead, imagePath)
			fmt.Printf("\n%d/%d. Image for api %s: %s", len(imageRead), allImageLen, apiName, imagePath)
			imageRead2 = append(imageRead2, imageRead...)
			return imageRead2, imagePath
		}

	}
}
