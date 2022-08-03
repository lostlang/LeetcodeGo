package leetcode

type MyCalendar struct {
	data [][2]int
}

func ConstructorCalendar() MyCalendar {
	return MyCalendar{[][2]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {

	for _, pair := range this.data {
		if pair[0] >= end || pair[1] <= start {
			continue
		}

		return false
	}

	this.data = append(this.data, [2]int{start, end})

	return true
}
