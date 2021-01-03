package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"quote/pkg/constants"
	"quote/pkg/env"
	"quote/pkg/model"
	"quote/pkg/repo"
	"quote/pkg/service"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	sqlite3file := env.GetStringWithDefault("SQLITE3_FILE", "./db/quote.db")
	sqlite3Repo, err := repo.NewSqlite3Repo(sqlite3file)
	if err != nil {
		fmt.Println("error=", err)
	}

	infoEventSerive := service.NewInfoEventService(sqlite3Repo)
	fmt.Println(infoEventSerive)

	//insertGitaJayanti(infoEventSerive) // inserted for 100 years
	//insertRamNavami(infoEventSerive) // Inserted for 100 years
	//insertSitaJiBirthday(infoEventSerive) // Inserted for 100 years
	insertHanumanJiBirthday(infoEventSerive) // Not inserted yet
	//insertKrishnaJanmasthmi(infoEventSerive) // Not inserted yet

	//parseDrinkPanchang("https://www.drikpanchang.com/dashavatara/lord-krishna/krishna-janmashtami-date-time.html?year=2022")
}

func insertKrishnaJanmasthmi(infoEventSerive service.InfoEventService) {
	// insert 100 gita jayanti for 100 years in quote-database
	apiKrishna := "https://www.drikpanchang.com/dashavatara/lord-krishna/krishna-janmashtami-date-time.html?year=2050"
	title := "Krishna appearance day"
	info := `Krishna Janmasthmi. Krishna birthday. Krishna Ji Birthday`
	findAndInsertEvents(2021, 100, infoEventSerive, apiKrishna, title, info)
}

func insertHanumanJiBirthday(infoEventSerive service.InfoEventService) {
	// insert 100 gita jayanti for 100 years in quote-database
	apiHanumanBirthdday := "https://www.drikpanchang.com/festivals/hanuman-jayanti/hanuman-jayanti-date-time.html?year=2050"
	title := "Hanuman Jayanti"
	info := `Hanuman Ji appearance day - Hanuman ji Birthday - Hanuman Birthday`
	findAndInsertEvents(2021, 100, infoEventSerive, apiHanumanBirthdday, title, info)
}

func insertSitaJiBirthday(infoEventSerive service.InfoEventService) {
	// insert 100 gita jayanti for 100 years in quote-database
	apiSitaBirthdday := "https://www.drikpanchang.com/festivals/sita-navami/sita-navami-date-time.html?year=2050"
	title := "Sita Ji appearance day"
	info := `Sita Ji appearance day - Sita ji Birthday - Sita Birthday`
	findAndInsertEvents(2021, 100, infoEventSerive, apiSitaBirthdday, title, info)
}

func insertRamNavami(infoEventSerive service.InfoEventService) {
	// insert 100 gita jayanti for 100 years in quote-database
	apiRamNavami := "https://www.drikpanchang.com/dashavatara/rama-navami/rama-navami-date-time.html?year=2050"
	title := "Shree Ram Ji appearance day"
	info := `ram navami - ram nawami`
	findAndInsertEvents(2021, 100, infoEventSerive, apiRamNavami, title, info)
}

func insertGitaJayanti(infoEventSerive service.InfoEventService) {
	// insert 100 gita jayanti for 100 years in quote-database
	apiGita := "https://www.drikpanchang.com/festivals/gita-jayanti/gita-jayanti-date-time.html?year=2050"
	title := "Gita Jayanti"
	info := `Gita jayanti celebration
	Krishna words to Arjun on this day`
	findAndInsertEvents(2021, 100, infoEventSerive, apiGita, title, info)
}

func findAndInsertEvents(startYear, yearsInFuture int, service service.InfoEventService, apiGita, title, info string) {
	original := apiGita
	api := original
	for s := startYear; s < startYear+yearsInFuture; s++ {
		println(s)
		api = strings.ReplaceAll(api, "2050", strconv.Itoa(s))

		parsedDate, err := parseDrinkPanchang(api)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		// Friday, December 22, 2023
		dateTime, err := time.Parse("Monday, January 2, 2006", parsedDate)
		if err != nil {
			fmt.Errorf("error parsing datetime. given datetime=%v, error=%v", parsedDate, err)
		}

		eventList, err := service.GetEventDetailByYearMonthDay(dateTime.Year(), int(dateTime.Month()), dateTime.Day())
		if err != nil {
			fmt.Errorf("error getting eventdetail list for year=%d, month=%d, day=%d. error=%v", dateTime.Year(), int(dateTime.Month()), dateTime.Day(), err)
		}
		fmt.Println("parsedDate", parsedDate)
		//fmt.Println("eventList", eventList)

		// if event does not exist, save to database
		if !model.SomeEventDetail(func(event model.EventDetail) bool { return event.Title == title }, eventList) {
			eventDetailForm := model.EventDetailForm{
				Title:     title,
				Info:      info,
				EventDate: dateTime.Format(constants.DATE_FORMAT_EVENT_DATE),
				Typ:       "different",
				Link:      api,
			}

			id, err := service.CreateNewEventDetail(eventDetailForm)
			if err != nil {
				fmt.Errorf("error creating event detail. id=%v, error=%v", id, err)
			}
			fmt.Printf("\neventdetail created successfully. id=%v\n", id)
		}
		api = original
	}
}

func parseDrinkPanchang(api string) (string, error) {
	//time.Sleep(2 * time.Second)
	fmt.Println("Finding events BS")
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return "", fmt.Errorf("\nerror creating http request for URL=%s, error=%s", api, err.Error())
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error reading=%v for api=%v", err, api)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", fmt.Errorf("readError=%v for api=%v", readErr, api)
	}

	//fmt.Println(string(body))

	// Friday, December 22, 2023
	re := regexp.MustCompile("[A-Z][a-z]+, [A-Z][a-z]+ [0-9]+, [0-9][0-9][0-9][0-9]")
	result := re.FindAllString(string(body), -1)
	var foundData string
	if len(result) > 0 {
		foundData = result[0]
	} else {
		return "", fmt.Errorf("result not found for api=%v", api)
	}
	return foundData, nil
}
