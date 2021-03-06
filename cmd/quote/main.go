package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"quote/pkg/service/emailservice"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/logic-building/functional-go/fp"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	"quote/pkg/api"
	"quote/pkg/env"
	"quote/pkg/log"
	"quote/pkg/repo"
	"quote/pkg/service"
	"quote/pkg/service/quote"
)

func main() {
	log.Init()

	// go tool pprof http://localhost:8080/debug/pprof/heap
	// go tool pprof http://localhost:8080/debug/pprof/profile
	go http.ListenAndServe(":8080", nil)

	var (
		httpPort = env.GetIntWithDefault("HTTP_PORT", 1922)
		grpcPort = env.GetIntWithDefault("GRPC_PORT", 1923)

		logLevel = env.GetLogLevelWithDefault("LOG_LEVEL", logrus.InfoLevel)

		sessionExpireMinutes = env.GetIntWithDefault("SESSION_EXPIRE_MINUTES", 60)
		// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
		webSessionSecretKey = env.GetStringWithDefault("WEB_SESSION_SECRET_KEY", "super-secret-key")

		serverRunTimeInMin  = env.GetIntWithDefault("SERVER_RUN_DURATION_MIN", 5)
		serverRunTimeInHour = env.GetIntWithDefault("SERVER_RUN_DURATION_HOUR", 2)

		devotionalImageMaxSize   = env.GetStringWithDefault("DEVOTIONAL_IMAGE_MAX_SIZE", "2400:1700")
		devotionalImageMinSize   = env.GetStringWithDefault("DEVOTIONAL_IMAGE_MIN_SIZE", "700:700")
		motivationalImageMaxSize = env.GetStringWithDefault("MOTIVATIONAL_IMAGE_MAX_SIZE", "2800:1700")
		motivationalImageMinSize = env.GetStringWithDefault("MOTIVATIONAL_IMAGE_MIN_SIZE", "700:700")

		emailServer          = env.GetStringWithDefault("EMAIL_SERVER", "smtp.gmail.com")
		emailServerPort      = env.GetIntWithDefault("EMAIL_SERVER_PORT", 587)
		emailFrom            = env.GetStringWithDefault("EMAIL_FROM", "abc@gmail.com")
		emailFromPwd         = env.GetStringWithDefault("EMAIL_FROM_PWD", "***")
		emailToForEvents     = strings.Split(env.GetStringWithDefault("EMAIL_TO_FOR_EVENTS", "abcdef@gmail.com, xyz@gmail.com"), ",")
		emailToForImageQuote = strings.Split(env.GetStringWithDefault("EMAIL_TO_FOR_QUOTE_IMAGE", "abcdef@gmail.com, xyz@gmail.com"), ",")
	)

	emailToForEvents = fp.MapStr(strings.TrimSpace, emailToForEvents)

	logrus.WithField("new_log_level", logLevel).Info("Setting log level")
	logrus.SetLevel(logLevel)

	logrus.Info("Welcome to the Quote App")

	sqlite3file := env.GetStringWithDefault("SQLITE3_FILE", "./db/quote.db")

	devotionalImageMaxWidth, devotionalImageMaxHeight, err := getImageSize(devotionalImageMaxSize, "DEVOTIONAL_IMAGE_MAX_SIZE")
	if err != nil {
		logrus.Errorf(err.Error())
		os.Exit(1)
	}

	devotionalImageMinWidth, devotionalImageMinHeight, err := getImageSize(devotionalImageMinSize, "DEVOTIONAL_IMAGE_MIN_SIZE")
	if err != nil {
		logrus.Errorf(err.Error())
		os.Exit(1)
	}

	motivationalImageMaxWidth, motivationalImageMaxHeight, err := getImageSize(motivationalImageMaxSize, "MOTIVATIONAL_IMAGE_MAX_SIZE")
	if err != nil {
		logrus.Errorf(err.Error())
		os.Exit(1)
	}

	motivationalImageMinWidth, motivationalImageMinHeight, err := getImageSize(motivationalImageMinSize, "MOTIVATIONAL_IMAGE_MIN_SIZE")
	if err != nil {
		logrus.Errorf(err.Error())
		os.Exit(1)
	}

	//pic := env.GetBoolWithDefault("PIC", false)
	//img := env.GetBoolWithDefault("IMG", false)
	//img2 := env.GetBoolWithDefault("IMAGE", false)

	quoteService := quote.NewQuoteService()
	quoteForTheDay := quoteService.QuoteForTheDay()

	wordList := strings.Split(quoteForTheDay, " ")
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
	word1, word2 := "ram", "gopal"
	fmt.Printf("\"")
	fmt.Println()
	fmt.Println("Links for the next quote")
	fmt.Printf("%s", blue("http://localhost:1922/quotes-devotional\n"))
	fmt.Printf("%s", blue("http://localhost:1922/quotes-motivational\n"))
	fmt.Printf("\n%s :%s", blue(fmt.Sprintf("http://localhost:1922/search/%s|%s", word1, word2)), red("search criteria can be delimited by '|'\n"))

	sqlite3Repo, err := repo.NewSqlite3Repo(sqlite3file)
	if err != nil {
		fmt.Println("error=", err)
	}

	infoEventSerive := service.NewInfoEventService(sqlite3Repo)

	fmt.Println("   ")

	eventsDayList := fp.RangeInt(0, 7)
	for _, day := range eventsDayList {
		today := time.Now()
		futureTime := today.AddDate(0, 0, day)

		switch day {
		case 0:
			fmt.Println("--------------Events Today------------------")
			fmt.Println("")
		case 1:
			fmt.Println("--------------Events Tomorrow------------------")
			fmt.Println("")
		case 2:
			fmt.Println("--------------Events Day After Tomorrow------------------")
			fmt.Println("")
		default:
			fmt.Printf("\n--------------Events on %s------------------", futureTime.Format("Monday Jan _2, 2006"))
			fmt.Println("")
		}

		eventsInFuture, err := infoEventSerive.EventsInFuture(futureTime)
		if err != nil {
			logrus.Errorf("error while getting EventsInFuture for tomorrow, error=%v", err)
		}
		for _, event := range eventsInFuture {
			event.DisplayEvent()
		}
	}

	//if pic || img || img2 {
	//	listDir("/image")
	//	image.DisplayImage("./image/competitionWithMySelf.jpg")
	//}

	emailStatusInMemory := emailservice.EmailStatusInMemory{}
	ctx := context.Background()
	// check every hour if email has already not been sent.
	// send email of events lists
	emailService := emailservice.NewEmailQuote(
		sqlite3Repo,
		emailStatusInMemory,
		infoEventSerive,
		emailServer,
		emailServerPort,
		emailFrom,
		emailFromPwd,
		emailToForEvents,
		emailToForImageQuote)

	c := cron.New()
	// hourly
	//c.AddFunc("@hourly", func() { emailService.SendEmailForEventDetail(ctx) })
	//every minute
	//c.AddFunc("* * * * *", func() { emailService.SendEmailForEventDetail(ctx) })
	// every 5 minutes
	c.AddFunc("*/5 * * * *", func() { emailService.SendEmailForEventDetail(ctx) })
	c.Start()

	cronForQuoteImage := cron.New()
	cronForQuoteImage.AddFunc("*/5 * * * *", func() { emailService.SendEmailForQuoteImage(ctx) })
	cronForQuoteImage.Start()

	currentTime := time.Now()
	currentTime = currentTime.Add(time.Duration(serverRunTimeInHour) * time.Hour)
	currentTime = currentTime.Add(time.Duration(serverRunTimeInMin) * time.Minute)

	const layout = "Jan 2, 2006 at 3:04pm MST"

	fmt.Printf("\n\nServer will be quit in %d hour and %d minutes at %v", serverRunTimeInHour, serverRunTimeInMin, currentTime.Format(layout))
	fmt.Printf("\n or press CTRL+C or CTRL +D to exit and stop docker container - 'quote' using commands- docker ps and docker stop \n")

	imageWidth := api.ImageSize{
		DevotionalImageMaxWidth:    devotionalImageMaxWidth,
		DevotionalImageMaxHeight:   devotionalImageMaxHeight,
		DevotionalImageMinWidth:    devotionalImageMinWidth,
		DevotionalImageMinHeight:   devotionalImageMinHeight,
		MotivationalImageMaxWidth:  motivationalImageMaxWidth,
		MotivationalImageMaxHeight: motivationalImageMaxHeight,
		MotivationalImageMinWidth:  motivationalImageMinWidth,
		MotivationalImageMinHeight: motivationalImageMinHeight,
	}
	apiServer := api.NewServer(httpPort, grpcPort, imageWidth, webSessionSecretKey, sessionExpireMinutes, infoEventSerive, quoteService)
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
