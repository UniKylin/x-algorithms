// link		: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
// Author	: Kylin
// Date		: 2021-12-27

package leetcode

// 哈希加上双向链表
// 增加节点：
//      第一种没有超出容量的情况下，添加新节点是在头部，如果添加节点存在那么更新节点值同时把节点移动到头部
//      第二种超出容量的条件下，先添加节点到头部，删除尾部节点，减少链表大小
// 删除：一般是删除尾部节点和删除中间节点两者有区别
// 获取节点：当节点存在那么直接获取，访问了节点后移动到链表头部
// 如果没有特殊声明，左侧是链表头部，右侧是链表尾部

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*LinkedNode
	head, tail *LinkedNode
}

func initLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

func LRUConstructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*LinkedNode{},
		capacity: capacity,
		head:     initLinkedNode(0, 0),
		tail:     initLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.move2Head(node)
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		node := initLinkedNode(key, value)
		l.cache[key] = node
		l.add2Head(node)
		l.size++

		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.cache, removed.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.move2Head(node)
	}
}

func (l *LRUCache) removeTail() *LinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func (l *LRUCache) move2Head(node *LinkedNode) {
	l.removeNode(node)
	l.add2Head(node)
}

func (l *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) add2Head(node *LinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}
