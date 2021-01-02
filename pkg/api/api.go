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
	"fmt"
	"quote/pkg/newrelicwrapper"

	newrelic "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrgrpc"

	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/justinas/alice"
	"github.com/sirupsen/logrus"

	grpc2 "quote/pkg/grpcquote"
	"quote/pkg/service"
	"quote/pkg/service/quote"
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
	httpPort  int
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

	views    Views
	grpc     *grpc.Server
	grpcPort int
}

func NewServer(httpPort int, grpcPort int, imageSize ImageSize, webSessionSecretKey string, sessionExpireMinutes int, infoEventService service.InfoEventService, quoteService quote.QuoteService) *Server {

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
		httpPort:  httpPort,
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
		grpcPort:           grpcPort,
	}

	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/quotes-devotional", s.quotesDevotional)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/quotes-motivational", s.quotesMotivational)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/events", s.events)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/events/{searchText}", s.events)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/info", s.info)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/info/{searchText}", s.info)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/search/{searchText}", s.search)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/find/{searchText}", s.search)).Methods(http.MethodGet)

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Admin section
	//router.HandleFunc("/", s.index)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/login", s.login)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/", s.login)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/admin", s.admin)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/admin-info", s.adminInfo)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/admin-event-detail", s.adminEvent)).Methods(http.MethodGet, http.MethodPost)

	putInfoHandler := http.HandlerFunc(s.putInfo)
	getInfoHandler := http.HandlerFunc(s.getInfo)
	getEventByMonthHandler := http.HandlerFunc(s.getEventByMonth)

	aliceChain := alice.New(s.authenticationHandler)

	//info api using alice middleware
	//router.Handle("/api/quote/v1/info/{id}", aliceChain.Then(getInfoHandler)).Methods(http.MethodGet)
	//router.Handle("/api/quote/v1/info/{id}", aliceChain.Then(putInfoHandler)).Methods(http.MethodPut)

	// info api using new relic and alice middleware
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/api/quote/v1/info/{id}", aliceChain.Then(getInfoHandler).ServeHTTP)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/api/quote/v1/info/{id}", aliceChain.Then(putInfoHandler).ServeHTTP)).Methods(http.MethodPut)
	router.HandleFunc(newrelic.WrapHandleFunc(newrelicwrapper.NewRelicApplication, "/api/quote/v1/eventsByMonth/{month}", aliceChain.Then(getEventByMonthHandler).ServeHTTP)).Methods(http.MethodGet)

	return s
}

// https://dev.to/techschoolguru/use-grpc-interceptor-for-authorization-with-jwt-1c5h
func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,

) (interface{}, error) {
	fmt.Println(" Unary Interceptor Nandeshwar here......", info.FullMethod)
	f := nrgrpc.UnaryServerInterceptor(newrelicwrapper.NewRelicApplication)
	f(ctx, req, info, handler)
	return handler(ctx, req)
}

func (s *Server) ServeGRPC(ready chan bool) error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return fmt.Errorf("Failed to open listening port for catalog grpc on address=%d", s.grpcPort)
	}
	s.grpc = grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		//grpc.UnaryInterceptor(nrgrpc.UnaryServerInterceptor(newrelicwrapper.NewRelicApplication)),
		grpc.StreamInterceptor(nrgrpc.StreamServerInterceptor(newrelicwrapper.NewRelicApplication)))
	grpc2.RegisterEventDetailServiceGRPCServer(s.grpc, s)
	// Register reflection service on gRPC server.`
	reflection.Register(s.grpc)
	if ready != nil {
		ready <- true
	}
	if err := s.grpc.Serve(listener); err != nil {
		return fmt.Errorf("failed to start catalog server: %v", err)
	}

	return nil
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

	err = s.ServeGRPC(nil)
	if err != nil {
		log.Fatalln("Failed to serve grpc: ", err)
	}

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

func (s *Server) GracefulStop() {
	if s.grpc != nil {
		s.grpc.GracefulStop()
	}
}
