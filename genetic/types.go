package genetic

// Activity is a representation of a business activity
type Activity struct {
	ID       int
	Duration int
	Price    int
}

// Ranges define the min/max of durations and prices in all loaded activities
type Ranges struct {
	Duration MinMax
	Price    MinMax
}

// MinMax is a min/max report
type MinMax struct {
	Min int
	Max int
}
