package decisionTree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	//tree := &DecisionTree{}
	//data := tree.splitData(GetData(),0 ,0)
	//fmt.Println(data)
}

func TestNew(t *testing.T) {
	tree := New(GetData())
	one := []interface{}{
		2,2,1,5,
	}
	fmt.Println(tree.Predict(one))
}

