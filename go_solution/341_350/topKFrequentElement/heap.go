package topKFrequentElement

// implement container heap interface
type Item struct {
	p int // priority
	v int // value
}

type Heap struct {
	// capability int
	data []Item
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Less(i, j int) bool {
	return h.data[i].p > h.data[j].p
}

func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) Pop() (v interface{}) {
	v, h.data = h.data[len(h.data)-1], h.data[:len(h.data)-1]
	return v
}

func (h *Heap) Push(v interface{}) {
	h.data = append(h.data, v.(Item))
}
