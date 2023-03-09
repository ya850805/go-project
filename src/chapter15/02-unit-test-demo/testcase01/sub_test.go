package testcase01

import "testing"

func TestGetSub(t *testing.T) {
	res := getSub(10, 3)

	if res != 7 {
		t.Fatalf("getSub(10, 3) 執行錯誤，期望值=%v 實際值=%v", 7, res)
	}

	t.Logf("getSub(10, 3) 執行正確")
}
