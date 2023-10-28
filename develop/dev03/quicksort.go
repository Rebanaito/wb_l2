package main

func quicksort(lines []string, low, high int, flags options) {
	if low >= high {
		return
	}
	pivot := quicksortPartition(lines, low, high, flags)
	quicksort(lines, low, pivot-1, flags)
	quicksort(lines, pivot+1, high, flags)
}

func quicksortPartition(lines []string, low, high int, flags options) int {
	pivot := lines[high]
	i := low - 1
	for j := low; j < high; j++ {
		if compare(lines[j], pivot, flags) {
			i++
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
	lines[i+1], lines[high] = lines[high], lines[i+1]
	return i + 1
}
