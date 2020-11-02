// Package Quote QuoteAPI.
//
//     Consumes:
//		- application/xml
//     Produces:
//      - application/json
//
// swagger:meta

package api

import (
	"context"
	"github.com/justinas/alice"
	"net/http"
	"net/http/pprof"
	"quote/pkg/service"
	"quote/pkg/service/quote"
	"strconv"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

//go:generate swagger generate spec -m -o ../../swagger-ui/swagger.json

type ImageSize struct {
	DevotionalImageMaxWidth  int
	DevotionalImageMaxHeight int
	DevotionalImageMinWidth  int
	DevotionalImageMinHeight int

	MotivationalImageMaxWidth  int
	MotivationalImageMaxHeight int
	MotivationalImageMinWidth  int
	MotivationalImageMinHeight int
}

type Views struct {
	Login            string
	Admin            string
	AdminEventDetail string
	AdminInfo        string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Server struct {
	server    *http.Server
	wg        sync.WaitGroup
	imageSize ImageSize

	loginService       service.ILogin
	infoService        service.IInfo
	eventDetailService service.IEventDetail
	quoteService       quote.QuoteService

	cookieName           string
	sessionCookieStore   *sessions.CookieStore
	sessionExpireSeconds int

	views Views
}

func NewServer(httpPort int, imageSize ImageSize, webSessionSecretKey string, sessionExpireMinutes int, infoEventService service.InfoEventService, quoteService quote.QuoteService) *Server {

	router := mux.NewRouter()
	router.PathPrefix("/image/").Handler(http.StripPrefix("/image/", http.FileServer(http.Dir("./image"))))
	router.PathPrefix("/image-motivational/").Handler(http.StripPrefix("/image-motivational/", http.FileServer(http.Dir("./image-motivational"))))

	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/")))
	router.PathPrefix("/swagger-ui/").Handler(sh)

	views := Views{
		Login:            "./views/login.gtpl",
		Admin:            "./views/admin.gtpl",
		AdminEventDetail: "./views/admin-event-detail.gtpl",
		AdminInfo:        "./views/admin-info.gtpl",
	}

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(httpPort),
		Handler:        router,
		ReadTimeout:    15 * time.Minute,
		WriteTimeout:   15 * time.Minute,
		MaxHeaderBytes: 1000000,
	}
	s := &Server{
		server:    server,
		imageSize: imageSize,

		loginService:       infoEventService,
		infoService:        infoEventService,
		eventDetailService: infoEventService,
		quoteService:       quoteService,

		sessionExpireSeconds: sessionExpireMinutes * 60,

		views: views,

		cookieName:         "nandeshwar-quote-cookie",
		sessionCookieStore: sessions.NewCookieStore([]byte(webSessionSecretKey)),
	}

	router.HandleFunc("/quotes-devotional", s.quotesDevotional).Methods(http.MethodGet)
	router.HandleFunc("/quotes-motivational", s.quotesMotivational).Methods(http.MethodGet)
	router.HandleFunc("/events", s.events).Methods(http.MethodGet)
	router.HandleFunc("/events/{searchText}", s.events).Methods(http.MethodGet)
	router.HandleFunc("/info", s.info).Methods(http.MethodGet)
	router.HandleFunc("/info/{searchText}", s.info).Methods(http.MethodGet)
	router.HandleFunc("/search/{searchText}", s.search).Methods(http.MethodGet)
	router.HandleFunc("/find/{searchText}", s.search).Methods(http.MethodGet)

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Admin section
	//router.HandleFunc("/", s.index)
	router.HandleFunc("/login", s.login).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/", s.login).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/admin", s.admin).Methods(http.MethodGet)
	router.HandleFunc("/admin-info", s.adminInfo).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/admin-event-detail", s.adminEvent).Methods(http.MethodGet, http.MethodPost)

	putInfoHandler := http.HandlerFunc(s.putInfo)
	getInfoHandler := http.HandlerFunc(s.getInfo)

	aliceChain := alice.New(s.authenticationHandler)

	//info api
	router.Handle("/api/quote/v1/info/{id}", aliceChain.Then(getInfoHandler)).Methods(http.MethodGet)
	router.Handle("/api/quote/v1/info/{id}", aliceChain.Then(putInfoHandler)).Methods(http.MethodPut)

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
