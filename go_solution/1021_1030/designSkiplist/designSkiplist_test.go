package designSkiplist

import "testing"

func TestSkiplist(t *testing.T) {
	sl := Constructor()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	if sl.Search(0) {
		t.Fatalf("should not find")
	}
	sl.Add(4)
	if !sl.Search(1) {
		t.Fatalf("should find")
	}
	sl.Add(5)

	if !sl.Search(3) {
		t.Fatalf("should find")
	}
	if sl.Search(6) {
		t.Fatalf("should not find")
	}
}
