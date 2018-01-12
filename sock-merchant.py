#!/bin/python3


def sock_merchant(sock_colors):
    sock_hash = {}
    sock_matches = 0

    for i, color in enumerate(sock_colors):
        sock_hash[color] = sock_hash[color] + 1 if color in sock_hash else 1
        if sock_hash[color] == 2:
            sock_matches = sock_matches + 1
            sock_hash[color] = 0

    return sock_matches


def main():
    n = int(input().strip())
    sock_colors = list(map(int, input().strip().split(' ')))
    result = sock_merchant(sock_colors)
    print(result)


def test():
    n = 9
    sock_colors = [10, 20, 20, 10, 10, 30, 50, 10, 20]
    actual_res = sock_merchant(sock_colors)
    expected_res = 3
    print('Good Code here' if actual_res == expected_res else 'Bad Code! Expected {}. Got {}'.format(expected_res,
                                                                                                     actual_res))


main()
# test()
