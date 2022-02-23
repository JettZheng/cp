package datastructure

func firstUniqChar(s string) int {
    var res = -1

	var queue = []byte{}
	var dict = make(map[byte]int)
	var dictIndex = make(map[byte]int)

	for i := range s {
		if _,ok := dict[s[i]];ok {
			dict[s[i]]++
			continue
		}else {
			dict[s[i]] = 1
			dictIndex[s[i]] = i
			queue = append(queue,s[i])
		}
	}

	for i := range queue {
		if dict[queue[i]] == 1 {
			res = dictIndex[queue[i]]
			break
		}
	}

	return res
}