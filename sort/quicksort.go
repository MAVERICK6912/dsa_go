package sort

import (
	"github.com/maverick6912/dsa_go/errors"
)

// QuickSort sorts the array as per the sortDirection passed.
// returns NoElements error if passed iterable is empty.
func QuickSort[T any](it []T, sortDir SortDirection, cmp func(T, T) int) ([]T, error) {
	if len(it) == 0 {
		return nil, errors.NoElements
	}
	low, high := 0, len(it)-1
	return partitionAndQuickSort(it, low, high, sortDir, cmp), nil
}

/*
	Dry run:
		- input: [4,3,2,1], sortDir: ascending
		- partitionAscending([4,3,2,1],0,3,cmp)
		- pivot,pRight:=it[3]->1,-1
		- for i:=0;i<3;i++:
			- if it[0]<pivot(1): (4<1)-> false
			- if it[1]<pivot(1): (3<1)-> false
			- if it[2]<pivot(1): (2<1)-> false
		- t:=it[-1+1]-> t:= it[0]-> t:=4
		- it[0]:=it[3]->it[0] = 1
		- it[high]=t-> it[3]= 4
		- return pRight+1 -> -1+1 -> 0
		- `it` at the end of this function call will be: [1,3,2,4]
		- partitionAndQuickSort([1,3,2,4],0,-1,ascending,cmp)-> this call will not do anything as low<high condition will be false.
		- partitionAndQuickSort([1,3,2,4],1,3,ascending,cmp):
			- if low<high: (1<3)-> true:
				- partitionAscending([4,3,2,1],1,3,cmp):
					- pivot, pRight :=(it[3])->4,(low-1)->0
					- for i:=1;i<3;i++:
						- if it[1]<pivot(4): (3<4)-> true:
							- swap(it[1],it[1])->[1,3,2,4]
							- pRight++=>1
						- if it[2]<pivot(4): (2<4)-> true:
							- swap(it[2],it[2])->[1,3,2,4]
							- pRight++=>2
					- t:=it[2+1]-> t:= it[3]-> t:=4
					- it[3]:=it[3]->it[3] = 4
					- it[high]=t-> it[3]= 4
					- return pRight+1 -> 2+1 -> 3
					- `it` at the end of this function call will be: [1,3,2,4]
		- partitionAndQuickSort([1,3,2,4],4,3,ascending,cmp)-> this call will not do anything as low<high condition will be false.
		- partitionAndQuickSort([1,3,2,4],1,2,ascending,cmp):
		- if low<high: (1<2)->true:
			- partitionAcending([1,3,2,4],1,2):
				- pivot,pRight:= (it[2])->2,0
				- for i:=1;1<2;i++:
					- if(3<2):false
				- swap(it[1],it[2])
				- Array after swap: [1,2,3,4]
				- return pRight+1-> 0+1-> 1
				- `it` at the end of this function call will be: [1,2,3,4]
		- partitionAndQuickSort([1,3,2,4],1,0,ascending,cmp)-> this call will not do anything as low<high condition will be false.
		- partitionAndQuickSort([1,3,2,4],2,2,ascending,cmp)-> this call will not do anything as low<high condition will be false.
		- Hence we will return: [1,2,3,4]
*/
// partitionAnQuickSort partitions and sorts the slice recursively.
func partitionAndQuickSort[T any](it []T, low, high int, sortDir SortDirection, cmp func(T, T) int) []T {
	if low < high {
		var partitionIndex int
		if sortDir == Ascending {
			partitionIndex = partitionAscending(it, low, high, cmp)
		} else {
			partitionIndex = partitionDescending(it, low, high, cmp)
		}

		partitionAndQuickSort(it, low, partitionIndex-1, sortDir, cmp)
		partitionAndQuickSort(it, partitionIndex+1, high, sortDir, cmp)
	}
	return it
}

// partitionAscending partitions and sorts the slice in ascending order
func partitionAscending[T any](it []T, low, high int, cmp func(T, T) int) int {
	pivot, pRight := it[high], low-1
	for i := low; i < high; i++ {
		if cmp(it[i], pivot) < 0 {
			pRight += 1
			t := it[pRight]
			it[pRight] = it[i]
			it[i] = t
		}
	}
	t := it[pRight+1]
	it[pRight+1] = it[high]
	it[high] = t

	return pRight + 1
}

// partitionDescending partitions and sorts the slice in descending order
func partitionDescending[T any](it []T, low, high int, cmp func(T, T) int) int {
	pivot, pRight := it[high], low-1
	for i := low; i < high; i++ {
		if cmp(it[i], pivot) > 0 {
			pRight += 1
			t := it[pRight]
			it[pRight] = it[i]
			it[i] = t
		}
	}
	t := it[pRight+1]
	it[pRight+1] = it[high]
	it[high] = t

	return pRight + 1
}
