package monster

import "testing"

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "九尾",
		Age:   999,
		Skill: "妖火",
	}
	err := monster.Store()
	if err != nil {
		t.Fatalf("Store err:%v", err)
	} else {
		t.Logf("Store Success")
	}
}

func TestReStore(t *testing.T) {
	var monster Monster
	err := monster.ReStore()
	if err != nil {
		t.Fatalf("Restore err:%v", err)
	} else {
		t.Logf("Restore success")
	}
}
