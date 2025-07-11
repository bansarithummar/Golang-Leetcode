type RoomHeap []int
func (h RoomHeap) Len() int           { return len(h) }
func (h RoomHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h RoomHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *RoomHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *RoomHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type BusyRoom struct {
	endTime int
	room    int
}
type BusyHeap []BusyRoom
func (h BusyHeap) Len() int           { return len(h) }
func (h BusyHeap) Less(i, j int) bool {
	if h[i].endTime == h[j].endTime {
		return h[i].room < h[j].room
	}
	return h[i].endTime < h[j].endTime
}
func (h BusyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *BusyHeap) Push(x interface{}) { *h = append(*h, x.(BusyRoom)) }
func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	count := make([]int, n)
	available := &RoomHeap{}
	for i := 0; i < n; i++ {
		heap.Push(available, i)
	}
	busy := &BusyHeap{}

	for _, m := range meetings {
		start, end := m[0], m[1]
		for busy.Len() > 0 && (*busy)[0].endTime <= start {
			room := heap.Pop(busy).(BusyRoom).room
			heap.Push(available, room)
		}

		duration := end - start
		if available.Len() > 0 {
			room := heap.Pop(available).(int)
			heap.Push(busy, BusyRoom{end, room})
			count[room]++
		} else {
			earliest := heap.Pop(busy).(BusyRoom)
			newEnd := earliest.endTime + duration
			heap.Push(busy, BusyRoom{newEnd, earliest.room})
			count[earliest.room]++
		}
	}

	maxCount := 0
	answer := 0
	for i, c := range count {
		if c > maxCount {
			maxCount = c
			answer = i
		}
	}
	return answer
}
