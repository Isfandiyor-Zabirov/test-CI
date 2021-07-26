package main

import "testing"

func TestIndexofMaxAndMaxElement(t *testing.T) {

	tests := []struct {
		name  string
		number []int
		indexwant  int
		MaxElementWant int
	}{
		{"All  of numbers are positive", []int{1, 4, 25, 25, 100}, 4, 100},
	//	{"All of numbers are negative", []int{-1, -40, -10, -1, -5}, 0, -1},
		{"Numbers are positive and negative", []int{-1, 5, -20, 2}, 1, 5},
	}

	for _, test := range tests {
		index, element := IndexofMaxAndMaxElement(test.number)

		if index != test.indexwant || element != test.MaxElementWant {
			t.Errorf("IndexofMaxAndMaxElement test %s gotIndex %d Indexwant %d gotElement %d Elementwant %d", test.name, index, test.indexwant, element, test.MaxElementWant )

		}
	}
}
