package kMeans

import "testing"

func TestNew(t *testing.T) {
	data := GetData()
	kM := New(data, 7, 200)
	kM.Predict()
}