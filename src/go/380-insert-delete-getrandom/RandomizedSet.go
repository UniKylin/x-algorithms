// link		: https://leetcode-cn.com/problems/insert-delete-getrandom-o1/
// Author	: Kylin
// Date		: 2021-10-13

package leetcode

import "math/rand"

/// Shit leetcode. Run many times. Always errors...... What's fuck...
type RandomizedSet struct {
	Map   map[int]int
	Array []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Map[val]; ok {
		return false
	}
	this.Array = append(this.Array, val)
	this.Map[val] = len(this.Array) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Map[val]; !ok {
		return false
	}

	/// update map and array
	removeIndex := this.Map[val]
	lastIndex := len(this.Array) - 1
	lastValue := this.Array[lastIndex]

	// pick last element to remove position
	this.Array[removeIndex], this.Array[lastIndex] = lastValue, val
	this.Array = this.Array[:lastIndex]

	this.Map[lastValue] = removeIndex
	delete(this.Map, val)

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.Array))
	return this.Array[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
