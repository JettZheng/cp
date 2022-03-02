package advanced

type MaxHeap struct {
	data []int
}

// heapifydown used for traverse start value from top to bottom.
func (m *MaxHeap) heapifydown(index int) {
	var heapifyRoot = index

	for {
		var RoundBig = heapifyRoot

		l := 2 * heapifyRoot + 1
		r := 2 * heapifyRoot + 2
		if l < len(m.data) && m.data[l] > m.data[RoundBig] {
			RoundBig = l
		}
		if r < len(m.data) && m.data[r] > m.data[RoundBig] {
			RoundBig = r
		}
	
		if RoundBig == heapifyRoot { // got biggest, we can return
			return
		}

		swap(&m.data[heapifyRoot],&m.data[RoundBig])// up roundbig to top level
		heapifyRoot = RoundBig //start new round
	}
}

// heapifyup used for traverse end value from bottom to up.
func (m *MaxHeap) heapifyup(index int) {
	var ci = index
	for {
		pi := (ci -1)/2

		if m.data[pi] < m.data[ci] {
			swap(&m.data[pi],&m.data[ci])// up bigger value to top level
			ci = pi // now parent index is the child index of next round
		}else {
			break
		}
	}
}

func (m *MaxHeap) Build(){
	for i := (len(m.data) - 1)/2; i >=0; i-- {
		m.heapifydown(i)
	}
}

func (m *MaxHeap) Pop() int {
	m.data[0],m.data[len(m.data) - 1] = m.data[len(m.data) - 1],m.data[0]
	res := m.data[len(m.data) - 1]
    m.data = m.data[:len(m.data)-1]
	m.heapifydown(0)
	return res
}

func (m *MaxHeap) Push(value int) {
	m.data = append(m.data, value)
	m.heapifyup(len(m.data) - 1)
}

func swap(a,b *int) {
	*a,*b = *b,*a
}