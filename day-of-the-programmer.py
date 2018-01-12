#!/bin/python3


def number_of_days_of_february(y):
    if y == 1918:
        return 15

    # Julian Calendar
    if y < 1918:
        return 29 if y % 4 == 0 else 28
    # Gregorian Calendar
    else:
        return 29 if y % 400 == 0 or y % 4 == 0 and y % 100 != 0 else 28


def find_day_of_programmer(y):
    if y < 1700 or y > 2700:
        raise ValueError('Year Must be between 1700 and 2700')

    total = 0
    days_in_month = [31, number_of_days_of_february(y), 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    for month in range(12):
        total += days_in_month[month]
        if total > 256:
            day = days_in_month[month] + 256 - total
            return "{}.{}.{}".format(str(day).zfill(2), str(month + 1).zfill(2), y)


####################################
#  BOILER PLATE FROM HACKER RANK  #
####################################
def solve(year):
    return find_day_of_programmer(year)


result = solve(int(input().strip()))
print(result)
