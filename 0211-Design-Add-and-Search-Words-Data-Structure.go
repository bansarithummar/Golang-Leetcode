type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type WordDictionary struct {
    root *TrieNode
}


func Constructor() WordDictionary {
    return WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}    
}


func (this *WordDictionary) AddWord(word string)  {
    node := this.root
    for _, ch := range word {
        if _, found := node.children[ch]; !found {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true    
}


func (this *WordDictionary) Search(word string) bool {
    return this.searchHelper(word, 0, this.root)    
}


func (this *WordDictionary) searchHelper(word string, index int, node *TrieNode) bool {
    if index == len(word) {
        return node.isEnd
    }
    
    ch := rune(word[index])
    if ch == '.' {
        for _, child := range node.children {
            if this.searchHelper(word, index+1, child) {
                return true
            }
        }
        return false
    } else {
        child, found := node.children[ch]
        if !found {
            return false
        }
        return this.searchHelper(word, index+1, child)
    }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
