package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

//* In: nums Array of ints, target int
//? Want: Indices of nums entries that equal target
/* Constraints
2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9  --> 2*10^9 is in 64 bit range
-10^9 <= target <= 10^9
Only one valid answer exists.

*/

// @lc code=start
func twoSum(nums []int, target int) []int {
	//* Naive approach would be to use two for loops
	//* But we can actually exploit maps
	// target = num_i + num_j <-> num_i = target - num_j
	seenNums := make(map[int]int, 0)
	for i, num_i := range nums {
		if j, isHit := seenNums[target-num_i]; isHit {
			return []int{j, i}
		}
		seenNums[num_i] = i
	}

	return []int{}

}

// @lc code=end
