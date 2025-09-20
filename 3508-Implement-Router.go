type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	limit int
	queue []Packet
	set   map[[3]int]bool
	index map[int][]int 
}


func Constructor(memoryLimit int) Router {
	return Router{
		limit: memoryLimit,
		queue: make([]Packet, 0),
		set:   make(map[[3]int]bool),
		index: make(map[int][]int),
	}
}


func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := [3]int{source, destination, timestamp}
	if this.set[key] {
		return false
	}

	if len(this.queue) == this.limit {
		old := this.queue[0]
		delete(this.set, [3]int{old.source, old.destination, old.timestamp})
		this.queue = this.queue[1:]

		tsList := this.index[old.destination]
		pos := sort.SearchInts(tsList, old.timestamp)
		if pos < len(tsList) && tsList[pos] == old.timestamp {
			this.index[old.destination] = append(tsList[:pos], tsList[pos+1:]...)
		}
	}

	p := Packet{source, destination, timestamp}
	this.queue = append(this.queue, p)
	this.set[key] = true
	this.index[destination] = append(this.index[destination], timestamp) 
	return true
}


func (this *Router) ForwardPacket() []int {
	if len(this.queue) == 0 {
		return []int{}
	}
	p := this.queue[0]
	this.queue = this.queue[1:]
	delete(this.set, [3]int{p.source, p.destination, p.timestamp})

	tsList := this.index[p.destination]
	pos := sort.SearchInts(tsList, p.timestamp)
	if pos < len(tsList) && tsList[pos] == p.timestamp {
		this.index[p.destination] = append(tsList[:pos], tsList[pos+1:]...)
	}
	return []int{p.source, p.destination, p.timestamp}
}


func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	tsList := this.index[destination]
	if len(tsList) == 0 {
		return 0
	}

	left := sort.SearchInts(tsList, startTime)
	right := sort.Search(len(tsList), func(i int) bool { return tsList[i] > endTime })
	return right - left
}

