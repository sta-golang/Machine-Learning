package decisionTree

import "math"

type Tree struct {
	data [][]interface{}
	features []string
}

func (t *Tree) ShannonEnt(data [][]interface{}) float64 {
	numData := len(data)
	labelCounts := map[interface{}]int{}
	for _, feature := range data {
		key := feature[len(feature) - 1]
		if val, ok := labelCounts[key]; ok {
			labelCounts[key] = val + 1
		} else {
			labelCounts[key] = 1
		}
	}
	shannonEnt := 0.0
	for _, val := range labelCounts {
		prod := float64(val) / float64(numData)
		shannonEnt -= prod * math.Log2(prod)
	}
	return shannonEnt
}