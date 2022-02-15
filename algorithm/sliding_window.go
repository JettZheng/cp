package algorithm
// 、"abcabcbb 3 // "tmmzuxt"
func lengthOfLongestSubstring(s string) int {
    var i,j,res int

	b := []byte(s)

	var charmap = make(map[byte]int)

	for ; i < len(b) && j < len(b); j++ {
		if x,ok := charmap[b[j]];ok && x > i{
			i = x + 1
		} else {
			res = max(res,j-i +1)
		}

		charmap[b[j]] = j
	}
	
	return res
}