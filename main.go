package main

import (
	"./bogosort"
	"./bubblesort"
	"./heapsort"
	"fmt"
)

func main() {
	slice := []int{2, 3, 1}
	fmt.Println("Bogo Sort: ", bogosort.Bogosort(slice))
	fmt.Println("Bubble Sort: ", bubblesort.BubbleSort(slice))
	fmt.Println("Heap Sort: ", heapsort.HeapSort(slice))
}
