package main

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	Key   int
	Value int
	Pre   *LRUCache
	Next  *LRUCache
}

func Constructor(capacity int) LRUCache {
	var lru LRUCache
	return lru
}

func (this *LRUCache) Get(key int) int {
	return -1
}

func (this *LRUCache) Put(key int, value int) {

}
