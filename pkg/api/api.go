package api

import (
	"context"
	"net/http"
	"net/http/pprof"
	"quote/pkg/service"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type ImageWidth struct {
	DevotionalImageMaxWidth  int
	DevotionalImageMaxHeight int
	DevotionalImageMinWidth  int
	DevotionalImageMinHeight int

	MotivationalImageMaxWidth  int
	MotivationalImageMaxHeight int
	MotivationalImageMinWidth  int
	MotivationalImageMinHeight int
}

type Server struct {
	server     *http.Server
	wg         sync.WaitGroup
	imageWidth ImageWidth

	loginService service.ILogin
}

func NewServer(httpPort int, imageWidth ImageWidth, loginService service.ILogin) *Server {

	router := mux.NewRouter()
	router.PathPrefix("/image/").Handler(http.StripPrefix("/image/", http.FileServer(http.Dir("./image"))))
	router.PathPrefix("/image-motivational/").Handler(http.StripPrefix("/image-motivational/", http.FileServer(http.Dir("./image-motivational"))))

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(httpPort),
		Handler:        router,
		ReadTimeout:    15 * time.Minute,
		WriteTimeout:   15 * time.Minute,
		MaxHeaderBytes: 1000000,
	}
	s := &Server{
		server:     server,
		imageWidth: imageWidth,

		loginService: loginService,
	}

	router.HandleFunc("/quotes-devotional", s.quotesAll)
	router.HandleFunc("/quotes-motivational", s.quotesMotivational)
	router.HandleFunc("/events", events)
	router.HandleFunc("/events/{searchText}", events)
	router.HandleFunc("/info", info)
	router.HandleFunc("/info/{searchText}", info)
	router.HandleFunc("/search/{searchText}", s.search)
	router.HandleFunc("/find/{searchText}", s.search)

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Admin section
	router.HandleFunc("/login", s.login)
	router.HandleFunc("/info2", adminInfo)
	router.HandleFunc("/event2", adminEvent)

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
