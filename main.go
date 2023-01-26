package main

import (
	"fmt"
	"sort"
)

type Storage interface {
	Add(int64)
	GetAll()
	GetByIndex(int64)
	DelByIndex(int64)
	Sort(func(a, b int64) bool)
}
type Note struct {
	PrevPoint *Note
	NextPoint *Note
	Data      int64
}
type List struct {
	FirstNote *Note
	LastNote  *Note
	Len       int64
}
type Array struct {
	Arr []int64
}

var s Storage

func main() {
	f := func(a, b int64) bool {
		return a > b
	}
	s = NewArray()
	s.Add(34)
	s.Add(-6)
	s.Add(2)
	s.Add(8)
	s.Add(-2)
	s.Add(3)
	s.DelByIndex(1)
	s.GetAll()
	s.Sort(f)
	s.GetAll()
	n := sort.Slice()
	sort.Sort()
}

func NewList() (l *List) {
	l = &List{
		FirstNote: nil,
		LastNote:  nil,
		Len:       0,
	}
	return
}

func (l *List) Add(i int64) {
	if l.Len == 0 {
		l.FirstNote = &Note{
			PrevPoint: nil,
			NextPoint: nil,
			Data:      i,
		}
		l.Len++
		l.LastNote = l.FirstNote
		return
	}
	l.LastNote.NextPoint = &Note{
		PrevPoint: l.LastNote,
		NextPoint: nil,
		Data:      i,
	}
	l.Len++
	l.LastNote = l.LastNote.NextPoint
}

func (l *List) GetAll() {
	n := l.FirstNote
	for j := 0; j < int(l.Len); j++ {
		fmt.Printf("%p %v \n", n, n)
		n = n.NextPoint
	}
	fmt.Printf(`len = %v`, l.Len)
}

func (l *List) GetByIndex(index int64) {
	n := l.FirstNote
	for j := 0; j < int(index); j++ {
		fmt.Printf("%p %v \n", n, n)
	}
}

func (l *List) DelByIndex(index int64) {
	if index >= l.Len || index < 0 {
		fmt.Println("index out of range")
		return
	}
	if index == 0 && l.Len == 1 {
		l.FirstNote = nil
		l.LastNote = nil
		l.Len--
		return
	}
	if index == 0 && l.Len > 1 {
		l.FirstNote = l.FirstNote.NextPoint
		l.FirstNote.PrevPoint = nil
		l.Len--
		return
	}
	n := l.FirstNote
	for j := 0; j < int(index); j++ {
		n = n.NextPoint
	}
	n.PrevPoint.NextPoint = n.NextPoint
	n.NextPoint.PrevPoint = n.PrevPoint
	l.Len--
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

func (l *List) Sort(increase func(a, b int64) bool) {
	if l.Len <= 1 {
		fmt.Println(`nothing to sort`)
		return
	}
	var fe, ne *Note
	var wall int64
	fe, ne = l.FirstNote, l.FirstNote.NextPoint
	wall = l.Len
	for co := 0; co <= int(wall); co++ {
		for ci := 0; ci < int(wall)-1; ci++ {
			if fe.Data > ne.Data {
				fe.Data, ne.Data = ne.Data, fe.Data
			}
			fe = ne
			ne = ne.NextPoint
		}
		fe, ne = l.FirstNote, l.FirstNote.NextPoint
		wall--
	}
}
