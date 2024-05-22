146. LRU Cache


type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node
    tail     *Node
}

type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        capacity: capacity,
        cache:    make(map[int]*Node),
        head:     &Node{},
        tail:     &Node{},
    }
    lru.head.next = lru.tail
    lru.tail.prev = lru.head
    return lru
}

func (this *LRUCache) moveToHead(node *Node) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.removeNode(node)
        this.moveToHead(node)
        return node.value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.cache[key]; ok {
        node.value = value
        this.removeNode(node)
        this.moveToHead(node)
    } else {
        newNode := &Node{
            key:   key,
            value: value,
        }
        this.cache[key] = newNode
        this.moveToHead(newNode)
        if len(this.cache) > this.capacity {
            delete(this.cache, this.tail.prev.key)
            this.removeNode(this.tail.prev)
        }
    }
}
