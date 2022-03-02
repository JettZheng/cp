package advanced

type MinHeap struct {
	data []int
}

func BuildMinHeap(arr []int) *MinHeap{
	m := &MinHeap{
		data:arr,
	}

	for i := (len(arr) -1)/2; i >= 0; i-- { //start from the end of array to make sure min value goes to front.
		m.heapify(i) //heapify starts from the last but one layer,
	}

	return m
}

func (m *MinHeap) Pop() int {
	m.data[0],m.data[len(m.data) - 1] = m.data[len(m.data) - 1],m.data[0]

	res := m.data[len(m.data) - 1]
	m.heapify(0)

	return res
}

func (m *MinHeap) Push(v int) {
	m.data = append(m.data, v)
	m.heapifyUp(len(m.data) - 1)
}

func (m *MinHeap) heapifyUp(i int) {
	if i == 0 { // 0 position means already be the min
		return
	}
	// 2i+1
	// 2i+2
	pi := (i - 1)/2

	if m.data[pi] > m.data[i] {// if child is small than parent,than swap value
		m.data[pi],m.data[i] = m.data[i],m.data[pi]
		m.heapifyUp(pi)
	}else {
		return
	}
}

func (m *MinHeap) heapify(i int) {
	l := 2 * i + 1
	r := 2 * i + 2

	var smallest = i // smallest index in parent and two children
	if l < len(m.data) && m.data[smallest] > m.data[l] {
		smallest = l
	}
	if r < len(m.data) && m.data[smallest] > m.data[r] {
		smallest = r
	}

	if smallest == i { // if i is still smallest,finished
		return
	}

	m.data[i],m.data[smallest] = m.data[smallest],m.data[i]// swap with the smallest

	m.heapify(smallest)
}

func minheap() {
	min := BuildMaxHeap([]int{1,2,3,4,5})
	a := min.MPop()
	print(a)
}

// max heap 

func BuildMaxHeap(arr []int) *MinHeap{
	m := &MinHeap{
		data:arr,
	}

	for i := (len(arr)-1)/2; i >= 0; i-- { //start from the end of array to make sure min value goes to front.
		m.Mheapify(i) //heapify starts from the last but one layer,
	}

	return m
}

func (m *MinHeap) MPop() int {
	m.data[0],m.data[len(m.data) - 1] = m.data[len(m.data) - 1],m.data[0]

	res := m.data[len(m.data) - 1]
	m.data = m.data[:len(m.data)-1]
	m.Mheapify(0)

	return res
}

func (m *MinHeap) Mheapify(i int) {
	l := 2 * i + 1
	r := 2 * i + 2

	var biggest = i // biggest index in parent and two children
	if l < len(m.data) && m.data[biggest] < m.data[l] {
		biggest = l
	}
	if r < len(m.data) && m.data[biggest] < m.data[r] {
		biggest = r
	}

	if biggest == i { // if i is still biggest,finished
		return
	}

	m.data[i],m.data[biggest] = m.data[biggest],m.data[i]// swap with the biggest

	m.Mheapify(biggest)
}