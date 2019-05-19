package genetic

import (
	"sort"
)

const (
	sacrificePerGeneration = 2 // number of individuals killed by generation
)

// Run calculates the iterations in the population until a single subject
// is left alive. This is the subject with the highest fitness and, thus,
// the best result for the given parameters
func Run(chromossomes Chromossomes, targetDuration int, targetPrice int, ranges Ranges) *Chromossome {
	population := chromossomes

	for len(population) > 1 {
		// Update population fitness
		population.UpdateFitness(targetDuration, targetPrice, ranges)

		// Sort by descending fitness
		sort.Sort(sort.Reverse(ByFitness(population)))

		// Sacrifice
		population = population[:len(population)-sacrificePerGeneration]

		// Next generation
		population = generateNextPopulation(population)
	}

	return population[0]
}

func generateNextPopulation(population Chromossomes) Chromossomes {
	nextPopulationSize := len(population) / 2

	nextPopulation := make(Chromossomes, nextPopulationSize)

	for i := 0; i < nextPopulationSize; i++ {
		// Fetch mother and father
		mother := population[2*i]
		father := population[2*i+1]

		// Crossover: generate a child
		newChromossome := mother.Crossover(father)

		nextPopulation[i] = newChromossome
	}

	return nextPopulation
}
