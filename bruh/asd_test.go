package bruh

import (
	"testing"
)

type TestData struct {
	Value1        int
	Value2        int
	ExpectedValue int
}

func TestAdd(t *testing.T) {
	testDatas := []TestData{
		{1, 2, 3},
		{0, 0, 0},
		{-2, -15, -17},
		{12, 5, 17},
		{-5, 4, -1},
	}

	for _, v := range testDatas {
		result := Add(v.Value1, v.Value2)
		if result != v.ExpectedValue {
			t.Errorf(
				"Add(%d, %d) FAILED. Expected %d, but got %d\n",
				v.Value1, v.Value2, v.ExpectedValue, result,
			)
		} else {
			t.Logf(
				"Add(%d, %d) PASSED. Expected %d, got %d\n",
				v.Value1, v.Value2, v.ExpectedValue, result,
			)
		}
	}
}

func TestSubstract(t *testing.T) {
	testCases := []TestData{
		{1, 2, -1},
		{5, 2, 3},
		{-2, -5, 3},
		{-5, 3, -8},
		{0, -2, 2},
	}

	for _, v := range testCases {
		result := Substract(v.Value1, v.Value2)
		if result != v.ExpectedValue {
			t.Errorf(
				"Substract(%d, %d) FAILED, expected %d but was %d\n",
				v.Value1, v.Value2, v.ExpectedValue, result,
			)
		} else {
			t.Logf(
				"Substract(%d, %d) PASSED, got %d\n",
				v.Value1, v.Value2, v.ExpectedValue,
			)
		}
	}
}
