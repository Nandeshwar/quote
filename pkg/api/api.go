package api

import (
	"context"
	"net/http"
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
