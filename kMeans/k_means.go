package kMeans

import (
	"github.com/sta-golang/Machine-Learning/common"
	"gonum.org/v1/gonum/stat"
	"math"
	"math/rand"
)

const (
	maxAnomalyValue = 5000
	minAnomalyValue = 1

	leftIndex = 0
	rightIndex = 1

	inf = 0x3f3f3f3f
)

type kMeans struct {
	clusters [][]float64
}

func New(data[][]float64, k, maxIters int) *kMeans {
	if len(data) <=0  {
		return nil
	}
	ret := &kMeans{}
	newData := ret.filterAnomalyValue(data, 0)
	ret.initCenters(k, len(data[0]))
	ret.fit(newData, k, maxIters)
	return ret
}

// 去除数据的异常值
func (km *kMeans) filterAnomalyValue(data [][]float64, row int) [][]float64 {
	arr := make([]float64, 0, len(data))
	for i := 0; i < len(data); i++ {
		arr = append(arr, data[i][row])
	}
	upper := stat.Mean(arr, nil) + 3 * stat.StdDev(arr, nil)
	if upper < maxAnomalyValue {
		upper = maxAnomalyValue
	}
	lower := stat.Mean(arr, nil) - 3 *stat.StdDev(arr, nil)
	if lower < minAnomalyValue {
		lower = minAnomalyValue
	}
	ret := make([][]float64, 0)
	for i := 0; i < len(data); i++ {
		if data[i][row] >= upper || data[i][row] <= lower {
			continue
		}
		ret = append(ret, data[i])
	}
	return ret
}

func (km *kMeans) initCenters(k int, dim int) {
	rand.Seed(100)
	km.clusters = make([][]float64, k)
	for i := 0; i < k; i++ {
		km.clusters[i] = make([]float64, dim)
		for j := 0; j < dim; j++ {
			km.clusters[i][j] = float64(rand.Intn(1000))
		}
	}
}

func (km *kMeans) distance(price1 []float64, price2 []float64) float64 {
	length := common.MinInt(len(price1), len(price2))
	temp := 0.0
	for i := 0; i < length; i++ {
		temp += math.Pow(price1[i] - price2[i], 2)
	}
	return math.Sqrt(temp)
}

func (km *kMeans) fit(data [][]float64, k, maxIters int) {
	clusterChanged := true
	i := 0
	centerMap := make(map[int][][]float64, 0)
	for clusterChanged {
		for j := 0; j < len(data); j++ {
			minDistance := float64(inf)
			minIndex := -1
			dataLine := data[j][leftIndex:rightIndex]
			for c, cluster := range km.clusters {
				dis := km.distance(dataLine, cluster)
				if dis < minDistance {
					minDistance = dis
					minIndex = c
				}
			}
			centerMap[minIndex] = append(centerMap[minIndex], dataLine)
		}
		newClusters := make([][]float64, len(km.clusters))
		for j := 0; j < k; j++ {
			newClusters[j] = km.clusters[j]
			if newData, ok := centerMap[j]; ok {
				newClusters[j] = common.MeansForRow(newData, 0, len(newData))
			}
		}
		if !km.hasDiff(newClusters) || i > maxIters {
			clusterChanged = false
		} else {
			km.clusters = newClusters
			i++
			centerMap = make(map[int][][]float64, 0)
		}
	}
}

func (km *kMeans) Predict() {

}

func (km *kMeans) hasDiff(otherClusters [][]float64) bool {
	if len(otherClusters) != len(km.clusters) {
		return true
	}
	for i := 0; i < len(otherClusters); i++ {
		if km.distance(otherClusters[i], km.clusters[i]) != 0 {
			return true
		}
	}
	return false
}