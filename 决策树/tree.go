package decisionTree

import (
	"errors"
	"fmt"
	"github.com/sta-golang/Machine-Learning/common"
	"github.com/sta-golang/go-lib-utils/algorithm/data_structure"
	"github.com/sta-golang/go-lib-utils/algorithm/data_structure/set/hashset"
	"math"
)

type DecisionTree struct {
	featuresMap map[interface{}]int
	root *node
}


type node struct {
	Feature interface{} `json:"feature"`
	Branch map[interface{}]*node `json:"branch"`
	IsLeafNode bool `json:"is_leaf_node"`
}

func New(data[][]interface{}, features []interface{}) *DecisionTree {
	ret := &DecisionTree{}
	ret.root = ret.createTree(data, features)
	ret.featuresMap = make(map[interface{}]int, len(features) << 1)
	for i := 0; i < len(features); i++ {
		ret.featuresMap[features[i]] = i
	}
	//ret.debug()
	return ret
}

func (dt *DecisionTree) Predict(data []interface{}) (interface{}, error) {
	if len(data) != len(dt.featuresMap) {
		return nil, common.ParameterErr
	}
	curNode := dt.root
	for !curNode.IsLeafNode {
		featureIndex := dt.featuresMap[curNode.Feature]
		oneLabel := data[featureIndex]
		if val, ok := curNode.Branch[oneLabel]; ok {
			curNode = val
			continue
		}
		return nil, parameterErr
	}
	return curNode.Feature, nil
}

func (dt *DecisionTree) createTree(data[][]interface{}, features []interface{}) *node {
	labelList := make([]interface{}, 0, len(data))
	labelCount := make(map[interface{}]int, len(data) << 1)
	ret := new(node)
	for i := range data {
		vote := data[i][len(data[i]) - 1]
		if val, ok := labelCount[vote]; ok {
			labelCount[vote] = val+1
		} else {
			labelCount[vote] = 1
		}
		labelList = append(labelList, vote)
	}
	if len(labelCount) == 1 {
		for key, _ := range labelCount {

			return &node{
				Feature:    key,
				Branch:     nil,
				IsLeafNode: true,
			}
		}
	}
	if len(data[0]) == 1 {
		return dt.majorityCnt(labelList)
	}
	ret.Branch = make(map[interface{}]*node)
	bestFeature := dt.chooseBestFeatureToSplit(data)
	bestFeatureLabel := features[bestFeature]
	ret.Feature = bestFeatureLabel
	featureSet := hashset.New()
	for i := 0; i < len(data); i++ {
		featureSet.Add(data[i][bestFeature])
	}
	for _, value := range featureSet.Iterator() {
		subFeatures := make([]interface{}, 0, len(features) - 1)
		subFeatures = append(subFeatures, features[:bestFeature]...)
		subFeatures = append(subFeatures, features[bestFeature+1:]...)
		ret.Branch[value] = dt.createTree(dt.splitData(data, bestFeature,value), subFeatures)
	}
	return ret
}

func (dt *DecisionTree) ShannonEnt(data [][]interface{}) float64 {
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

func (dt *DecisionTree) splitData(data [][]interface{}, index int, value interface{}) [][]interface{} {
	ret := make([][]interface{}, 0)
	for _, feature := range data {
		temp := make([]interface{},0, len(feature) - 1)
		if feature[index] == value {
			for i := range feature {
				if i == index {
					continue
				}
				temp = append(temp, feature[i])
			}
			ret = append(ret, temp)
		}

	}

	return ret
}

func (dt *DecisionTree) chooseBestFeatureToSplit(data [][]interface{}) int {
	if len(data) <= 0 {
		return -1
	}

	featureNum := len(data[0]) - 1
	baseEntropy := dt.ShannonEnt(data)
	bestInfoGain, bestFeature := 0.0, -1
	for i := 0; i < featureNum; i++ {

		featureSet := hashset.New()
		for j := 0; j < len(data); j++ {
			featureSet.Add(data[j][i])
		}
		newEntropy := 0.0
		for _, val := range featureSet.Iterator() {
			splitD := dt.splitData(data, i, val)
			prob := float64(len(splitD)) / float64(len(data))
			newEntropy += prob * dt.ShannonEnt(splitD)
		}
		infoGain := baseEntropy - newEntropy
		if infoGain > bestInfoGain {
			bestFeature = i
			bestInfoGain = infoGain
		}
	}
	return bestFeature
}

func (dt *DecisionTree) majorityCnt(labelsList []interface{}) *node {
	if len(labelsList) == 0 {
		return nil
	}
	labelCount := map[interface{}]int{}
	for _, label := range labelsList {
		if val, ok := labelCount[label]; ok {
			labelCount[label] = val+1
			continue
		}
		labelCount[label] = 1
	}
	maxCnt := 0
	ret := &node{
		Feature:    nil,
		Branch:     nil,
		IsLeafNode: true,
	}
	for key, val := range labelCount {
		if val > maxCnt {
			maxCnt = val
			ret.Feature = key
		}
	}
	return ret
}

func (dt *DecisionTree) debug() {
	dt.show()
}

func (dt *DecisionTree) show() {
	if dt.root == nil {
		return
	}
	type showNode struct {
		n *node
		level int
	}
	q := data_structure.NewQueue()
	q.Push(showNode{
		n:     dt.root,
		level: 0,
	})
	for !q.Empty() {
		n := q.Pop().(showNode)
		fmt.Printf("level : %d  node %p : %v\n", n.level, n.n, n.n)
		for _, val := range n.n.Branch {
			q.Push(showNode{
				n:     val,
				level: n.level + 1,
			})
		}
	}
}
