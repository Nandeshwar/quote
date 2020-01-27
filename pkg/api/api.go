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
	"strings"
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

func events(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	allEvents := event.AllEvents()

	var filteredEvents []*event.EventDetail

	filterBySearch := func(event *event.EventDetail) bool {

		if strings.Contains(strings.ToLower(event.Info), searchText) ||
			strings.Contains(strings.ToLower(event.Title), searchText) ||
			strings.Contains(strings.ToLower(event.URL), searchText) {
			return true
		}
		return false
	}

	if searchText != "" {
		searchText = strings.ToLower(searchText)
		filteredEvents = event.FilterEventDetailPtr(filterBySearch, allEvents)
	} else {
		filteredEvents = allEvents
	}

	fmt.Fprintf(w, "<title>Events</title>")

	fmt.Fprintf(w, fmt.Sprintf("<table border='2'>"))

	fmt.Fprintf(w, fmt.Sprintf("<tr>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Number</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Title</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Info</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Link</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Date</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Creattion Date</th>"))
	fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	for i, event := range filteredEvents {
		fmt.Fprintf(w, fmt.Sprintf("<tr>"))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d.</td>", i+1))
		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", event.Title))
		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", event.Info))

		// Display URL in different table under td
		fmt.Fprintf(w, fmt.Sprintf("<td>"))
		fmt.Fprintf(w, fmt.Sprintf("<table>"))
		for i, url := range strings.Split(event.URL, ";") {
			fmt.Fprintf(w, fmt.Sprintf("<tr><td><a href='%s'>Link%d </a></td></tr>", url, i+1))
		}
		fmt.Fprintf(w, fmt.Sprintf("</td>"))
		fmt.Fprintf(w, fmt.Sprintf("</table>"))

		fmt.Fprintf(w, fmt.Sprintf("<td>%d-%d-%d</td>", event.Year, event.Month, event.Day))
		fmt.Fprintf(w, fmt.Sprintf("<td>%v</td>", event.CreationDate))

		fmt.Fprintf(w, fmt.Sprintf("</tr>"))
		fmt.Fprintf(w, fmt.Sprintf("</br>"))

	}
	fmt.Fprintf(w, fmt.Sprintf("</table>"))
}

func info(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	allInfo := info2.GetAllInfo()

	var filteredInfo []info2.Info

	filterBySearch := func(info info2.Info) bool {

		if strings.Contains(strings.ToLower(info.Info), searchText) ||
			strings.Contains(strings.ToLower(info.Title), searchText) {
			return true
		}
		return false
	}

	if searchText != "" {
		searchText = strings.ToLower(searchText)
		filteredInfo = info2.Filter(filterBySearch, allInfo)
	} else {
		filteredInfo = allInfo
	}

	fmt.Fprintf(w, "<title>Info</title>")

	fmt.Fprintf(w, fmt.Sprintf("<table border='2'>"))

	fmt.Fprintf(w, fmt.Sprintf("<tr>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Number</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Title</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Info</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Link</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Info Added</th>"))
	fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	for i, info := range filteredInfo {
		fmt.Fprintf(w, fmt.Sprintf("<tr>"))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d.</td>", i+1))
		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", info.Title))
		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", info.Info))

		// Display URL in different table under td
		fmt.Fprintf(w, fmt.Sprintf("<td>"))
		fmt.Fprintf(w, fmt.Sprintf("<table>"))
		for i, url := range info.Link {
			fmt.Fprintf(w, fmt.Sprintf("<tr><td><a href='%s'>Link%d </a></td></tr>", url, i+1))
		}
		fmt.Fprintf(w, fmt.Sprintf("</td>"))
		fmt.Fprintf(w, fmt.Sprintf("</table>"))

		fmt.Fprintf(w, fmt.Sprintf("<td>%v</td>", info.CreationDate))
		fmt.Fprintf(w, fmt.Sprintf("</tr>"))
		fmt.Fprintf(w, fmt.Sprintf("</br>"))

	}
	fmt.Fprintf(w, fmt.Sprintf("</table>"))
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
	router.HandleFunc("/quotes-devotional", quotesAll)
	router.HandleFunc("/quotes-motivational", quotesMotivational)
	router.HandleFunc("/events", events)
	router.HandleFunc("/events/{searchText}", events)
	router.HandleFunc("/info", info)
	router.HandleFunc("/info/{searchText}", info)

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
