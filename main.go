package main

var s Storage

func main() {
	f := func(a, b int64) bool {
		return a > b
	}
	s = NewList()
	s.Add(34)
	s.Add(-6)
	s.Add(2)
	s.Add(8)
	s.Add(-2)
	s.Add(3)
	s.DelByIndex(5)
	s.GetAll()
	s.Sort(f)
	s.GetAll()
	// n := sort.Slice()
	// sort.Sort()
	s.DelByIndex(0)
}
