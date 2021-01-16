package naive_bayesian

import (
	"github.com/sta-golang/Machine-Learning/common"
	"github.com/sta-golang/go-lib-utils/algorithm/data_structure/set/hashset"
	"gonum.org/v1/gonum/stat"
	"math"
)

const (
	abnormal = 1
	normal = 0

	abnormalStr = "异常"
	normalStr = "正常"
)


type naiveBayesian struct {
	alpha float64
	classP map[interface{}]float64
	classPFeature map[interface{}]map[int]node
	numFeature int
}

type node struct {
	mu float64
	sigma float64
}

func New(data [][]float64, labels []interface{}, alpha float64) *naiveBayesian {
	nb := &naiveBayesian{
		alpha:         alpha,
		classP:        make(map[interface{}]float64),
		classPFeature: make(map[interface{}]map[int]node),
		numFeature: 0,
	}
	nb.train(data, labels)
	return nb
}

// 返回平均数和标准差
func (nb *naiveBayesian) calMuAndSigma(data []float64) (float64, float64) {
	return stat.Mean(data, nil), stat.StdDev(data, nil)
}

// 训练朴素贝叶斯算法模型
func (nb *naiveBayesian) train(data [][]float64, labels []interface{}) {
	if len(data) <= 0 {
		return
	}
	numData := len(data)
	numFeature := len(data[0])
	nb.numFeature = numFeature
	// 异常用户的取值
	set := hashset.New()
	set.Add(labels...)
	nb.classP[abnormal] = (float64(nb.sumCnt(labels,abnormal)) + nb.alpha) / (float64(numData) + nb.alpha * float64(set.Size()))
	nb.classP[normal] = 1- nb.classP[abnormal]
	for _, c := range set.Iterator() {
		nb.classPFeature[c] = make(map[int]node)
		for i := 0; i < numFeature; i++ {
			feature := make([]float64, 0)
			for j := 0; j < len(data); j++ {
				if c == labels[j] {
					feature = append(feature, data[j][i])
				}
			}
			mu, sigma := nb.calMuAndSigma(feature)
			nb.classPFeature[c][i] = node{
				mu:    mu,
				sigma: sigma,
			}
		}
	}
}

func (nb *naiveBayesian) Predict(oneData []float64) (string, error) {
	if len(oneData) != nb.numFeature {
		return "", common.ParameterErr
	}
	var label interface{}
	maxP := 0.0
	for key := range nb.classP{
		labelP := nb.classP[key]
		currentP := 1.0
		featureP := nb.classPFeature[key]
		j := 0

		for fp := range featureP {
			currentP *= nb.gaussian(featureP[fp].mu, featureP[fp].sigma, oneData[j])
			j++
		}
		if currentP * labelP > maxP {
			maxP = currentP * labelP
			label = key
		}
	}
	 return nb.interpreter(label), nil
}

func (nb *naiveBayesian) interpreter(arg interface{}) string {
	ret := "未知结果"
	switch arg {
	case abnormal:
		ret = abnormalStr
	case normal:
		ret = normalStr
	}
	return ret
}

// 高斯正态分布
func (nb *naiveBayesian) gaussian(mu, sigma, x float64) float64 {
	return 1.0 / (math.Sqrt(2*math.Pi) * sigma) * math.Pow(math.E,(-math.Pow((x-mu), 2)/(2*math.Pow(sigma, 2))))
}

func (nb *naiveBayesian) sumCnt(labels []interface{}, value interface{}) int {
	cnt := 0
	for i := 0; i < len(labels); i++ {
		if labels[i] == value {
			cnt++
		}
	}
	return cnt
}

