# Bounded Square Sum

You are given two arrays of integers a and b, and two integers lower and upper.

Your task is to find the number of pairs `(i, j)` such that `lower ≤ a[i] * a[i] + b[j] * b[j] ≤ upper`.

### **Example 1:**

**For:** 
- a = [3, -1, 9]
- b = [100, 5, -2]
- lower = 7 
- upper = 99

**Output:**

`boundedSquareSum(a, b, lower, upper) = 4`.

**Explanation:**

- If i = 0 and j = 1, then a[0] = 3, b[1] = 5, and 7 ≤ 3 * 3 + 5 * 5 = 9 + 25 = 36 ≤ 99.
- If i = 0 and j = 2, then a[0] = 3, b[2] = -2, and 7 ≤ 3 * 3 + (-2) * (-2) = 9 + 4 = 13 ≤ 99.
- If i = 1 and j = 1, then a[1] = -1, b[1] = 5, and 7 ≤ (-1) * (-1) + 5 * 5 = 1 + 25 = 26 ≤ 99.
- If i = 2 and j = 2, then a[2] = 9, b[2] = -2, and 7 ≤ 9 * 9 + (-2) * (-2) = 81 + 4 = 85 ≤ 99.

There are only four pairs that satisfy the requirement.


### **Example 2:**

**For:**
- a = [1, 2, 3, -1, -2, -3]
- b = [10]
- lower = 0
- upper = 100

**Output:**

`boundedSquareSum(a, b, lower, upper) = 0`.

**Explanation:**

Since the array `b` contains only one element `10` and the array `a` does not contain `0`, 
it is not possible to satisfy `0 ≤ a[i] * a[i] + 10 * 10 ≤ 100`.