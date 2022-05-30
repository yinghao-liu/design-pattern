package main

// 迭代器接口
type Iterator interface {
	GetNext() int
	HasMore() bool
}

// 顺序遍历器
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

// 逆序遍历器
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
