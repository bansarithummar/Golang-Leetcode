981. Time Based Key-Value Store


type TimeMap struct {
    store map[string][]pair // Maps keys to a list of (value, timestamp) pairs
}

type pair struct {
    value     string
    timestamp int
}

func Constructor() TimeMap {
    return TimeMap{store: make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.store[key] = append(this.store[key], pair{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    pairs := this.store[key]
    i := sort.Search(len(pairs), func(i int) bool {
        return pairs[i].timestamp > timestamp
    })
    if i > 0 {
        return pairs[i-1].value
    }
    return "" 
}



