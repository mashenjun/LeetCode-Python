package designSkiplist

import (
	"math/rand"
)

type Skiplist struct {
	skipFactor int // <100 用来控制间隔多少个节点提取一个节点
	maxLevel   int // 允许的最高层数
	levelCount int // 当前的最高层数

	root *Node // 头节点
}

type Node struct {
	Val      int
	MaxLevel int     // 当前节点最高到那一层
	Forward  []*Node // 当前节点所在所有层的后继节点，len(Forward) = maxLevel, 如果不在的i层的话 Forward[i] = nil
}

func Constructor() Skiplist {
	sl := Skiplist{
		skipFactor: 49,
		maxLevel:   16,
		levelCount: 1,
		root:       nil,
	}
	return sl
}

func (sl *Skiplist) Search(target int) bool {
	if sl.root == nil {
		return false
	}
	// 关键逻辑
	// 从头节点开始，以从高到低的顺序，检查节点所在层的后继节点和目标值的大小
	// 目的是为了找到这一层中比仅仅小于目标值的节点，然后从这个节点下潜至下一层
	// 当跳出循环的时候，我们已经完成了阶梯型下潜的过程，我们可以直接检查当前节点最底层的后继节点是否等于目标值
	curr := sl.root
	for i := sl.levelCount - 1; i >= 0; i-- {
		for curr.Forward[i] != nil && curr.Forward[i].Val < target {
			curr = curr.Forward[i]
		}
	}
	if curr.Forward[0] != nil && curr.Forward[0].Val == target {
		return true
	}
	return false
}

func (sl *Skiplist) Add(num int) {
	if sl.root == nil {
		sl.root = &Node{
			Val:      num,
			MaxLevel: 1,
			Forward:  make([]*Node, sl.maxLevel),
		}
	}
	level := sl.randomLevel()
	node := &Node{
		Val:      num,
		MaxLevel: level,
		Forward:  make([]*Node, sl.maxLevel),
	}
	update := make([]*Node, level)
	for i := 0; i < level; i++ {
		update[i] = sl.root
	}
	// 关键逻辑
	// 和查询类似，从高到低，查询这个新节点被分配的层中，仅比新加值小的节点，记录在update数组中
	curr := sl.root
	for i := level - 1; i >= 0; i-- {
		for curr.Forward[i] != nil && curr.Forward[i].Val < num {
			curr = curr.Forward[i]
		}
		update[i] = curr
	}
	// 通过update数组，把新节点一次插入到每个层中
	for i := 0; i < level; i++ {
		node.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = node
	}
	// 更新跳表的层高
	if sl.levelCount < level {
		sl.levelCount = level
	}
}

func (sl *Skiplist) Erase(num int) bool {
	update := make([]*Node, sl.levelCount)
	// 关键逻辑
	// 和查询类似，从高到低，查询这个新节点被分配的层中，仅比新加值小的节点，记录在update数组中
	curr := sl.root
	for i := sl.levelCount - 1; i >= 0; i-- {
		for curr.Forward[i] != nil && curr.Forward[i].Val < num {
			curr = curr.Forward[i]
		}
		update[i] = curr
	}
	// 如果在最层中找到了匹配的节点，意味着可以删除该节点
	// 通过update数组，删除匹配的节点
	shouldDel := false
	if curr.Forward[0] != nil && curr.Forward[0].Val == num {
		shouldDel = true
		for i := sl.levelCount - 1; i >= 0; i-- {
			if update[i].Forward[i] != nil && update[i].Forward[i].Val == num {
				update[i].Forward[i] = update[i].Forward[i].Forward[i]
			}
		}
	}
	for sl.levelCount > 1 && sl.root.Forward[sl.levelCount-1] == nil {
		sl.levelCount--
	}
	return shouldDel
}

// randomLevel按照skipFactor/100的概率计算新节点被分配的层数
func (sl *Skiplist) randomLevel() int {
	level := 1

	for rand.Intn(100) > sl.skipFactor && level < sl.maxLevel {
		level++
	}
	return level
}
