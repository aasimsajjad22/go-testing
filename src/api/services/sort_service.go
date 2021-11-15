package services

import "github.com/aasimsajjad22/go-testing/src/api/utils/sort"

const (
	privateConst = "private"
)
func Sort(elements []int) {
	if len(elements) < 10000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}