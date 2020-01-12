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
	allImageLen, _ := quote.AllQuotesImage()

	var imagePath string
	for {
		if len(quote.AllImageRead) == allImageLen {
			quote.AllImageRead = nil
			imagePath = quote.QuoteForTheDayImage()
			break
		}

		imagePath = quote.QuoteForTheDayImage()
		if !fp.ExistsStr(imagePath, quote.AllImageRead) {
			quote.AllImageRead = append(quote.AllImageRead, imagePath)
			break
		}
	}

	width, height := image2.GetImageDimension(imagePath)

	if width < 400 || height < 400 {
		width += 100
		height += 100
	}

	fmt.Fprintf(w, "<h1>Quote for the day!</h1>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<img src='%s' alt='gopher' style='width:%vpx;height:%vpx;'>", imagePath, width, height))
}

func quotesMotivational(w http.ResponseWriter, r *http.Request) {
	allImageLen, _ := quote.AllMotivationalImage()

	var imagePath string
	for {
		if len(quote.MotivationalImageRead) == allImageLen {
			quote.MotivationalImageRead = nil
			imagePath = quote.QuoteMotivationalImage()
			break
		}

		imagePath = quote.QuoteMotivationalImage()
		if !fp.ExistsStr(imagePath, quote.MotivationalImageRead) {
			quote.MotivationalImageRead = append(quote.MotivationalImageRead, imagePath)
			break
		}
	}

	width, height := image2.GetImageDimension(imagePath)

	if width < 400 || height < 400 {
		width += 100
		height += 100
	}

	fmt.Fprintf(w, "<head>Quote for the day! <meta http-equiv='refresh' content='300' /> </head>")
	fmt.Fprintf(w, "<h1>Quote for the day!</h1>")
	fmt.Fprintf(w, "<title>Quote</title>")
	fmt.Fprintf(w, fmt.Sprintf("<img src='%s' alt='gopher' style='width:%vpx;height:%vpx;'>", imagePath, width, height))
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
			logrus.Info("\nHttpServer : Service stopping : Error=%v\n", err)
		}
	}
	logrus.Info("\nHttpServer : Stopped\n")
}
