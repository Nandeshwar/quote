package main

import (
	"fmt"
	"quote/pkg/event"
	"strings"
	"time"

	"github.com/gookit/color"

	"quote/pkg/quote"
)

func main() {
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
	fmt.Println("\n")

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

	//image.DisplayImage("./pkg/image/competitionWithMySelf.jpg")
}
