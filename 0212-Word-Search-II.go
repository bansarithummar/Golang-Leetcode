type TrieNode struct {
    children map[byte]*TrieNode
    word     string
}

type Trie struct {
    root *TrieNode
}

func Constructor() Trie {
    return Trie{root: &TrieNode{children: make(map[byte]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
    node := this.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if _, found := node.children[ch]; !found {
            node.children[ch] = &TrieNode{children: make(map[byte]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.word = word
}

func findWords(board [][]byte, words []string) []string {
    trie := Constructor()
    for _, word := range words {
        trie.Insert(word)
    }
    result := []string{}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(board, i, j, trie.root, &result)
        }
    }
    return result
}

func dfs(board [][]byte, i, j int, node *TrieNode, result *[]string) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
        return
    } 
    ch := board[i][j]
    if ch == '#' || node.children[ch] == nil {
        return
    }
    node = node.children[ch]
    if node.word != "" {
        *result = append(*result, node.word)
        node.word = "" 
    }
    board[i][j] = '#' 
    dfs(board, i-1, j, node, result) 
    dfs(board, i+1, j, node, result) 
    dfs(board, i, j-1, node, result)
    dfs(board, i, j+1, node, result) 
    board[i][j] = ch 
}
