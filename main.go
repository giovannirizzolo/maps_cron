package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/giovannirizzolo/maps_cron/file"
	"github.com/giovannirizzolo/maps_cron/maps"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

var baseUrl = "https://api.mapbox.com"

func main() {

	origin := os.Getenv("ORIGIN")
	destination := os.Getenv("DESTINATION")

	origLat, origLng := maps.GetCoordinatesPair(origin)
	destLat, destLng := maps.GetCoordinatesPair(destination)

	urlCoordinates := url.PathEscape(fmt.Sprintf("%s;%s", encodeCoordinates(origLat, origLng), encodeCoordinates(destLat, destLng)))

	duration, duration_typical := getEtas(urlCoordinates)


	// test := time.Duration(duration) * time.Second

	// fmt.Printf("%s min\n", strconv.FormatFloat(math.Round(test.Minutes()), 'f', -1, 64))

	// c := cron.New(cron.WithSeconds())
	c := cron.New()

	// c.AddFunc("* */3 6-11 * 3-6 1-5", func() {
	c.AddFunc("*/2 6-10 * 3-6 1-5", func() {
	// c.AddFunc("* * * * *", func() {
		fmt.Println("Executing cron job at", time.Now().Format(time.RFC1123))

		record := file.GenerateFileRecord(origin, destination, duration, duration_typical)
		file.WriteToFile("report.txt", record)
	})

	fmt.Println("Starting cronjob at", time.Now().Format((time.RFC1123)))
	c.Start()

	select {}
}

func getEtas(urlCoordinates string) (float64, float64) {
	profile := os.Getenv("ROUTING_PROFILE")

	apiKey := os.Getenv("MAPS_API_KEY")
	apiName := "directions"

	url := fmt.Sprintf("%s/%s/v5/mapbox/%s/%s?alternatives=true&geometries=geojson&overview=full&steps=false&access_token=%s", baseUrl, apiName, profile, urlCoordinates, apiKey)
	resp, err := http.Get(url)

	// fmt.Println(url)

	if err != nil {
		fmt.Println("[getEtas]@main: Error:", err)
		return 0, 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("[getEtas]@main: Error getting route: %s", err)
		return 0, 0
	}

	var response maps.Routes

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("[getEtas]@main: Error on unmarshaling JSON:", err)
		return 0, 0
	}

	route := response.Routes[0]

	return route.Duration, route.Duration_typical

}

func encodeCoordinates(lat string, lng string) string {
	return fmt.Sprintf("%s,%s", lat, lng)
}
