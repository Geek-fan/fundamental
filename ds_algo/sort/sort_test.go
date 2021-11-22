package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type array []int
func (a array) Len() int { return len(a) }
func (a array) Less(i, j int) bool { return a[i] < a[j] }
func (a array) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func newArray(n int) array {
	var a array
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		a = append(a, rand.Int()%100)
	}
	return a
}

func testSort(t *testing.T, s Sorter) {
	a := newArray(20)
	t.Logf("Before sort: %v", a)
	Sort(a, s)
	t.Logf("After sort: %v", a)
}

func TestBubbleSorter(t *testing.T) {
	testSort(t, BubbleSorter())
}

func TestInsertSorter(t *testing.T) {
	testSort(t, InsertSorter())
}

func TestQuickSorter(t *testing.T) {
	testSort(t, QuickSorter())
}

func benchSort(b *testing.B, s Sorter) {
	for i := 0; i < b.N; i++ {
		a := newArray(1000)
		Sort(a, s)
	}
}

func BenchmarkBubbleSorter(b *testing.B) {
	benchSort(b, BubbleSorter())
}

func BenchmarkInsertSorter(b *testing.B) {
	benchSort(b, InsertSorter())
}

func BenchmarkQuickSorter(b *testing.B) {
	benchSort(b, QuickSorter())
}

func BenchmarkGoLibSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := newArray(1000)
		sort.Sort(a)
	}
}