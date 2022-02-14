package range_extr

import (
	"sort"
	"strconv"
)

type List struct {
	data    []int
	id      int
	cur     int
	prev    int
	isRange bool
	count   int
	res     string
}

func Solution(list []int) string {
	if len(list) == 0 {
		return ""
	}

	l := newList(list)
	for l.iter() {
		if l.run() {
			break
		}

	}
	return l.res
}

func newList(data []int) (list *List) {
	sort.Ints(data)
	list = &List{
		data: data,
	}
	list.init()
	return
}

func (l *List) init() {
	l.isRange = false
	l.prev = l.data[l.id]
	l.count = len(l.data)
	l.res += strconv.Itoa(l.prev)
}

func (l *List) iter() bool {
	l.prev = l.data[l.id]
	l.id += 1
	if l.id == l.count {
		return false
	}

	l.cur = l.data[l.id]
	return true
}

func (l *List) run() bool {
	if l.isApplyRes() || l.isLastForRange() {
		return false
	}

	if l.isLast() {
		return true
	}

	l.res += l.getSplitForRange()
	return false
}

func (l *List) getSplitForRange() string {
	l.isRange = true
	if l.data[l.id+1] == l.cur+1 {
		return "-"
	}
	return ","
}

func (l *List) isLast() bool {
	if l.count-1 == l.id {
		l.res += "," + strconv.Itoa(l.cur)
		return true
	}
	return false
}

func (l *List) isLastForRange() bool {
	if l.isRange && l.count == l.id+1 {
		l.res += strconv.Itoa(l.cur)
	}
	return l.isRange
}

func (l *List) isApplyRes() bool {
	if l.prev+1 == l.cur {
		return false
	}

	if l.isRange {
		l.res += strconv.Itoa(l.prev)
		l.isRange = false
	}
	l.res += "," + strconv.Itoa(l.cur)
	return true
}
