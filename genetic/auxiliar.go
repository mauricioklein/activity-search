package genetic

func normalize(min, max, given int) float64 {
	fMin, fMax, fGiven := float64(min), float64(max), float64(given)

	return (fGiven - fMin) / (fMax - fMin)
}
