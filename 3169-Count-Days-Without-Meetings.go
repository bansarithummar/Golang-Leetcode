func countDays(days int, meetings [][]int) int {
    events := make(map[int]int)

    // Record start and end events for meetings
    for _, meeting := range meetings {
        start, end := meeting[0], meeting[1]
        events[start]++
        events[end+1]--
    }

    // Process events in sorted order
    keys := make([]int, 0, len(events))
    for k := range events {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    // Calculate the number of free days
    busy := 0
    unscheduledDays := 0
    prev := 1

    for _, day := range keys {
        if day > days {
            break
        }
        // Count unscheduled days between intervals
        if busy == 0 {
            unscheduledDays += day - prev
        }
        // Update the busy counter
        busy += events[day]
        prev = day
    }

    // Account for any remaining days
    if prev <= days && busy == 0 {
        unscheduledDays += days - prev + 1
    }

    return unscheduledDays
    
}
