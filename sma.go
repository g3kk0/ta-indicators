package taindicators

import "log"

func Sma(values []float64, period int) []float64 {
	var res []float64

	if len(values) < period {
		log.Println("not enough values to calculate")
		return res
	}

	for i := period; i <= len(values); i++ {
		res = append(res, Sum(values[i-period:i])/float64(period))
	}

	return res
}
