package genetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateFitness_ExactMatch(t *testing.T) {
	c := Chromossome{
		Activities: []Activity{
			{Duration: 30, Price: 50},
			{Duration: 45, Price: 80},
		},
	}

	ranges := Ranges{
		Duration: MinMax{Min: 30, Max: 45},
		Price:    MinMax{Min: 50, Max: 80},
	}

	c.UpdateFitness(75, 130, ranges)

	assert.Equal(t, 1.0, c.fitness)
}

func TestUpdateFitness_PartialMatch(t *testing.T) {
	c := Chromossome{
		Activities: []Activity{
			{Duration: 30, Price: 50},
			{Duration: 45, Price: 80},
		},
	}

	ranges := Ranges{
		Duration: MinMax{Min: 30, Max: 45},
		Price:    MinMax{Min: 50, Max: 80},
	}

	c.UpdateFitness(75, 100, ranges)

	assert.Equal(t, 0.7, c.fitness)
}

func TestCrossover(t *testing.T) {
	mother := &Chromossome{
		Activities: []Activity{
			{Duration: 1, Price: 1},
			{Duration: 2, Price: 2},
			{Duration: 3, Price: 3},
		},
	}

	father := &Chromossome{
		Activities: []Activity{
			{Duration: 4, Price: 4},
			{Duration: 5, Price: 5},
			{Duration: 6, Price: 6},
		},
	}

	child := mother.Crossover(father)

	assert.Equal(t, child.Activities, []Activity{
		mother.Activities[0],
		mother.Activities[1],
		father.Activities[2],
	})
}

func TestMutation(t *testing.T) {
	c := &Chromossome{
		Activities: []Activity{
			{Duration: 1, Price: 1},
			{Duration: 2, Price: 2},
			{Duration: 3, Price: 3},
		},
	}

	mutationAct := Activity{Duration: 4, Price: 4}

	c.Mutation(mutationAct)

	found := false
	for _, act := range c.Activities {
		if act == mutationAct {
			found = true
		}
	}

	if !found {
		t.Fail()
	}
}
