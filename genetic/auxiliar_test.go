package genetic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		min   int
		max   int
		given int
		want  float64
	}{
		{min: 0, max: 100, given: 80, want: 0.8},
		{min: 0, max: 200, given: 80, want: 0.4},
		{min: 0, max: 200, given: 100, want: 0.5},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Testcase %d", i), func(t *testing.T) {
			assert.Equal(t, tc.want, normalize(tc.min, tc.max, tc.given))
		})
	}
}
