func numTilePossibilities(tiles string) int {
   // Build frequency map
   freq := make(map[rune]int)
   for _, t := range tiles {
       freq[t]++
   }
   
   // Backtrack to generate sequences
   var dfs func() int
   dfs = func() int {
       count := 0
       
       // Try each available character
       for c, f := range freq {
           if f > 0 {
               // Use this character
               freq[c]--
               count++
               
               // Recursively try remaining chars
               count += dfs()
               
               // Backtrack
               freq[c]++
           }
       }
       
       return count
   }
   
   return dfs()
}
