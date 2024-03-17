package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var baseUrl = "https://api.mapbox.com"

type Feature struct {
	Id string `json:"id"`
	PlaceType []string `json:"place_type"`
	Relevance float64 `json:"relevance"`
	Property struct {
		Accuracy string `json:"accuracy"`
		MapboxId string `json:"mapbox_id"`
	} `json:"property"`
	Text string `json:"text"`
	PlaceName string `json:"place_name"`
	Center []float64 `json:"center"`
    Geometry struct {
        Type string `json:"type"`
        Coordinates []float64 `json:"coordinates"`
    } `json:"geometry"`
}

type FeatureCollection struct {
	Query []string `json:"query"`
	Features []Feature `json:"features"`
}

func main() {
	// apiKey := getApiKey()
	loadEnvModule()

	getCoordinatesPair(os.Getenv("ORIGIN"), os.Getenv("DESTINATION"))

	// c := cron.New()

	// c.AddFunc("* */3 6-11 * 3-6 1-5", func() {
	// 	fmt.Println("Executing cron job at", time.Now().Format(time.RFC1123))

	// 	getEta(apiKey)
	// })

	// fmt.Println("Starting cronjob at", time.Now().Format((time.RFC1123)))
	// c.Start()

	// select {}
}

func getApiKey() string {
	err := loadEnvModule()
	if  err != nil {
		return ""
	}
	
	apiKey := os.Getenv("MAPS_API_KEY")
	return apiKey
}

// func initClient(apiKey string) *maps.Client {
// 	c, err := maps.NewClient(maps.WithAPIKey(apiKey))

// 	if err != nil {
// 		log.Fatalf("Fatal error on client init: %s", err)
// 		return nil
// 	}

// 	return c
// }

func loadEnvModule() error {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	return nil
}

// func getEta(apiKey string) {
// 	profile := os.Getenv("ROUTING_PROFILE")
// 	apiName :=
// 	// resp, err := http.Get("%s/%s/v5/mapbox/%s/", baseUrl, apiName, profile, )
// 	return nil

// 	if err!= nil {
//         log.Fatalf("Error getting route: %s", err)
//     }
// }

func getCoordinatesPair(origin string, destination string) (string, string) {
	apiName := "geocoding/v5/mapbox.places"
	urlEncodedOrigin := url.PathEscape(origin)
	// urlEncodedDestination := url.PathEscape(destination)

	baseUrl = ""
	fmt.Println("urlEncodedOrigin", urlEncodedOrigin)

	url := fmt.Sprintf("https://api.mapbox.com/%s/%s.json?proximity=ip&access_token=%s", apiName, urlEncodedOrigin, getApiKey())

	fmt.Println("url", url)

	resp, err := http.Get(url)
	// encodeCoordinates()
	if err != nil {
        fmt.Println("Error:", err)
        return "", ""
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return "", ""
    }

	var featureCollection FeatureCollection

	err = json.Unmarshal(body, &featureCollection)

	if err!= nil {
        fmt.Println("Error:", err)
        return "", ""
    }

	originCoordinates := featureCollection.Features[0].Center[0]
	destinationCoordinates := featureCollection.Features[0].Center[1]

	fmt.Printf("%s %s", strconv.FormatFloat(originCoordinates, 'f', -1, 64), strconv.FormatFloat(destinationCoordinates, 'f', -1, 64))
	return strconv.FormatFloat(originCoordinates, 'f', -1, 64), strconv.FormatFloat(destinationCoordinates, 'f', -1, 64)

}

// func encodeCoordinates(origin string, destination string)(string,string)  {
// 	return url.QueryEscape(origin), url.QueryEscape(destination)
// }