#!/bin/python3


def bon_appetit(n, k, b, items):
    actual_price = 0
    for i in range(n):
        actual_price = actual_price + items[i] / 2 if i != k else actual_price
    return 'Bon Appetit' if actual_price == b else int(b - actual_price)


def main():
    n, k = input().strip().split(' ')
    n, k = [int(n), int(k)]
    items = list(map(int, input().strip().split(' ')))
    b = int(input().strip())
    result = bon_appetit(n, k, b, items)
    print(result)


def test():
    n = 4
    k = 1
    b = 12
    items = [3, 10, 2, 9]
    actual_res = bon_appetit(n, k, b, items)
    expected_res = 5
    print('Good Code here' if actual_res == expected_res else 'Bad Code! Expected {}. Got {}'.format(expected_res,
                                                                                                     actual_res))


main()
# test()
