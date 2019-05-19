package main

// Schedule is a JSON response for schedule object
type Schedule struct {
	Summary Summary `json:"summary"`
	Days    []Day   `json:"days"`
}

// Summary is a JSON response for summary object
type Summary struct {
	BudgetSpent int `json:"budget_spent"`
	TimeSpent   int `json:"time_spent"`
}

// Day is a JSON response for day object
type Day struct {
	Ord       int   `json:"day"`
	Itinerary []Leg `json:"itinerary"`
}

// Leg is a JSON response for leg object
type Leg struct {
	StartAt string   `json:"start"`
	Act     activity `json:"activity"`
}
