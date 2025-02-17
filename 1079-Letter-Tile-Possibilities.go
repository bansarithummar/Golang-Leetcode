func numTilePossibilities(tiles string) int {
   freq := make(map[rune]int)
   for _, t := range tiles {
       freq[t]++
   }

   var dfs func() int
   dfs = func() int {
       count := 0
       
       for c, f := range freq {
           if f > 0 {
               freq[c]--
               count++
               
               count += dfs()
               freq[c]++
           }
       }     
       return count
   }  
   return dfs()
}
