package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"

	"quote/pkg/api"
	"quote/pkg/env"
	"quote/pkg/event"
	"quote/pkg/quote"
)

func main() {
	serverRunTimeInMin := env.GetIntWithDefault("SERVER_RUN_DURATION_MIN", 5)
	serverRunTimeInHour := env.GetIntWithDefault("SERVER_RUN_DURATION_HOUR", 2)

	devotionalImageMaxSize := env.GetStringWithDefault("DEVOTIONAL_IMAGE_MAX_SIZE", "2400:1700")
	devotionalImageMinSize := env.GetStringWithDefault("DEVOTIONAL_IMAGE_MIN_SIZE", "700:700")
	motivationalImageMaxSize := env.GetStringWithDefault("MOTIVATIONAL_IMAGE_MAX_SIZE", "2800:1700")
	motivationalImageMinSize := env.GetStringWithDefault("MOTIVATIONAL_IMAGE_MIN_SIZE", "700:700")

	devotionalImageMaxWidth, devotionalImageMaxHeight, err := getImageSize(devotionalImageMaxSize, "DEVOTIONAL_IMAGE_MAX_SIZE")
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	devotionalImageMinWidth, devotionalImageMinHeight, err := getImageSize(devotionalImageMinSize, "DEVOTIONAL_IMAGE_MIN_SIZE")
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	motivationalImageMaxWidth, motivationalImageMaxHeight, err := getImageSize(motivationalImageMaxSize, "MOTIVATIONAL_IMAGE_MAX_SIZE")
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	motivationalImageMinWidth, motivationalImageMinHeight, err := getImageSize(motivationalImageMinSize, "MOTIVATIONAL_IMAGE_MIN_SIZE")
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	//pic := env.GetBoolWithDefault("PIC", false)
	//img := env.GetBoolWithDefault("IMG", false)
	//img2 := env.GetBoolWithDefault("IMAGE", false)

	quote := quote.QuoteForTheDay()

	wordList := strings.Split(quote, " ")
	wordListLen := len(wordList)

	red := color.FgRed.Render
	green := color.FgGreen.Render
	yellow := color.FgYellow.Render
	blue := color.FgBlue.Render
	cyan := color.FgCyan.Render

	fmt.Println("Quote for the day")
	fmt.Println("-----------------")
	fmt.Printf("\t\"")
	for i := 0; i < wordListLen; i++ {
		if i%15 == 0 {
			fmt.Printf("%s ", cyan(wordList[i]))
		} else if i%10 == 0 {
			fmt.Printf("%s ", blue(wordList[i]))
		} else if i%5 == 0 {
			fmt.Printf("%s ", yellow(wordList[i]))
		} else if i%2 == 0 {
			fmt.Printf("%s ", red(wordList[i]))
		} else {
			fmt.Printf("%s ", green(wordList[i]))
		}
	}
	fmt.Printf("\"")
	fmt.Println()
	fmt.Println("Link for the next quote")
	fmt.Printf("%s", blue("http://localhost:1922/quotes-devotional\n"))
	fmt.Printf("%s", blue("http://localhost:1922/quotes-motivational\n"))
	fmt.Printf("\n%s :%s", blue("http://localhost:1922/search/krishna&radha"), red("search criteria can be delimited by '&'\n"))

	todayEvents := event.TodayEvents()
	if len(todayEvents) > 0 {
		today := time.Now()
		todayDateStr := today.Format("Mon 2006-01-2")
		fmt.Printf("\nToday %v is auspicious day\n", todayDateStr)
		fmt.Println("-------------")
	}

	for _, event := range todayEvents {
		event.DisplayEvent()
	}

	//if pic || img || img2 {
	//	listDir("/image")
	//	image.DisplayImage("./image/competitionWithMySelf.jpg")
	//}

	currentTime := time.Now()
	currentTime = currentTime.Add(time.Duration(serverRunTimeInHour) * time.Hour)
	currentTime = currentTime.Add(time.Duration(serverRunTimeInMin) * time.Minute)

	const layout = "Jan 2, 2006 at 3:04pm MST"

	fmt.Printf("\n\nServer will be quit in %d hour and %d minutes at %v", serverRunTimeInHour, serverRunTimeInMin, currentTime.Format(layout))
	fmt.Printf("\n or press CTRL+C or CTRL +D to exit and stop docker container - 'quote' using commands- docker ps and docker stop \n")
	const httpPort int = 1922
	apiServer := api.NewServer(httpPort, devotionalImageMaxWidth, devotionalImageMaxHeight, devotionalImageMinWidth, devotionalImageMinHeight, motivationalImageMaxWidth, motivationalImageMaxHeight, motivationalImageMinWidth, motivationalImageMinHeight)
	go func() {
		apiServer.Run()
	}()

	time.Sleep(time.Duration(serverRunTimeInMin) * time.Minute)
	time.Sleep(time.Duration(serverRunTimeInHour) * time.Hour)
	apiServer.Close()
}

func getImageSize(imageSize, envVarName string) (width, height int, err error) {
	if !strings.Contains(imageSize, ":") {
		return 0, 0, fmt.Errorf("wrong format of %s=%s. width and height must be separated by cololn(:)", envVarName, imageSize)
	}
	imageSizes := strings.Split(imageSize, ":")

	if len(imageSizes) == 1 {
		return 0, 0, fmt.Errorf("wrong format of %s=%s. only only value is provided. Provide width and height separate by colon(:)", envVarName, imageSize)
	} else if len(imageSizes) > 2 {
		return 0, 0, fmt.Errorf("wrong format of %s=%s. extra value is provided. Provide width and height separate by colon(:)", envVarName, imageSize)
	}
	width, err = strconv.Atoi(strings.TrimSpace(imageSizes[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("wrong value of %s=%s. 1st value width must be integer", envVarName, imageSize)
	}

	height, err = strconv.Atoi(strings.TrimSpace(imageSizes[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("wrong value of %s=%s. 2nd value height must be integer", envVarName, imageSize)
	}
	return width, height, nil
}
