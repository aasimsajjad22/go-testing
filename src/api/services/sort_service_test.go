package services

import (
	"github.com/aasimsajjad22/go-testing/src/api/utils/sort"
	"testing"
)

func TestConstants(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	elements := []int{3,2,1,5,6,3,7,8}
	Sort(elements)
	if elements[0] != 1 {
		t.Error("first element should be 1")
	}
	if elements[len(elements) - 1] != 8 {
		t.Error("last element should be 8")
	}
}

func TestSortMoreThan10000(t *testing.T) {
	elements := sort.GetElements(20001)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 20000 {
		t.Error("last element should be 20000")
	}
}

func BenchmarkBubbleSort10K(b *testing.B) {
	elements := sort.GetElements(20000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100K(b *testing.B) {
	elements := sort.GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
