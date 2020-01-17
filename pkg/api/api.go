package api

import (
	"context"
	"fmt"
	"net/http"
	image2 "quote/pkg/image"
	"quote/pkg/quote"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/logic-building/functional-go/fp"
	"github.com/sirupsen/logrus"
)

func intro(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Coming soon"))
}

type Server struct {
	server *http.Server
	wg     sync.WaitGroup
}

func quotesAll(w http.ResponseWriter, r *http.Request) {
	allImageLen, allImages := quote.AllQuotesImage()

	imageReadList, imagePath := getNonReadImage("All Image", allImageLen, quote.AllImageRead, quote.QuoteForTheDayImage, allImages)
	quote.AllImageRead = imageReadList

	width, height := image2.GetImageDimension(imagePath)

	if width < 400 || height < 400 {
		width += 100
		height += 100
	}

	width, height = increaseImageSize(width, height, 700, 700, 100)
	width, height = reduceImageSize(width, height, 2800, 1700, 100)

	fmt.Fprintf(w, "<head>Quote for the day! <meta http-equiv='refresh' content='300' /> </head>")
	fmt.Fprintf(w, "<h1>Quote for the day!</h1>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<img src='%s' alt='gopher' style='width:%vpx;height:%vpx;'>", imagePath, width, height))
}

func quotesMotivational(w http.ResponseWriter, r *http.Request) {
	allImageLen, allImages := quote.AllMotivationalImage()
	imageReadList, imagePath := getNonReadImage("MotivationalImage", allImageLen, quote.MotivationalImageRead, quote.GetQuoteMotivationalImage, allImages)
	quote.MotivationalImageRead = imageReadList

	width, height := image2.GetImageDimension(imagePath)

	width, height = increaseImageSize(width, height, 700, 700, 100)
	width, height = reduceImageSize(width, height, 2800, 1700, 100)

	fmt.Fprintf(w, "<head>Quote for the day! <meta http-equiv='refresh' content='300' /> </head>")
	fmt.Fprintf(w, "<h1>Quote for the day!</h1>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<img src='%s' alt='gopher' style='width:%vpx;height:%vpx;'>", imagePath, width, height))
}

func reduceImageSize(width, height, maxAllowedWidth, maxAllowedHeight, reduceFactor int) (newWidth, newHeight int) {
	for {
		if width > maxAllowedWidth || height > maxAllowedHeight {
			width -= reduceFactor
			height -= reduceFactor
		} else {
			return width, height
		}
	}
}

func increaseImageSize(width, height, minAllowedWidth, minAllowedHeight, reduceFactor int) (newWidth, newHeight int) {
	for {
		if width < minAllowedWidth || height < minAllowedHeight {
			width += reduceFactor
			height += reduceFactor
		} else {
			return width, height
		}
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

func NewServer(httpPort int) *Server {
	router := mux.NewRouter()
	router.PathPrefix("/image/").Handler(http.StripPrefix("/image/", http.FileServer(http.Dir("./image"))))
	router.PathPrefix("/image-motivational/").Handler(http.StripPrefix("/image-motivational/", http.FileServer(http.Dir("./image-motivational"))))
	router.HandleFunc("/intro", intro)
	router.HandleFunc("/quotes-all", quotesAll)
	router.HandleFunc("/quotes-motivational", quotesMotivational)

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(httpPort),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1000000,
	}
	s := &Server{server: server}
	return s
}

func (s *Server) Run() error {
	s.wg.Add(1)
	var err error
	go func() {
		err = s.server.ListenAndServe()
		if err == http.ErrServerClosed {
			err = nil
		}
		s.wg.Done()
	}()
	s.wg.Wait()
	return err
}

func (s *Server) Close() {
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := s.server.Close(); err != nil {
			logrus.Infof("\nHttpServer : Service stopping : Error=%v\n", err)
		}
	}
	logrus.Info("\nHttpServer : Stopped\n")
}
