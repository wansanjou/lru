package lru

import "fmt"

// Node represents a single entry in the cache. It is not exported.
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// Cache holds the state. It is the main exported type.
type Cache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

// New creates a new LRUCache instance with a given capacity.
func New(capacity int) *Cache {
	c := &Cache{
		capacity: capacity,
		cache:    make(map[int]*Node),
	}

	return c
}

func (c *Cache) Get(key int) (int, bool) {
	if node, found := c.cache[key]; found {
		c.moveToHead(node)
		return node.value, true
	}
	return 0, false
}

func (c *Cache) Put(key int, value int) {
	if node, found := c.cache[key]; found {
		node.value = value
		c.moveToHead(node)
	} else {
		newNode := &Node{
			key:   key,
			value: value,
		}

		c.cache[key] = newNode
		c.addToHead(newNode)

		if len(c.cache) > c.capacity {
			c.removeTail()
		}
	}
}

func (c *Cache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *Cache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		c.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		c.tail = node.prev
	}

}

func (c *Cache) addToHead(node *Node) {
	node.prev = nil
	node.next = c.head
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node

	if c.tail == nil {
		c.tail = node
	}
}

func (c *Cache) removeTail() {
	if c.tail == nil {
		return
	}

	delete(c.cache, c.tail.key)

	if c.tail.prev != nil {
		c.tail = c.tail.prev
		c.tail.next = nil
	} else {
		c.head = nil
		c.tail = nil
	}

}

func (c *Cache) PrintList() {
	if c.head == nil {
		fmt.Println("Empty cache")
		return
	}

	fmt.Print("Head -> ")
	current := c.head
	for current != nil {
		fmt.Printf("[%d:%d]", current.key, current.value)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println(" <- Tail")
}
