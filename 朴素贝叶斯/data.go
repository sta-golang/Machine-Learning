package naive_bayesian

func GetData() ([][]float64, []interface{}) {
	return TestData(),TestLabels()
}

func TestData() [][]float64 {
	data := [][]float64{
		{320,204,198,265},
		{253,53,15,2243},
		{53,32,5,325},
		{63,50,42,98},
		{1302,523,202,5430},
		{32,22,5,143},
		{105,85,70,332},
		{872,730,840,2762},
		{16,15,13,52},
		{92,70,21,693},
	}
	return data
}

func TestLabels() []interface{} {
	return []interface{} {
		1,0,0,1,0,0,1,1,1,0,
	}
}
