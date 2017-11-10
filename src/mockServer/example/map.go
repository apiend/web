/**
Go Map介绍
Go 中 Map是一种无序的键值对的集合。Map最重要的一点是通过key来快速检索数据，
key类似于索引，指向数据的值。Map是一种集合，所以我们可以像迭代数组和切片那样迭代它。
不过，Map是无序的，我们无法决定它的返回顺序，这是因为Map是使用链式hash表来实现的。

c++中的实现
在C++ STL 中map 采用红黑树实现，可以实现有序的Map.

Go 中实现
实现原理
这个实现方法的主要的方法是用空间换取时间。通过list 和 map 两种数据结构，
保存相同的一份数据。list 用来做顺序遍历，map 用来做查找，删除操作

*/
package main

import (
	"container/list"
	"fmt"
)

type Keyer interface {
	GetKey() string
}

type MapList struct {
	dataMap  map[string]*list.Element
	dataList *list.List
}

func NewMapList() *MapList {
	return &MapList{
		dataMap:  make(map[string]*list.Element),
		dataList: list.New(),
	}
}

func (mapList *MapList) Exists(data Keyer) bool {
	_, exists := mapList.dataMap[string(data.GetKey())]
	return exists
}

func (mapList *MapList) Push(data Keyer) bool {
	if mapList.Exists(data) {
		return false
	}
	elem := mapList.dataList.PushBack(data)
	mapList.dataMap[data.GetKey()] = elem
	return true
}

func (mapList *MapList) Remove(data Keyer) {
	if !mapList.Exists(data) {
		return
	}
	mapList.dataList.Remove(mapList.dataMap[data.GetKey()])
	delete(mapList.dataMap, data.GetKey())
}

func (mapList *MapList) Size() int {
	return mapList.dataList.Len()
}

func (mapList *MapList) Walk(cb func(data Keyer)) {
	for elem := mapList.dataList.Front(); elem != nil; elem = elem.Next() {
		cb(elem.Value.(Keyer))
	}
}

type Elements struct {
	value string
}

func (e Elements) GetKey() string {
	return e.value
}

func main() {
	fmt.Println("Starting test...")
	ml := NewMapList()
	var a, b, c Keyer
	a = &Elements{"Alice"}
	b = &Elements{"Bob"}
	c = &Elements{"Conrad"}
	ml.Push(a)
	ml.Push(b)
	ml.Push(c)
	cb := func(data Keyer) {
		fmt.Println(ml.dataMap[data.GetKey()].Value.(*Elements).value)
	}
	fmt.Println("Print elements in the order of pushing:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
	ml.Remove(b)
	fmt.Println("After removing b:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
}
