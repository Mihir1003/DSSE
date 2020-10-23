package array

type array struct {
	array []int
}

func (arr *array) add(num int) {
	(*arr).array = append((*arr).array, num)
}

func (arr *array) delete(index int) {
	(*arr).array = append((*arr).array[:index], (*arr).array[index+1:]...)
}

func (arr *array) extractMin() int {
	min := (*arr).array[0]
	for _, x := range (*arr).array {
		if x < min {
			min = x
		}
	}
	return min
}

func (arr *array) extractMax() int {
	max := (*arr).array[0]
	for _, x := range (*arr).array {
		if x > max {
			max = x
		}
	}
	return max
}

func (arr *array) search(num int) int {

	for i, x := range (*arr).array {
		if x == num {
			return i
		}
	}
	return -1
}
