// Package sorting contains small, readable implementations for learning.
// For application code, prefer the Go standard library's slices package.
package sorting

// BubbleSort sorts values in ascending order by swapping neighboring values.
// It changes the given slice and stops early when a pass makes no swaps.
func BubbleSort(values []int) {
	for end := len(values) - 1; end > 0; end-- {
		swapped := false
		for i := 0; i < end; i++ {
			if values[i] > values[i+1] {
				values[i], values[i+1] = values[i+1], values[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// MergeSort sorts values in ascending order by sorting two halves and merging
// them. It changes the given slice and uses an extra slice of the same length.
func MergeSort(values []int) {
	if len(values) < 2 {
		return
	}
	mergeSort(values, make([]int, len(values)))
}

func mergeSort(values, buffer []int) {
	if len(values) < 2 {
		return
	}

	middle := len(values) / 2
	mergeSort(values[:middle], buffer[:middle])
	mergeSort(values[middle:], buffer[middle:])

	i, j, k := 0, middle, 0
	for i < middle && j < len(values) {
		if values[i] <= values[j] { // Taking the left value first keeps equal values in order.
			buffer[k] = values[i]
			i++
		} else {
			buffer[k] = values[j]
			j++
		}
		k++
	}
	for i < middle {
		buffer[k] = values[i]
		i++
		k++
	}
	for j < len(values) {
		buffer[k] = values[j]
		j++
		k++
	}
	copy(values, buffer[:len(values)])
}

// QuickSort sorts values in ascending order around a pivot. It changes the
// given slice and uses its last value as the pivot in each partition.
func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func quickSort(values []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(values, low, high)
	quickSort(values, low, pivotIndex-1)
	quickSort(values, pivotIndex+1, high)
}

func partition(values []int, low, high int) int {
	pivot := values[high]
	nextSmall := low
	for i := low; i < high; i++ {
		if values[i] < pivot {
			values[nextSmall], values[i] = values[i], values[nextSmall]
			nextSmall++
		}
	}
	values[nextSmall], values[high] = values[high], values[nextSmall]
	return nextSmall
}
