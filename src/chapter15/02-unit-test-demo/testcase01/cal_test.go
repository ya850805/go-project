package testcase01

import (
	"testing"
)

// 編寫測試案例，去測試addUpper是否正確
func TestAddUpper(t *testing.T) {
	res := addUpper(10)

	if res != 55 {
		//fmt.Printf("AddUpper(10)執行錯誤，期望值=%v，實際值=%v", 55, res)
		t.Fatalf("AddUpper(10)執行錯誤，期望值=%v，實際值=%v", 55, res)
	}

	//如過正確，輸出日誌
	t.Logf("AddUpper(10) 執行正確")
}
