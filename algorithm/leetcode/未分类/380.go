/*
	设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
		insert(val)：当元素 val 不存在时，向集合中插入该项。
		remove(val)：元素 val 存在时，从集合中移除该项。
		getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。

	对于 getRandom 方法，如果想「等概率」且「在 O(1) 的时间」取出元素，一定要满足：底层用数组实现，且数组必须是紧凑的。
	但如果用数组存储元素的话，插入，删除的时间复杂度怎么可能是 O(1) 呢？对数组尾部进行插入和删除操作不会涉及数据搬移，时间复杂度是 O(1)。
*/
type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		valToIndex: make(map[int]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	// 已存在
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	// 不存在，插入到尾部
	this.valToIndex[val] = len(this.nums)
	// 记录val对应的索引值
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	// 不存在
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	// 拿到val索引
	index := this.valToIndex[val]
	// 将最后一个元素对应的索引值修改为index(question  这一步不好想)
	this.valToIndex[this.nums[len(this.nums)-1]] = index
	// 交换val和最后一个元素
	this.nums[index], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[index]
	// 删除数组最后一个元素
	this.nums = this.nums[:len(this.nums)-1]
	// 删除val及其索引
	delete(this.valToIndex, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(1<<32)%len(this.nums)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */