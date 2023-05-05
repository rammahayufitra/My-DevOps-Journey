package math_test

import "testing"


func TestAdd(t *testing.T){
	result := Add(1, 2)
	if result != 3 {
		t.Logf("Got: %d But expect %d", result, 3)
		t.Fail()
	}
}

func TestSubtract(t *testing.T){
	result :=  Subtract(1, 2)
	if result != -1 {
		t.Logf("Got: %d But expect %d", result, -1)
		t.Fail()
	}
}