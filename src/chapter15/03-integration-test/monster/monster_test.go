package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "monster1",
		Age:   500,
		Skill: "skill1",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()錯誤")
	}

	t.Logf("monster.Store()測試成功")
}

func TestReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()

	if !res {
		t.Fatalf("monster.ReStore()錯誤")
	}

	if monster.Name != "monster1" {
		t.Fatalf("monster名稱錯誤 實際值=%v 期望值=%v", monster.Name, "monster1")
	}

	t.Logf("monster.ReStore()測試成功")
	t.Logf("monster=%v", monster)
}
