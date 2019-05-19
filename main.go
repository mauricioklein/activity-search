package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mauricioklein/activity-scheduler/genetic"
)

const (
	initialPopulation       = 16000
	activitiesPerPopulation = 3
	dayDuration             = (12 * 60) // 12h in minutes
)

func main() {
	// Fetch arguments from CLI
	activitiesFiles := os.Args[1]
	budget, _ := strconv.Atoi(os.Args[2])
	numberOfDays, _ := strconv.Atoi(os.Args[3])
	budgetPerDay := budget / numberOfDays

	// Load activities and ranges
	acts, _ := loadActivities(activitiesFiles)
	ranges := toGeneticRange(acts.Metrics())

	days := make([]Day, numberOfDays)
	summaries := make([]Summary, numberOfDays)

	for i := 0; i < numberOfDays; i++ {
		// Generates the initial population
		initialPopulation := generateChromossomes(acts, initialPopulation, activitiesPerPopulation)

		// Calculates the survivor (i.e. the last chromossome alive)
		survivor := genetic.Run(initialPopulation, dayDuration, budgetPerDay, ranges)

		// Generate the days and summaries entries
		days[i], summaries[i] = toDayAndSummary(i+1, survivor)
	}

	// Json output
	jsonBody, err := json.Marshal(toSchedule(days, summaries))

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	fmt.Println(string(jsonBody))
}

func generateChromossomes(acts activities, nChromossomes int, activitiesPerChromossome int) genetic.Chromossomes {
	chromossomes := make(genetic.Chromossomes, nChromossomes)

	for i := 0; i < nChromossomes; i++ {
		chromossomes[i] = &genetic.Chromossome{
			Activities: selectGeneticActivities(acts, activitiesPerChromossome),
		}
	}

	return chromossomes
}

func selectGeneticActivities(acts activities, nActivities int) []genetic.Activity {
	subset := make([]genetic.Activity, nActivities)

	for i := 0; i < nActivities; i++ {
		rand.Seed(time.Now().UnixNano())

		idx := rand.Intn(len(acts))

		subset[i] = toGeneticActivity(acts[idx])
	}

	return subset
}

func toGeneticActivity(ac activity) genetic.Activity {
	return genetic.Activity{
		ID:       ac.ID,
		Duration: ac.Duration,
		Price:    ac.Price,
	}
}

func toGeneticRange(am activitiesMetrics) genetic.Ranges {
	return genetic.Ranges{
		Duration: genetic.MinMax{
			Min: am.MinDuration,
			Max: am.MaxDuration,
		},
		Price: genetic.MinMax{
			Min: am.MinPrice,
			Max: am.MaxPrice,
		},
	}
}

func toDayAndSummary(ord int, c *genetic.Chromossome) (Day, Summary) {
	legs := make([]Leg, len(c.Activities))
	summary := Summary{}

	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, time.UTC)

	for i, act := range c.Activities {
		legs[i] = Leg{
			StartAt: t.Format("3:04"),
			Act: activity{
				ID:       act.ID,
				Duration: act.Duration,
				Price:    act.Price,
			},
		}

		summary.TimeSpent += act.Duration
		summary.BudgetSpent += act.Price

		t = t.Add(time.Minute * time.Duration(act.Duration))
	}

	return Day{
		Ord:       ord,
		Itinerary: legs,
	}, summary
}

func toSchedule(days []Day, summaries []Summary) Schedule {
	totalSummary := Summary{}

	for _, s := range summaries {
		totalSummary.TimeSpent += s.TimeSpent
		totalSummary.BudgetSpent += s.BudgetSpent
	}

	return Schedule{
		Summary: totalSummary,
		Days:    days,
	}
}
