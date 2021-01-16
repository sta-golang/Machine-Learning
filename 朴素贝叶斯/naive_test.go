package naive_bayesian

import (
	"fmt"
	"github.com/grd/statistics"
	"gonum.org/v1/gonum/stat"
	"testing"
)

func TestMu(t *testing.T) {
	arr := statistics.Float64{5,4,3,2,1}
	variance := statistics.Variance(&arr)
	fmt.Println(variance)
}

func TestNew(t *testing.T) {
	data, labels := GetData()
	nb := New(data, labels, 1.0)
	one := []float64{134,84,235,349}
	fmt.Println(nb.Predict(one))
}

func TestGoNum(t *testing.T)  {
	dev := stat.StdDev([]float64{1, 2, 3, 4, 5}, nil)
	fmt.Println(dev)
}