package api

import (
	"fmt"
	"net/http"
	"quote/pkg/event"
	info2 "quote/pkg/info"
	"strings"
	"sync"

	"github.com/logic-building/functional-go/fp"

	"github.com/gorilla/mux"
)

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]
	searchTextList := strings.Split(searchText, "&")

	searchTextList = searchIntelligence(searchTextList)

	var wg sync.WaitGroup

	infoCh := make(chan []info2.Info, 1)
	eventCh := make(chan []*event.EventDetail, 1)
	imageCh := make(chan []string, 1)

	wg.Add(1)
	go func(infoCh chan []info2.Info) {
		defer wg.Done()

		var filteredInfo []info2.Info
		for _, searchTxt := range searchTextList {
			filteredInfo = append(filteredInfo, findInfo(searchTxt)...)
		}
		infoCh <- filteredInfo
		close(infoCh)
	}(infoCh)

	wg.Add(1)
	go func(eventCh chan []*event.EventDetail) {
		defer wg.Done()
		var filteredEvents []*event.EventDetail
		for _, searchTxt := range searchTextList {
			filteredEvents = append(filteredEvents, findEvents(searchTxt)...)
		}
		eventCh <- filteredEvents
		close(eventCh)
	}(eventCh)

	wg.Add(1)
	go func(imageCh chan []string) {
		defer wg.Done()
		var foundImages []string
		for _, searchTxt := range searchTextList {
			foundImages = append(foundImages, findImage(searchTxt)...)
		}
		imageCh <- foundImages
		close(imageCh)
	}(imageCh)

	wg.Wait()

	func(infoCh chan []info2.Info, eventCh chan []*event.EventDetail, imageCh chan []string) {
		filteredInfo := <-infoCh
		filteredEvents := <-eventCh
		foundImages := <-imageCh

		fmt.Fprintf(w, "<title>Search</title>")
		fmt.Fprintf(w, fmt.Sprintf("<h1>Info: %d, Events: %d, Images: %d</h1>", len(filteredInfo), len(filteredEvents), len(foundImages)))

		displayInfo(filteredInfo, w)
		displayEvents(filteredEvents, w)
		displayImage(foundImages, w)
	}(infoCh, eventCh, imageCh)

}

func searchIntelligence(searchStrList []string) []string {

	var newSearchStrList []string
	m := map[string][]string{
		"ram":       []string{"ram", "rama", "raam", "rama", "raaama", "ram-ji", "shree-ram-ji", "shree ram ji"},
		"sita":      []string{"sita", "seeta", "sitaa", "seetaa", "setaa", "seeta"},
		"krishna":   []string{"krishna", "krisna", "kirisna", "kirishna", "krishnaa", "krisna", "krishna-ji", "krishnaji"},
		"radha":     []string{"radha", "rada", "raadha", "radhha", "radhaa", "radhe", "radhaji", "radha ji", "radha-ji"},
		"hanuman":   []string{"hanuman", "hanumanji", "hanuman-ji", "hanumaan", "hanman", "hanmaan", "hanmanji", "hanman ji", "hanumaan-ji"},
		"vedbyas":   []string{"ved vyas", "vedvyas", "ved byas", "vedbyas", "vyasbed", "vyas ved", "vedh vyas", "veddhvyas"},
		"sukdev":    []string{"sukdev", "suk dev", "shukdev", "shuk dev", "sutji", "sutji maharaj", "shutji", "shut ji", "sut ji", "sut ji maharaj", "shut ji maharaj", "sut", "shut"},
		"kripaluji": []string{"kripalu", "kripaluji", "kripalu ji", "kripalu ji maharaj", "kripaluji maharaj", "ram kripalu", "ramkripalu", "ram kripalu tripathhi", "kripalu-ji", "maharaj ji", "mahrajji", "kripaalu", "kreepalu", "krepalu", "kirpalu", "kerpalu"},
	}

	newSearchStrList = append(newSearchStrList, searchStrList...)
	for _, searchTxt := range searchStrList {
		for _, data := range m {
			if fp.ExistsStrIgnoreCase(searchTxt, data) {
				newSearchStrList = append(newSearchStrList, data...)
			}
		}
	}

	return fp.DistinctStrIgnoreCase(newSearchStrList)
}
