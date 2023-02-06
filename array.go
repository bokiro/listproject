package main

import "fmt"

type Array struct {
	Arr []int64
}

func NewArray() (a *Array) {
	a = &Array{
		Arr: []int64{},
	}
	return
}
func (a *Array) Add(i int64) {
	a.Arr = append(a.Arr, i)
}
func (a *Array) GetAll() {
	for _, v := range a.Arr {
		fmt.Printf(`%v `, v)
	}
	fmt.Println()
}
func (a *Array) GetByIndex(index int64) {
	fmt.Println(a.Arr[index])
}
func (a *Array) DelByIndex(index int64) {
	a.Arr = append(a.Arr[:index], a.Arr[index+1:]...)
}

func (a *Array) Sort(increase func(a, b int64) bool) {
	if len(a.Arr) <= 2 {
		fmt.Println(`nothing to sort`)
	}
	for j := 0; j < len(a.Arr); j++ {
		for i := 0; i < len(a.Arr)-j-1; i++ {
			if increase(a.Arr[i], a.Arr[i+1]) {
				a.Arr[i], a.Arr[i+1] = a.Arr[i+1], a.Arr[i]
			}
		}
	}
}