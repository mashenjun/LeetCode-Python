package implementQueueUsingStacks

// 232. Implement Queue using Stacks
type MyQueue struct {
	in  *Stack
	out *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  NewStack(),
		out: NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.out.Len() > 0 {
		return this.out.Pop()
	}
	// 把in中的元素全部灌到out中
	for this.in.Len() > 0 {
		this.out.Push(this.in.Pop())
	}
	return this.out.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.out.Len() > 0 {
		return this.out.Top()
	}
	for this.in.Len() > 0 {
		this.out.Push(this.in.Pop())
	}
	return this.out.Top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.in.Len()+this.out.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
