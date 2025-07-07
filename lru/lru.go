package lru

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

	c.head = &Node{}
	c.tail = &Node{}
	c.head.next = c.tail
	c.tail.prev = c.head

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

		if len(c.cache) >= c.capacity {
			c.removeTail()
		}

		c.cache[key] = newNode
		c.addToHead(newNode)
	}
}
func (c *Cache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *Cache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *Cache) addToHead(node *Node) {
	node.prev = c.head
	node.next = c.head.next

	c.head.next.prev = node
	c.head.next = node
}

func (c *Cache) removeTail() *Node {
	lastNode := c.tail.prev
	c.removeNode(lastNode)
	delete(c.cache, lastNode.key)
	return lastNode
}
