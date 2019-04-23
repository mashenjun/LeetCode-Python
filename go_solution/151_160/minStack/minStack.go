package minStack

// 155. Min Stack

type MinStack struct {
	data []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:make([]int,0),
	}
}


func (this *MinStack) Push(x int)  {
	if(len(this.data) == 0){
		this.data = append(this.data, x)
		this.data = append(this.data, x)
	}else{
		tmp := this.data[len(this.data)-1]
		this.data = append(this.data, x)
		if(tmp<x){
			this.data = append(this.data, tmp)
		}else{
			this.data = append(this.data, x)
		}
	}
}


func (this *MinStack) Pop()  {
	this.data = this.data[:len(this.data)-2]
}


func (this *MinStack) Top() int {
	return this.data[len(this.data)-2]
}


func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
