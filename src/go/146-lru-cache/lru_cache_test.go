package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lRUCache := LRUConstructor(2)
	lRUCache.Put(1, 1)                  // 缓存是 {1=1}
	lRUCache.Put(2, 2)                  // 缓存是 {1=1, 2=2}
	fmt.Println(">>>", lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)                  // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(">>>", lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)                  // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(">>>", lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(">>>", lRUCache.Get(3)) // 返回 3
	fmt.Println(">>>", lRUCache.Get(4)) // 返回 4
}
