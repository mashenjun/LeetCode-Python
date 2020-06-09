package lruCache

import (
	"testing"
)

type LRUCache struct {
	data map[int]*Node
	head *Node
	tail *Node
	capacity int
}

type Node struct {
	Key int
	Val int
	Next *Node
	Prev *Node
}


func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		data: make(map[int]*Node),
		head: &Node{},
		tail: &Node{},
	}
	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head
	return cache
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key];
	if !ok {
		return -1
	}
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	this.head.Next, node.Next, this.head.Next.Prev, node.Prev = node, this.head.Next, node, this.head
	return node.Val
}


func (this *LRUCache) Put(key int, value int)  {
	if old, ok := this.data[key]; !ok {
		node := &Node{Key:key, Val:value, Prev:this.head, Next:this.head.Next}
		this.data[key] = node
		this.head.Next, node.Next.Prev= node, node
	}else {
		old.Val = value

		old.Prev.Next, old.Next.Prev = old.Next, old.Prev
		this.head.Next, old.Next, this.head.Next.Prev, old.Prev = old, this.head.Next, old, this.head
	}
	if len(this.data)> this.capacity {
		curr := this.tail.Prev
		delete(this.data, curr.Key)
		this.tail.Prev, curr.Prev.Next = curr.Prev, this.tail
	}
}

func TestLRU(t *testing.T) {
	//["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	//[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	obj := Constructor(3);
	obj.Put(1,1)
	obj.Put(2,2)
	obj.Put(3,3)
	obj.Put(4,4)
	t.Log(obj.Get(4))
	t.Log(obj.Get(3))
	t.Log(obj.Get(2))
	t.Log(obj.Get(1))
	obj.Put(5,5)
	t.Log(obj.Get(1))
	t.Log(obj.Get(2))

}