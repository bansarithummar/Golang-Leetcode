981. Time Based Key-Value Store

package main

import (
	"fmt"
	"sort"
)

type TimeMap struct {
	store map[string][][2]string
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][][2]string),
	}
}

func (tm *TimeMap) Set(key, value string, timestamp int) {
	if _, ok := tm.store[key]; !ok {
		tm.store[key] = [][2]string{}
	}
	tm.store[key] = append(tm.store[key], [2]string{value, fmt.Sprintf("%d", timestamp)})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	values, ok := tm.store[key]
	if !ok {
		return ""
	}

	// Perform binary search
	left, right := 0, len(values)-1
	for left <= right {
		mid := left + (right-left)/2
		if ts, _ := strconv.Atoi(values[mid][1]); ts <= timestamp {
			return values[mid][0]
		} else {
			right = mid - 1
		}
	}

	return ""
}


