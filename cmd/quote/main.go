package main

import (
	"fmt"
	"os"
	"path/filepath"
	"quote/pkg/api"
	"strings"
	"time"

	"github.com/gookit/color"

	//"quote/pkg/env"
	"quote/pkg/event"
	"quote/pkg/quote"
)

func main() {
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
	fmt.Println("\n")
	fmt.Println("Link for the next quote")
	fmt.Printf("%s", blue("http://localhost:9797/image\n"))

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

	fmt.Println("\nCTRL+C or CTRL +D to exit")
	const httpPort int = 9797
	apiServer := api.NewServer(httpPort)
	apiServer.Run()
	defer apiServer.Close()
}

func listDir(dirName string) {
	var files []string

	root := dirName
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
