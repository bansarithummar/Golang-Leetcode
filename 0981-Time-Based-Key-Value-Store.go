type Entry struct {
	value     string
	timestamp int
}

type TimeMap struct {
    store map[string][]Entry    
}

func Constructor() TimeMap {
    return TimeMap{store: make(map[string][]Entry)}    
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.store[key] = append(this.store[key], Entry{value, timestamp})    
}

func (this *TimeMap) Get(key string, timestamp int) string {
    entries, exists := this.store[key]
	if !exists {
		return ""
	}
	idx := sort.Search(len(entries), func(i int) bool {
		return entries[i].timestamp > timestamp
	})
	if idx == 0 {
		return ""
	}
	return entries[idx-1].value    
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
