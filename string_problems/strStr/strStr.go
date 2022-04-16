package main

import "fmt"

/*
Problem: Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

Problem link: https://leetcode.com/problems/implement-strstr/

Approach:
- If len(needle)==0 then return 0.
- Scan from left.
- Inside the loop check if currIndex+len(neeedle)<= len(haystack)
- If above condition is true, then check if needle==haystack[currIndex:currIndex+len(neeedle)].
- If above condition is true, then return currIndex.
- If nothing matches then return -1.
*/

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen == 0 {
		return 0
	}
	for ind := range haystack {
		end := ind + needleLen
		if ind+needleLen <= haystackLen {
			if needle == haystack[ind:end] {
				return ind
			}
		}
	}
	return -1
}

/*
                                 .__        __     ________________  ____________
  _____ _____ ___  __ ___________|__| ____ |  | __/  _____/   __   \/_   \_____  \
 /     \\__  \\  \/ // __ \_  __ \  |/ ___\|  |/ /   __  \\____    / |   |/  ____/
|  Y Y  \/ __ \\   /\  ___/|  | \/  \  \___|    <\  |__\  \  /    /  |   /       \
|__|_|  (____  /\_/  \___  >__|  |__|\___  >__|_ \\_____  / /____/   |___\_______ \
      \/     \/          \/              \/     \/      \/                       \/

*/
