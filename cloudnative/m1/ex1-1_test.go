package m1

import "testing"

func TestArr(t *testing.T) {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	for i, _ := range arr {
		if i == 2 {
			arr[i] = "smart"
		}
		if i == 4 {
			arr[i] = "strong"
		}
	}
	t.Logf("arr=%v\n", arr)
}
