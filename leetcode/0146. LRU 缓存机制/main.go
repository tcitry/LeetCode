package main

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	node := &LRUCache{
		capacity: capacity,
		head:     &DLinkedNode{0, 0, nil, nil},
		tail:     &DLinkedNode{0, 0, nil, nil},
		cache:    map[int]*DLinkedNode{},
	}
	return *node
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	ret := node.value
	return ret
}

func (this *LRUCache) Put(key int, value int) {

}
