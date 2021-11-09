package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Sorter interface {
	sort(Interface)
}

type bubbleSorter struct {}
func BubbleSorter() Sorter { return bubbleSorter{} }

func (s bubbleSorter) sort(data Interface) {
	length := data.Len()
	for i := 0; i < length; i++ {
		for j := 0; j < length - 1 - i; j++ {
			if data.Less(j + 1, j) {
				data.Swap(j, j + 1)
			}
		}
	}
}

type insertSorter struct {}
func InsertSorter() Sorter { return insertSorter{} }

func (s insertSorter) sort(data Interface) {
	for i := 0; i < data.Len(); i++ {
		idx := i
		for j := i; j >= 0; j-- {
			if data.Less(idx, j) {
				data.Swap(idx, j)
				idx = j
			} else {
				continue
			}
		}
	}
}

type quickSorter struct{}
func QuickSorter() Sorter { return quickSorter{} }

func quickSort(data Interface, left, right int) {
	if left >= right {
		return
	}
	pivot := left
	l := left
	r := right
	for l < r {
		for !data.Less(r, pivot) && pivot < r {
			r--
		}
		data.Swap(pivot, r)
		pivot = r
		for !data.Less(pivot, l) && l < pivot {
			l++
		}
		data.Swap(l, pivot)
		pivot = l
	}
	quickSort(data, left, pivot - 1)
	quickSort(data, pivot + 1, right)
}

func (s quickSorter) sort(I Interface) {
	quickSort(I, 0, I.Len() - 1)
}

func Sort(data Interface, S Sorter) {
	S.sort(data)
}
