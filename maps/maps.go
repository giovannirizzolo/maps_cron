package maps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"maps_cron/env"
	"net/http"
	"net/url"
	"strconv"
)

type Feature struct {
	Id        string   `json:"id"`
	PlaceType []string `json:"place_type"`
	Relevance float64  `json:"relevance"`
	Property  struct {
		Accuracy string `json:"accuracy"`
		MapboxId string `json:"mapbox_id"`
	} `json:"property"`
	Text      string    `json:"text"`
	PlaceName string    `json:"place_name"`
	Center    []float64 `json:"center"`
	Geometry  struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

type FeatureCollection struct {
	Query    []string  `json:"query"`
	Features []Feature `json:"features"`
}

type Route struct {
	Weight_typical   float64 `json:"weight_typical"`
	Duration_typical float64 `json:"duration_typical"`
	Weight_name      string  `json:"weight_name"`
	Weight           float64 `json:"weight"`
	Duration         float64 `json:"duration"`
	Distance         float64 `json:"distance"`
}

type Routes struct {
	Routes []Route `json:"routes"`
}

func GetCoordinatesPair(point string) (string, string) {
	apiName := "geocoding/v5/mapbox.places"
	urlEncodedPoint := url.PathEscape(point)

	url := fmt.Sprintf("https://api.mapbox.com/%s/%s.json?proximity=ip&access_token=%s", apiName, urlEncodedPoint, env.GetApiKey())
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("[GetCoordinatesPair]@maps: Error:", err)
		return "", ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[GetCoordinatesPair]@maps: Error reading response body:", err)
		return "", ""
	}
	var featureCollection FeatureCollection

	err = json.Unmarshal(body, &featureCollection)

	if err != nil {
		fmt.Println("[GetCoordinatesPair]@maps: Error on JSON Unmarshal:", err)
		return "", ""
	}

	lat := featureCollection.Features[0].Center[0]
	lng := featureCollection.Features[0].Center[1]

	return strconv.FormatFloat(lat, 'f', -1, 64), strconv.FormatFloat(lng, 'f', -1, 64)
}
