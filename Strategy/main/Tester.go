package main

import (
	"fmt"
	"math/rand"
)
type DataSet struct {
	data[] int
	sortedList[] int
	strategy func(a []int)
}
func (this *DataSet)SetData()  {

	this.data = append(this.data, 198,56,12,34,62,99,145,87,26,37,13,16)
	this.sortedList = append(this.sortedList, 198,56,12,34,62,99,145,87,26,37,13,16)
}

func (this *DataSet)SetStrategy(strategy func([]int)){
	this.strategy = strategy
}

func (this *DataSet) DoSort()  {
	this.strategy(this.sortedList)
}

func (this *DataSet) Display(){
	fmt.Println(this.sortedList)
}

func (this *DataSet) ResetData(){
	copy(this.sortedList,this.data)
	fmt.Println("Data Reset Done!")
}
func (this *DataSet) ChangeStatregy(s func([]int))  {
	this.strategy = s
}

func split(A []int) ([]int, []int) {
	return A[0:len(A) / 2], A[len(A) / 2:]
}


func merge(A, B []int) []int {
	arr := make([]int, len(A) + len(B))

	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		}
		if A[j] > B[k] {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}

	return arr
}

func  MergeSort(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

func QuickSort(a []int) []int{
	if len(a) < 2 { return a }

	left, right := 0, len(a) - 1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	QuickSort(a[:left])
	QuickSort(a[left + 1:])


	return a
}



func main()  {
	bubbleSort := func(arr []int){
		for i:=1; i< len(arr); i++ {
			for j:=0; j < len(arr)-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		fmt.Println("List sorted with  BubbleSort.",)
		fmt.Println(arr)
	}
	mergeSort := func(A []int){
		if len(A) <= 1 {
			fmt.Println("List sorted with  BubbleSort.",)
			fmt.Println(A)
		}

		left, right := split(A)
		left = MergeSort(left)
		right = MergeSort(right)
		var arr []int
		arr = merge(left, right)
		fmt.Println("List sorted with  MergeSort.",)
		fmt.Println(arr)
	}
	quickSort := func(a []int) {
		var arr = QuickSort(a)
		fmt.Println("List sorted with  QuickSort.",)
		fmt.Println(arr)
	}
	var dataset = new(DataSet)
	dataset.SetData()
	dataset.SetStrategy(bubbleSort)
	dataset.Display()
	dataset.DoSort()
	dataset.ChangeStatregy(mergeSort)
	dataset.ResetData()
	dataset.Display()
	dataset.DoSort()
	dataset.ChangeStatregy(quickSort)
	dataset.ResetData()
	dataset.Display()
	dataset.DoSort()
}
