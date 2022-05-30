package main

// 集合接口
type Collector interface {
	GetIterator() Iterator
	GetReverseIterator() Iterator
}

// 切片集合
type CollectionSlice []int

func (c *CollectionSlice) GetIterator() Iterator {
	iter := new(orderIterator)
	iter.Set(c)
	return iter
}
func (c *CollectionSlice) GetReverseIterator() Iterator {
	iter := new(reverseIterator)
	iter.Set(c)
	return iter
}

/***************************迭代器****************************/
// 迭代器接口
type Iterator interface {
	GetNext() int
	HasMore() bool
}

// 顺序迭代器
type orderIterator struct {
	collection *CollectionSlice
	index      int
}

func (iter *orderIterator) Set(collect *CollectionSlice) {
	iter.collection = collect
}
func (iter *orderIterator) GetNext() int {
	next := (*iter.collection)[iter.index]
	iter.index++
	return next

}
func (iter orderIterator) HasMore() bool {
	if iter.index < len(*iter.collection) {
		return true
	}
	return false
}

// 逆序迭代器
type reverseIterator struct {
	collection *CollectionSlice
	index      int
}

func (iter *reverseIterator) Set(collect *CollectionSlice) {
	iter.collection = collect
}
func (iter reverseIterator) GetNext() int {
	return 0
}
func (iter reverseIterator) HasMore() bool {
	return false
}
