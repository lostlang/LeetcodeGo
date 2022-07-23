package leetcode

import "sort"

type NumberContainers struct {
	val map[int][]int
}

func Constructor() NumberContainers {
	return NumberContainers{make(map[int][]int)}
}

func (this *NumberContainers) Change(index int, number int) {
	this.Remove(index)
	if this.val[number] != nil {
		this.val[number] = append(this.val[number], index)
		sort.Ints(this.val[number])
	} else {
		this.val[number] = []int{index}
	}
}

func (this *NumberContainers) Find(number int) int {
	var out = -1
	if this.val[number] != nil && len(this.val[number]) > 0 {
		out = this.val[number][0]
	}

	return out
}

func (this *NumberContainers) Remove(number int) {
	for i := range this.val {
		for j := range this.val[i] {
			if this.val[i][j] == number {
				this.val[i] = append(this.val[i][:j], this.val[i][j+1:]...)
				return
			}
		}
	}
}
