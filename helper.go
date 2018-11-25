package taindicators

func Avg(values []float64) float64 {
	return Sum(values) / float64(len(values))
}

func Sum(values []float64) float64 {
	var res float64

	for _, x := range values {
		res += x
	}

	return res
}
