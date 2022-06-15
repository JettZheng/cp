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

/*
    "123"
    +"45"

    168
*/
func addStrings(num1 string, num2 string) string {
    i:= len(num1)-1
    j:= len(num2)-1
    
    var res []byte
    if i > j {
        res = make([]byte,i+2)   
    } else {
        res = make([]byte,j+2)
    }
   
    var carry int
    var k = len(res)-1
    for i>=0 && j>=0 {
        temp := int(num1[i] - '0') + int(num2[j] - '0') + carry
        
        if temp/10 >= 1 {
            carry = 1
        }else {
            carry = 0
        }
        curr := temp%10
        
        res[k] = byte('0' + curr)
        k--
        j--
        i--
    }
    
    return string(res[k+1:])
}