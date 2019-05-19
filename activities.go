package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type activity struct {
	ID       int `json:"id"`
	Duration int `json:"duration"`
	Price    int `json:"price"`
}

type activitiesMetrics struct {
	MinDuration int
	MaxDuration int

	MinPrice int
	MaxPrice int
}

type activities []activity

// LoadActivities loads the activities from the JSON file
func loadActivities(filepath string) (activities, error) {
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	// Read file content
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Unmarshal
	var acts activities
	json.Unmarshal(byteValue, &acts)

	return acts, nil
}

// Calculate metrics for the activities
// (min and max duration and price)
func (acts activities) Metrics() activitiesMetrics {
	minDuration, maxDuration := acts[0].Duration, acts[0].Duration
	minPrice, maxPrice := acts[0].Price, acts[0].Price

	for _, act := range acts {
		if act.Duration < minDuration {
			minDuration = act.Duration
		}

		if act.Duration > maxDuration {
			maxDuration = act.Duration
		}

		if act.Price < minPrice {
			minPrice = act.Price
		}

		if act.Price > maxPrice {
			maxPrice = act.Price
		}
	}

	return activitiesMetrics{
		MinDuration: minDuration,
		MaxDuration: maxDuration,
		MinPrice:    minPrice,
		MaxPrice:    maxPrice,
	}
}
