package implementStackUsingQueues

// 225. Implement Stack using Queues
// 思路：利用两个queue互相之间暂存元素来拿到最后那个元素
type MyStack struct {
	q   *Queue
	tmp *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q:   InitQueue(),
		tmp: InitQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.q.Len() == 0 {
		for this.tmp.Len() > 1 {
			this.q.Push(this.tmp.Pop())
		}
		return this.tmp.Pop()
	}
	for this.q.Len() > 1 {
		this.tmp.Push(this.q.Pop())
	}
	return this.q.Pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.q.Len() == 0 {
		for this.tmp.Len() > 1 {
			this.q.Push(this.tmp.Pop())
		}
		rlt := this.tmp.Pop()
		this.q.Push(rlt)
		return rlt
	}
	for this.q.Len() > 1 {
		this.tmp.Push(this.q.Pop())
	}
	return this.q.Top()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.Len()+this.tmp.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
