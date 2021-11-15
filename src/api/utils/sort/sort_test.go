package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {
	elements := []int{3,2,1,5,6,3,7,8}
	assert.NotNil(t, elements)
	assert.EqualValues(t, 8, len(elements))
	assert.EqualValues(t, 8, elements[len(elements) - 1])
	assert.EqualValues(t, 3, elements[0])

	//BubbleSort(elements)
	timeoutChan := make(chan bool, 1)
	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		timeoutChan <- true
	}()
	if <- timeoutChan {
		assert.Fail(t, "bubble sort took more then 50 ms")
	}

	assert.NotNil(t, elements)
	assert.EqualValues(t, 8, len(elements))
	assert.EqualValues(t, 8, elements[len(elements) - 1])
	assert.EqualValues(t, 1, elements[0])

}

func TestSortIncreasingOrder(t *testing.T) {
	elements := []int{3,2,1,5,6,3,7,8}
	Sort(elements)
	if elements[0] != 1 {
		t.Error("first element should be 1")
	}
	if elements[len(elements) - 1] != 8 {
		t.Error("last element should be 8")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(1000)
	for i := 0; i < b.N; i ++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(1000)
	for i := 0; i < b.N; i ++ {
		Sort(elements)
	}
}