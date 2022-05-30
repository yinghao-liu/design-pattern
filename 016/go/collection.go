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

// 字符换集合
type CollectionString string
