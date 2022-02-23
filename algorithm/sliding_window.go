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

//Input: s1 = "abc", s2 = "bbbac"
//Output: true
//Explanation: s2 contains one permutation of s1 ("ba").

// solution: use fixed sliding window to iterative,
func checkInclusion(s1 string, s2 string) bool {
	a := len(s1)
	b := len(s2)
	if a > b {
		return false
	}

	var start = int('a') 

	var map1 =make([]int,26)
	var map2 =make([]int,26)
	for i := 0; i < a; i++ {
		map1[int(s1[i])-start]++
		map2[int(s2[i])-start]++
	}

	for i := 0; i < b - a; i ++ {
		if matched(map1,map2) {
			return true
		} else {
			print(int(s2[i+a]))
			map2[int(s2[i+a])-start]++
			map2[int(s2[i])-start]--
		}

	}

	return matched(map1,map2)
}

func matched(map1,map2 []int) bool {
	for i := 0; i < len(map1); i++ {
		if map1[i] != map2[i] {
			return false
		}
	}

	return true
}

func findAnagrams(p,s string) []int {
    var res []int
    
   	a := len(p)
	b := len(s)
	if a > b {
		return res
	}

	var start = int('a') 

	var map1 =make([]int,26)
	var map2 =make([]int,26)
	for i := 0; i < a; i++ {
		map1[int(p[i])-start]++
		map2[int(s[i])-start]++
	}

	for i := 0; i < b - a; i ++ {
		if matched(map1,map2) {
            res = append(res,i)
		}
		map2[int(s[i+a])-start]++
		map2[int(s[i])-start]--
	}
    
		if matched(map1,map2) {
            res = append(res,b-a)
		} 

	return res
}


