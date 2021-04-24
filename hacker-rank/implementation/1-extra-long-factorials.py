#!/bin/python3


def extra_long_factorial(n):
    if n < 0 or type(n).__name__ != 'int':
        raise ValueError('Not a valid input')
    if n <= 1:
        return n

    return n * extra_long_factorial( n - 1)


def main():
    n = int(input().strip())
    res = extra_long_factorial(n)
    print(res)


def test(n, expected_res):
    actual_res = extra_long_factorial(n)
    print('Good Code here' if actual_res == expected_res else 'Bad Code! Expected {}. Got {}'.format(expected_res,
                                                                                                     actual_res))


main()
# test(4, 24)
# test(25, 15511210043330985984000000)
