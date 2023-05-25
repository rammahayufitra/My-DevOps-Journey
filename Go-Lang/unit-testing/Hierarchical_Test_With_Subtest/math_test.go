package math

import (
	"testing"
	
)


func TestAdd(t *testing.T){

	testTable := []struct {
		a, b			int
		expectedOutcome	int 
	}{
		{a: 1, b: 2, expectedOutcome: 3}, 
		{a: -1, b: 2, expectedOutcome: 1},
		{a: -1, b: -2, expectedOutcome: -3},
		{a: 0, b: 0, expectedOutcome: 0},
	}

	for _, test := range testTable{
		result := Add(test.a, test.b)
		if result != test.expectedOutcome{
			t.Logf("Got: %d But expect %d", result, test.expectedOutcome)
			t.Fail()
		}
	}
}

func TestAddWithHierarchical(t *testing.T){

	t.Run("a=positive", func(t *testing.T){
		a := 10
		t.Run("b=positive", func(t *testing.T){
			result := Add(a, 5)
			if result != 15{
				t.Logf("Got: %d But expect %d", result, 15)
				t.Fail()
			}
		})

		t.Run("b=negatif", func(t *testing.T){
			result := Add(a, -5)
			if result != 5{
				t.Logf("Got: %d But expect %d", result, 5)
				t.Fail()
			}
		})



	})

}

