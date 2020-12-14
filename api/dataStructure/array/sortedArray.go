package array

type SortedArray struct {
	Array []int
}

func (arr *SortedArray) Add(num int) {
	for i, x := range (*arr).Array {
		if num > x {
			(*arr).Array = append((*arr).Array[:i], num)
			(*arr).Array = append((*arr).Array[:], (*arr).Array[i:]...)
		}
	}

}

func (arr *SortedArray) Delete(index int) {
	if index < len(arr.Array) {
		(*arr).Array = append((*arr).Array[:index], (*arr).Array[index+1:]...)
	}
}

func (arr *SortedArray) ExtractMin() int {
	if len(arr.Array) > 0 {
		return (*arr).Array[0]
	}
	return -1
}

func (arr *SortedArray) ExtractMax() int {
	if len(arr.Array) > 0 {
		return (*arr).Array[len(arr.Array)-1]
	}
	return -1
}

func (arr *SortedArray) Find(num int) int {
	if len(arr.Array) > 0 {
		return binarySearch(num, arr.Array)
	}
	return -1

}

func binarySearch(needle int, haystack []int) int {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else if haystack[median] == needle {
			return median
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return -1
	}

	return -1

}
