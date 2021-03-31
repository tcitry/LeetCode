package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 1
	r := n
	mid := 0
	for l < r {
		mid = (l + r) / 2
		if isBadVersion(mid) == true {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(n int) bool {
	return true
}
