package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

//實現Interface介面

func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 決定使用什麼標準來排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age

	//修改成對Name排序
	//return hs[i].Name > hs[j].Name
}

// Swap 元素交換的方法
func (hs HeroSlice) Swap(i, j int) {
	//var temp = hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp

	//上面3句等同下面這句
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//先定義一個切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//排序
	//1. 冒泡排序...
	//2. 使用系統提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//實現對`Hero`結構體切片的結構體切片的排序：`sort.Sort(data Interface)`。
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//將hero append到heroes切片
		heroes = append(heroes, hero)
	}

	//查看排序後
	for _, hero := range heroes {
		fmt.Println(hero)
	}

	sort.Sort(heroes)
	fmt.Println("=================")

	//查看排序前
	for _, hero := range heroes {
		fmt.Println(hero)
	}
}
