package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"strconv"
	"strings"
)

var (
	logger = log.New(os.Stdout, "[EVENT-MAKER] ", 0)

	projectID      = flag.String("project", "", "GCP Project ID")
	topicName      = flag.String("topic", "", "Name of the GCP PubSub topic")
	numOfSources   = flag.Int("sources", 1, "Number of event sources [1-n]")
	metricLabel    = flag.String("metric", "utilization", "Name of the metric label")
	metricRange    = flag.String("range", "0-100", "Numeric metric range [0-100]")
	eventFreq      = flag.String("freq", "5s", "Event frequency [5s]")
	maxNumOfErrors = flag.Int("maxErrors", 10, "Max number of errors [10]")

	errorInvalidMetricRange = errors.New("Invalid metric range format. Expected min-max (e.g. 1-10)")
)

func main() {

	flag.Parse()

	min, max := mustParseRange(*metricRange)
	freq, err := time.ParseDuration(*eventFreq)
	failOnErr(err)

	ctx := context.Background()
	q, err := newQueue(ctx, *projectID, *topicName)
	failOnErr(err)

	sendErrorCount := 0
	for {
		for d := 0; d < *numOfSources; d++ {
			data := makeEvent(fmt.Sprintf("device-%d", d), min, max)
			logger.Printf("Publishing: %v", data)
			if err := q.push(ctx, []byte(data)); err != nil {
				logger.Printf("Error on push: %v", err)
				sendErrorCount++
				if sendErrorCount > *maxNumOfErrors {
					logger.Fatalf("Too many push errors: %d", sendErrorCount)
				}
			}
		}
		time.Sleep(freq)
	}

}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mustParseRange(r string) (min, max float64) {
	rangeParts := strings.Split(r, "-")
	if len(rangeParts) != 2 {
		log.Fatal(errorInvalidMetricRange)
	}

	min, minErr := strconv.ParseFloat(rangeParts[0], 64)
	max, maxErr := strconv.ParseFloat(rangeParts[1], 64)
	if minErr != nil || maxErr != nil {
		log.Fatal(errorInvalidMetricRange)
	}

	return min, max
}
