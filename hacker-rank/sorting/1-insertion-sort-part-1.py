#!/bin/python3


def print_arr(arr):
    print(" ".join(str(val) for val in arr))


def insertion_sort_1(n, arr):
    val = arr[n - 1]

    for i in range(len(arr) - 1, -1, -1):
        arr[i] = arr[i - 1]
        if arr[i - 1] < val or i == 0:
            arr[i] = val
            break
        print_arr(arr)
    print_arr(arr)

    return None


def main():
    n = int(input().strip())
    arr = list(map(int, input().strip().split(' ')))
    insertion_sort_1(n, arr)


main()
# insertion_sort_1(10, [2, 3, 4, 5, 6, 7, 8, 9, 10, 1])
# insertion_sort_1(5, [2, 4, 6, 8, 3])
