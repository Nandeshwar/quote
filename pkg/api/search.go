package api

import (
	"fmt"
	"net/http"
	"quote/pkg/event"
	"quote/pkg/model"
	"strings"
	"sync"

	"github.com/logic-building/functional-go/fp"

	"github.com/gorilla/mux"
)

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]
	searchTextList := strings.Split(searchText, "|")

	// trim space
	searchTextList = fp.MapStr(func(searchTxt string) string {
		return strings.TrimSpace(searchTxt)
	}, searchTextList)

	searchTextList = searchIntelligence(searchTextList)

	var wg sync.WaitGroup

	infoCh := make(chan []model.Info, 1)
	eventCh := make(chan []*event.EventDetail, 1)
	imageCh := make(chan []string, 1)

	wg.Add(1)
	go func(infoCh chan []model.Info) {
		defer wg.Done()

		var filteredInfo []model.Info
		for _, searchTxt := range searchTextList {
			//if info2.Some()
			foundList, err := s.infoService.GetInfoByTitleOrInfo(searchTxt)
			if err != nil {
				fmt.Println("Error in findInfo", err)
			}

			for _, foundInfo := range foundList {

				isSame := func(info model.Info) bool {
					if foundInfo.Title == info.Title && foundInfo.CreationDate == info.CreationDate {
						return true
					}
					return false
				}

				if !model.SomeInfo(isSame, filteredInfo) {
					filteredInfo = append(filteredInfo, foundInfo)
				}
			}
		}
		infoCh <- filteredInfo
		close(infoCh)
	}(infoCh)

	wg.Add(1)
	go func(eventCh chan []*event.EventDetail) {
		defer wg.Done()
		var filteredEvents []*event.EventDetail

		for _, searchTxt := range searchTextList {
			foundList := findEvents(searchTxt)
			for _, foundEvent := range foundList {

				isTitleExist := func(event *event.EventDetail) bool {
					if event.Title == foundEvent.Title && event.Year == foundEvent.Year && event.Month == foundEvent.Month && event.Day == foundEvent.Day {
						return true
					}
					return false
				}

				if !event.SomeEventDetailPtr(isTitleExist, filteredEvents) {
					filteredEvents = append(filteredEvents, foundEvent)
				}
			}
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
		foundImages = fp.DistinctStrIgnoreCase(foundImages)
		imageCh <- foundImages
		close(imageCh)
	}(imageCh)

	wg.Wait()

	func(infoCh chan []model.Info, eventCh chan []*event.EventDetail, imageCh chan []string) {
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
		"ram":        []string{"ram", "rama", "raam", "rama", "raaama", "ram-ji", "shree-ram-ji", "shree ram ji"},
		"sita":       []string{"sita", "seeta", "sitaa", "seetaa", "setaa", "seeta"},
		"krishna":    []string{"krishna", "krisna", "kirisna", "kirishna", "krishnaa", "krisna", "krishna-ji", "krishnaji"},
		"radha":      []string{"radha", "rada", "raadha", "radhha", "radhaa", "radhe", "radhaji", "radha ji", "radha-ji"},
		"hanuman":    []string{"hanuman", "hanumanji", "hanuman-ji", "hanumaan", "hanman", "hanmaan", "hanmanji", "hanman ji", "hanumaan-ji"},
		"vedbyas":    []string{"ved vyas", "vedvyas", "ved byas", "vedbyas", "vyasbed", "vyas ved", "vedh vyas", "veddhvyas"},
		"sukdev":     []string{"sukdev", "suk dev", "shukdev", "shuk dev", "sutji", "sutji maharaj", "shutji", "shut ji", "sut ji", "sut ji maharaj", "shut ji maharaj", "sut", "shut"},
		"kripaluji":  []string{"kripalu", "kripaluji", "kripalu ji", "kripalu ji maharaj", "kripaluji maharaj", "ram kripalu", "ramkripalu", "ram kripalu tripathhi", "kripalu-ji", "maharaj ji", "mahrajji", "kripaalu", "kreepalu", "krepalu", "kirpalu", "kerpalu"},
		"dashrathji": []string{"dasrat", "dashrath", "dashrat", "dasrath", "dasrat ji", "dashrath ji", "dashrat ji", "dasrath ji", "dasratji", "dashrathji", "dashratji", "dasrathji"},
		"nandeshwar": []string{"nandeshwar blog", "my blog", "nandeshwar meditation", "my meditation", "my  meditation", "my  blog"},
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
