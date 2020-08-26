package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type hero struct {
	name string
	age  int
}

type heroSlice []hero

func (hs heroSlice) Len() int {
	return len(hs)
}

func (hs heroSlice) Less(i, j int) bool {
	return hs[i].age < hs[j].age
}

func (hs heroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var intSlice = []int{0, -1, 10, 4, 53}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heros heroSlice

	for i := 0; i < 10; i++ {
		hero := hero{
			name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	for _, v := range heros {
		fmt.Printf("%v", v)
	}

	sort.Sort(heros)
	fmt.Println()
	for _, v := range heros {
		fmt.Printf("%v", v)
	}
	fmt.Println()
}
