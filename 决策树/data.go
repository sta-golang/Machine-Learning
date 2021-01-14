package decisionTree

func GetData() ([][]interface{}, []interface{}) {
	return TestData(), TestFeatures()
}

func TestFeatures() []interface{} {
	ret := []interface{}{
		"天气","温度","湿度","风速",
	}
	return ret
}

func TestData() [][]interface{} {
	data := [][]interface{}{
		{2,2,1,0,"yes"},
		{2,2,1,1,"no"},
		{1,2,1,0,"yes"},
		{0,0,0,0,"yes"},
		{0,0,0,1,"no"},
		{1,0,0,1,"yes"},
		{2,1,1,0,"no"},
		{2,0,0,0,"yes"},
		{0,1,0,0,"yes"},
		{2,1,0,1,"yes"},
		{1,2,0,0,"no"},
		{0,1,1,1,"no"},
	}
	return data
}