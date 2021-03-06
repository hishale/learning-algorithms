package problem0007

/*
[7] Reverse Integer

https://leetcode.com/problems/reverse-integer/description/

algorithms Easy (25.2%)
Total Accepted:    631.6K(631606)
Total Submissions: 2.5M(2506273)
Testcase Example:  ''

Given a 32-bit signed integer, reverse digits of an integer.

**Example 1:**

```
Input: 123
Output: 321
```

**Example 2:**

```
Input: -123
Output: -321
```

**Example 3:**

```
Input: 120
Output: 21
```

**Note:**

Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
func reverse(x int) int {
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}
	var Max = 1<<31 - 1
	if x == 0 {
		return 0
	}
	var y int
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > Max {
		y = 0
	}
	if isNeg {
		y = -y
	}
	return y
}
