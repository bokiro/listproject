package main

type Storage interface {
	Add(int64)
	GetAll()
	GetByIndex(int64)
	DelByIndex(int64)
	Sort(func(a, b int64) bool)
}