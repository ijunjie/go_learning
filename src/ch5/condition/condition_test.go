package condition_test

import "testing"

func TestCondition(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	if a, err := someFunc(); err == nil {
		t.Log(a)
	} else {

	}
}

func someFunc() (string, error) {
	return "aaa", nil
}
