package main

import "fmt"

type Monkey struct {
	Name string
}

func (m *Monkey) Climbing() {
	fmt.Println(m.Name, "生來會爬樹")
}

// LittleMonkey 小猴子繼承Monkey
type LittleMonkey struct {
	Monkey
}

type BirdAble interface {
	Flying()
}

func (m *LittleMonkey) Flying() {
	fmt.Println(m.Name, "通過學習會飛翔")
}

type FishAble interface {
	Swimming()
}

func (m *LittleMonkey) Swimming() {
	fmt.Println(m.Name, "通過學習會游泳")
}

func main() {
	littleMonkey := LittleMonkey{
		Monkey{
			Name: "little monkey1",
		},
	}

	littleMonkey.Climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()
}
