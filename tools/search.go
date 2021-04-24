package tools

// Binary search accepts an ordered array and return the index of the target. If target does not exist will return -1
func BinaryCount(arr []int, target int, lowerB, upperB int) int {
	if upperB == lowerB {
		if arr[lowerB] == target {
			return 1
		}
		return 0
	}

	lowerBound := lowerB
	upperBound := upperB
	for lowerBound <= upperBound {
		median := (lowerBound + upperBound) / 2

		if arr[median] == target {
			return BinaryCount(arr, target, lowerB, median-1) + 1 + BinaryCount(arr, target, median+1, upperB)
		} else if arr[median] > target {
			upperBound = median - 1
		} else {
			lowerBound = median + 1
		}
	}
	if lowerBound == upperBound && arr[lowerBound] == target {
		return lowerBound
	}
	return 0
}
