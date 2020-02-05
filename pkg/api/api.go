package api

import (
	"context"
	"fmt"
	"net/http"
	"quote/pkg/event"
	image2 "quote/pkg/image"
	info2 "quote/pkg/info"
	"quote/pkg/quote"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func intro(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Coming soon"))
}

type Server struct {
	server                   *http.Server
	wg                       sync.WaitGroup
	devotionalImageMaxWidth  int
	devotionalImageMaxHeight int
	devotionalImageMinWidth  int
	devotionalImageMinHeight int

	motivationalImageMaxWidth  int
	motivationalImageMaxHeight int
	motivationalImageMinWidth  int
	motivationalImageMinHeight int
}

func (s *Server) quotesAll(w http.ResponseWriter, r *http.Request) {
	allImageLen, allImages := quote.AllQuotesImage()

	imageReadList, imagePath := getNonReadImage("All Image", allImageLen, quote.AllImageRead, quote.QuoteForTheDayImage, allImages)
	quote.AllImageRead = imageReadList

	width, height := image2.GetImageDimension(imagePath)

	width, height = increaseImageSize(width, height, s.devotionalImageMinWidth, s.devotionalImageMinHeight, 100)
	width, height = reduceImageSize(width, height, s.devotionalImageMaxWidth, s.devotionalImageMaxHeight, 100)

	fmt.Fprintf(w, "<head><meta http-equiv='refresh' content='300' /> </head>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<img src='%s' alt='Nandeshwar' style='width:%vpx;height:%vpx;'>", imagePath, width, height))
}

func (s *Server) quotesMotivational(w http.ResponseWriter, r *http.Request) {
	allImageLen, allImages := quote.AllMotivationalImage()
	imageReadList, imagePath := getNonReadImage("MotivationalImage", allImageLen, quote.MotivationalImageRead, quote.GetQuoteMotivationalImage, allImages)
	quote.MotivationalImageRead = imageReadList

	width, height := image2.GetImageDimension(imagePath)

	width, height = increaseImageSize(width, height, s.motivationalImageMinWidth, s.motivationalImageMinHeight, 100)
	width, height = reduceImageSize(width, height, s.motivationalImageMaxWidth, s.motivationalImageMaxHeight, 100)

	fmt.Fprintf(w, "<head>Quote for the day! <meta http-equiv='refresh' content='300' /> </head>")
	fmt.Fprintf(w, "<h1>Quote for the day!</h1>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<img src='%s' alt='Nandeshwar' style='width:%vpx;height:%vpx;'>", imagePath, width, height))
}

func events(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	filteredEvents := findEvents(searchText)
	displayEvents(filteredEvents, w)
}

func info(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	filteredInfo := findInfo(searchText)
	displayInfo(filteredInfo, w)
}

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	var wg sync.WaitGroup

	infoCh := make(chan []info2.Info, 1)
	eventCh := make(chan []*event.EventDetail, 1)
	imageCh := make(chan []string, 1)

	wg.Add(1)
	go func(infoCh chan []info2.Info) {
		defer wg.Done()
		filteredInfo := findInfo(searchText)
		infoCh <- filteredInfo
		close(infoCh)
	}(infoCh)

	wg.Add(1)
	go func(eventCh chan []*event.EventDetail) {
		defer wg.Done()
		filteredEvents := findEvents(searchText)
		eventCh <- filteredEvents
		close(eventCh)
	}(eventCh)

	wg.Add(1)
	go func(imageCh chan []string) {
		defer wg.Done()
		foundImages := findImage(searchText)
		imageCh <- foundImages
		close(imageCh)
	}(imageCh)

	wg.Add(1)
	go func(infoCh chan []info2.Info, eventCh chan []*event.EventDetail, imageCh chan []string) {
		defer wg.Done()

		filteredInfo := <-infoCh
		filteredEvents := <-eventCh
		foundImages := <-imageCh

		fmt.Fprintf(w, "<title>Search</title>")
		fmt.Fprintf(w, fmt.Sprintf("<h1>Info: %d, Events: %d, Images: %d</h1>", len(filteredInfo), len(filteredEvents), len(foundImages)))

		displayInfo(filteredInfo, w)
		displayEvents(filteredEvents, w)
		displayImage(foundImages, w)
	}(infoCh, eventCh, imageCh)

	wg.Wait()

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

func NewServer(httpPort, devotionalImageMaxWidth, devotionalImageMaxHeight, devotionalImageMinWidth, devotionalImageMinHeight, motivationalImageMaxWidth, motivationalImageMaxHeight, motivationalImageMinWidth, motivationalImageMinHeight int) *Server {
	router := mux.NewRouter()
	router.PathPrefix("/image/").Handler(http.StripPrefix("/image/", http.FileServer(http.Dir("./image"))))
	router.PathPrefix("/image-motivational/").Handler(http.StripPrefix("/image-motivational/", http.FileServer(http.Dir("./image-motivational"))))

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(httpPort),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1000000,
	}
	s := &Server{
		server:                     server,
		devotionalImageMaxWidth:    devotionalImageMaxWidth,
		devotionalImageMaxHeight:   devotionalImageMaxHeight,
		devotionalImageMinWidth:    devotionalImageMinWidth,
		devotionalImageMinHeight:   devotionalImageMinHeight,
		motivationalImageMaxWidth:  motivationalImageMaxWidth,
		motivationalImageMaxHeight: motivationalImageMaxHeight,
		motivationalImageMinWidth:  motivationalImageMinWidth,
		motivationalImageMinHeight: motivationalImageMinHeight,
	}

	router.HandleFunc("/intro", intro)
	router.HandleFunc("/quotes-devotional", s.quotesAll)
	router.HandleFunc("/quotes-motivational", s.quotesMotivational)
	router.HandleFunc("/events", events)
	router.HandleFunc("/events/{searchText}", events)
	router.HandleFunc("/info", info)
	router.HandleFunc("/info/{searchText}", info)
	router.HandleFunc("/search/{searchText}", s.search)
	router.HandleFunc("/find/{searchText}", s.search)

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
