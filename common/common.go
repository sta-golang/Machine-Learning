package common

import (
	"gonum.org/v1/gonum/stat"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MeansForRow(data [][]float64, left, right int) []float64 {
	if len(data) <= 0 {
		return nil
	}
	ret := make([]float64, 0)
	for j := 0;j < len(data[0]); j++ {
		arr := TakeAColumnFloat(data, j)
		mean := stat.Mean(arr, nil)
		ret = append(ret, mean)
	}
	return ret
}

func TakeAColumnFloat(data [][]float64, col int) []float64 {
	ret := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		ret[i] = data[i][col]
	}
	return ret
}