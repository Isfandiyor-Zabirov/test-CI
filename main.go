package main

func IndexofMaxAndMaxElement(numbers []int) (int, int)  {
	Index := 0
	MaxElement := 0

	for index, value := range numbers {
		if MaxElement < value {
			MaxElement = value
			Index = index
		}
	}
	return Index, MaxElement
}

func IndexofMaxElement(numbers []int) int  {
	Index := 0
	MaxElement := numbers[0]

	for index, value := range numbers {
		if MaxElement < value {
			MaxElement = value
			Index = index
		}
	}
	return Index
}

func MaxElement(numbers []int) int  {

	MaxElement := numbers[0]

	for _, value := range numbers {
		if MaxElement < value {
			MaxElement = value

		}
	}
	return MaxElement
}

func main() {

	IndexofMaxElement([]int{10, 25, -5, 1, -100, 25})
	MaxElement([]int{10, 25, -5, 1, -100, 25})


}
