func canCompleteCircuit(gas []int, cost []int) int {
    total_gas := 0
    current_gas := 0
    start_index := 0
    for i := 0; i < len(gas); i++ {
        total_gas += gas[i] - cost[i]
        current_gas += gas[i] - cost[i]
        if current_gas < 0 {
            start_index = i + 1
            current_gas = 0
        }
    }
    if total_gas < 0 {
        return -1
    }
    return start_index
}
