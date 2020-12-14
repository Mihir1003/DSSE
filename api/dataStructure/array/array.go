package array

type Array struct {
	Array []int
}

func (arr *Array) Add(num int) {
	(*arr).Array = append((*arr).Array, num)
}

func (arr *Array) Delete(index int) {
	if index < len(arr.Array) {
		(*arr).Array = append((*arr).Array[:index], (*arr).Array[index+1:]...)
	}

}

func (arr *Array) ExtractMin() int {
	if len(arr.Array) > 0 {
		min := (*arr).Array[0]
		for _, x := range (*arr).Array {
			if x < min {
				min = x
			}
		}
		return min
	}
	return -1
}

func (arr *Array) ExtractMax() int {
	if len(arr.Array) > 0 {
		max := (*arr).Array[0]
		for _, x := range (*arr).Array {
			if x > max {
				max = x
			}
		}
		return max
	}
	return -1

}

func (arr *Array) Find(num int) int {

	for i, x := range (*arr).Array {
		if x == num {
			return i
		}
	}
	return -1
}
