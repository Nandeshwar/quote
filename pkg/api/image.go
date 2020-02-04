package api

import (
	"fmt"
	"net/http"
	"quote/pkg/quote"
	"strings"

	"github.com/logic-building/functional-go/fp"
)

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
	for _, imagePath := range foundImages {
		imagePathName := strings.Split(imagePath, "/")
		fmt.Fprintf(w, fmt.Sprintf("<a href='http://localhost:1922/%s'> <img src='%s' alt='%s' style='width:%vpx;height:%vpx;'> click me</a>", imagePath, imagePath, imagePathName[1], 200, 200))
	}
	fmt.Fprintf(w, "</table>")
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
