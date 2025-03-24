func countDays(days int, meetings [][]int) int {
    events := make(map[int]int)

    for _, meeting := range meetings {
        start, end := meeting[0], meeting[1]
        events[start]++
        events[end+1]--
    }

    keys := make([]int, 0, len(events))
    for k := range events {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    busy := 0
    unscheduledDays := 0
    prev := 1

    for _, day := range keys {
        if day > days {
            break
        }
        if busy == 0 {
            unscheduledDays += day - prev
        }
        busy += events[day]
        prev = day
    }

    if prev <= days && busy == 0 {
        unscheduledDays += days - prev + 1
    }
    return unscheduledDays    
}
