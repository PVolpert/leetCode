package main

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	base := 10
	// Guard 1 negative numbers
	// Guard 2 Last digit is 0 but x is not 0
	if x < 0 || x%base == 0 && x != 0 {
		return false
	}
	var reverseNumber int
	for x > reverseNumber {
		//To  extract rightmost digit of x we can use moduloand int division with base 10
		// eg 121 mod 10 = 1; 121 / 10 = 12;
		digit := x % base
		x = x / base
		//To align the extracted digits we can use former * base + current digit
		// eg former is 1 and current is 2 --> 1*10+2 = 12
		reverseNumber = reverseNumber*base + digit
	}
	// If x we need to exclude middle number
	return x == reverseNumber || x == reverseNumber/10
}

// @lc code=end
