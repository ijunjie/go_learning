package map_test

import "testing"

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2[4])
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2])
	// 上面两种情况：不存在 key 时默认为 0；存在 key 且值为 0
	m[3] = 2
	if v, f := m[3]; f {
		t.Logf("3 exists, value is %d", v)
	} else {
		t.Log("3 not exist")
	}
}

func TestTraverse(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
