49.Group Anagrams

package main

import (
    "fmt"
    "sort"
    "strings"
)

func groupAnagrams(strs []string) [][]string {
    // Create a map to store the groups of anagrams
    groups := make(map[string][]string)

    // Iterate over each string in strs
    for _, str := range strs {
        // Convert the string to a sorted slice of characters
        chars := strings.Split(str, "")
        sort.Strings(chars)
        sortedStr := strings.Join(chars, "")

        // Add the string to the corresponding group in the map
        groups[sortedStr] = append(groups[sortedStr], str)
    }

    // Convert the map values to a slice of slices
    var result [][]string
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}

func main() {
    strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    result := groupAnagrams(strs)
    fmt.Println(result)
}

