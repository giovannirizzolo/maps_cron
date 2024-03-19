package file

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func WriteToFile(fileName string, data string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)

	if err != nil {
		log.Fatalf("Error while writing to file: %s", err)
		return
	}
}

func GenerateFileRecord(origin string, destination string, duration float64, duration_typical float64) string {
	return fmt.Sprintf("%s-%s@%s->%s (typical):%s\n", origin, destination, time.Now().Format(time.RFC1123), formatDurationToMin(duration), formatDurationToMin(duration_typical))
}

func formatDurationToMin(duration float64) string{
	durationInSec := time.Duration(duration) * time.Second
	return fmt.Sprintf("%s min", strconv.FormatFloat(math.Round(durationInSec.Minutes()), 'f', -1, 64))
}
