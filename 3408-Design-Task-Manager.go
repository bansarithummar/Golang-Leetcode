type taskItem struct {
	userId   int
	taskId   int
	priority int
}

type TaskManager struct {
	h   []taskItem       
	pos map[int]int      
}

func (tm *TaskManager) less(a, b int) bool {
	ia, ib := tm.h[a], tm.h[b]
	if ia.priority == ib.priority {
		return ia.taskId > ib.taskId 
	}
	return ia.priority > ib.priority
}

func (tm *TaskManager) swap(i, j int) {
	tm.pos[tm.h[i].taskId], tm.pos[tm.h[j].taskId] = j, i
	tm.h[i], tm.h[j] = tm.h[j], tm.h[i]
}

func (tm *TaskManager) heapUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if tm.less(i, p) {
			tm.swap(i, p)
			i = p
		} else {
			break
		}
	}
}

func (tm *TaskManager) heapDown(i int) {
	n := len(tm.h)
	for {
		l := 2*i + 1
		r := 2*i + 2
		best := i
		if l < n && tm.less(l, best) {
			best = l
		}
		if r < n && tm.less(r, best) {
			best = r
		}
		if best == i {
			break
		}
		tm.swap(i, best)
		i = best
	}
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		h:   make([]taskItem, 0, len(tasks)),
		pos: make(map[int]int),
	}
	for _, t := range tasks {
		userId, taskId, priority := t[0], t[1], t[2]
		tm.h = append(tm.h, taskItem{userId: userId, taskId: taskId, priority: priority})
		tm.pos[taskId] = len(tm.h) - 1
	}
	
	for i := len(tm.h)/2 - 1; i >= 0; i-- {
		tm.heapDown(i)
	}
	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	it := taskItem{userId: userId, taskId: taskId, priority: priority}
	this.h = append(this.h, it)
	idx := len(this.h) - 1
	this.pos[taskId] = idx
	this.heapUp(idx)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	idx, ok := this.pos[taskId]
	if !ok {
		return
	}
	this.h[idx].priority = newPriority
	this.heapUp(idx)
	this.heapDown(idx)
}

func (this *TaskManager) Rmv(taskId int) {
	idx, ok := this.pos[taskId]
	if !ok {
		return
	}
	last := len(this.h) - 1
	if idx != last {
		this.swap(idx, last)
	}

	this.h = this.h[:last]
	delete(this.pos, taskId)
	if idx < len(this.h) {		
        this.heapUp(idx)
		this.heapDown(idx)
	}
}

func (this *TaskManager) ExecTop() int {
	if len(this.h) == 0 {
		return -1
	}
	top := this.h[0]
	last := len(this.h) - 1
	if last == 0 {
		this.h = this.h[:0]
		delete(this.pos, top.taskId)
		return top.userId
	}

	this.swap(0, last)
	this.h = this.h[:last]
	delete(this.pos, top.taskId)
	this.heapDown(0)
	return top.userId
}

