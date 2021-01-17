package kMeans

import (
	"encoding/csv"
	"os"
	"strconv"
)

func readCSV() [][]float64 {
	file, err := os.Open("skuid_price.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	result, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	ret := make([][]float64,  len(result) - 1)
	for i := 1; i < len(result); i++ {
		ret[i - 1] = make([]float64, len(result[i]))
		for j := 0; j < len(result[i]); j++ {
			ret[i - 1][j], err = strconv.ParseFloat(result[i][j], len(result[i][j]))
			if err != nil {
				panic(err)
			}
		}
	}
	return ret
}

func GetData() [][]float64 {
	data := readCSV()
	ret := make([][]float64, len(data))
	for i := 1; i < len(data[0]); i++ {
		for j := 0; j < len(data); j++ {
			ret[j] = append(ret[j], data[j][i])
		}
	}
	return ret
}
