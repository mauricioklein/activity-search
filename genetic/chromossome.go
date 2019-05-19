package genetic

import (
	"math"
	"math/rand"
	"time"
)

// Chromossome is a subject in the population
type Chromossome struct {
	Activities []Activity
	fitness    float64
}

// UpdateFitness calculate the fitness of the chromossome considering the
// target budget and duration
func (c *Chromossome) UpdateFitness(targetDuration int, targetPrice int, ranges Ranges) {
	duration := 0
	price := 0

	for _, act := range c.Activities {
		duration += act.Duration
		price += act.Price
	}

	normalizedDuration := normalize(0, targetDuration, duration)
	normalizedPrice := normalize(0, targetPrice, price)

	durationScore := 1 - math.Abs(float64(normalizedDuration-1.0))
	priceScore := 1 - math.Abs(float64(normalizedPrice-1.0))

	c.fitness = durationScore * priceScore
}

// Crossover generates a new chromossome using the first half of the
// genetic material of the mother (caller) and the second half of the
// father (argument)
func (c *Chromossome) Crossover(other *Chromossome) *Chromossome {
	motherMaterial := c.Activities[:2]
	fatherMaterial := other.Activities[2:]

	return &Chromossome{
		Activities: append(motherMaterial, fatherMaterial...),
	}
}

// Mutation changes a random activity on the chromossome's
// genetic material by a random activity
func (c *Chromossome) Mutation(act Activity) {
	rand.Seed(time.Now().UnixNano())

	idx := rand.Intn(len(c.Activities))

	c.Activities[idx] = act
}

// Chromossomes is a slice of Chromossome
type Chromossomes []*Chromossome

// UpdateFitness triggers the UpdateFitness in all chromossomes
func (cs Chromossomes) UpdateFitness(targetDuration int, targetPrice int, ranges Ranges) {
	for _, c := range cs {
		c.UpdateFitness(targetDuration, targetPrice, ranges)
	}
}

// ByFitness sorts the chromossomes in an
// ascending order of fitness
type ByFitness Chromossomes

func (bf ByFitness) Len() int           { return len(bf) }
func (bf ByFitness) Swap(i, j int)      { bf[i], bf[j] = bf[j], bf[i] }
func (bf ByFitness) Less(i, j int) bool { return bf[i].fitness < bf[j].fitness }
