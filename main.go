package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

func main() {
	rateString := os.Getenv("RATE")
	rate, err := strconv.ParseFloat(rateString, 32)
	if err != nil {
		panic("RATE could not be converted to a float")
	}

	ticker := time.NewTicker(time.Second / time.Duration(rate))

	gofakeit.Seed(time.Now().UnixNano())

	var ip, httpMethod, path, httpVersion, referrer, userAgent string
	var statusCode, bodyBytesSent int
	var timeLocal time.Time

	httpVersion = "HTTP/1.1"
	referrer = "-"

	for range ticker.C {
		timeLocal = time.Now()

		ip = gofakeit.IPv4Address()
		httpMethod = weightedHTTPMethod(50, 20)
		path = randomPath(1, 5)
		statusCode = weightedStatusCode(80)
		bodyBytesSent = realisticBytesSent(statusCode)
		userAgent = gofakeit.UserAgent()

		fmt.Printf("%s - - [%s] \"%s %s %s\" %v %v \"%s\" \"%s\"\n", ip, timeLocal.Format("02/Jan/2006:15:04:05 -0700"), httpMethod, path, httpVersion, statusCode, bodyBytesSent, referrer, userAgent)
	}
}

func realisticBytesSent(statusCode int) int {
	if statusCode != 200 {
		return gofakeit.Number(30, 120)
	}

	return gofakeit.Number(800, 3100)
}

func weightedStatusCode(percentageOk int) int {
	roll := gofakeit.Number(0, 100)
	if roll <= percentageOk {
		return 200
	}

	return gofakeit.SimpleStatusCode()
}

func weightedHTTPMethod(percentageGet, percentagePost int) string {
	if percentageGet+percentagePost >= 100 {
		panic("percentageGet and percentagePost add up to more than 100%")
	}

	roll := gofakeit.Number(0, 100)
	if roll <= percentageGet {
		return "GET"
	} else if roll <= percentagePost {
		return "POST"
	}

	return gofakeit.HTTPMethod()
}

func randomPath(min, max int) string {
	var path strings.Builder
	length := gofakeit.Number(min, max)

	path.WriteString("/")

	for i := 0; i < length; i++ {
		if i > 0 {
			path.WriteString(gofakeit.RandString([]string{"-", "-", "_", "%20", "/", "/", "/"}))
		}
		path.WriteString(gofakeit.BuzzWord())
	}

	path.WriteString(gofakeit.RandString([]string{".hmtl", ".php", ".htm", ".jpg", ".png", ".gif", ".svg", ".css", ".js"}))

	result := path.String()
	return strings.Replace(result, " ", "%20", -1)
}
