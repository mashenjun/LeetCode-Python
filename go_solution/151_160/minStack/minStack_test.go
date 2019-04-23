package minStack

import "testing"

func TestConstructor(t *testing.T) {
	var data []int
	minStack := Constructor()
	t.Log(minStack.data == nil)
	t.Log(data == nil)
}
