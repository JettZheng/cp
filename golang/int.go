package golang

import (
	"fmt"
	"math"
)

func testMaxInt() {
	const UnsignedMin = uint(0)
	const UnsignedMax = ^uint(0)
	var SignedMax = int(^uint(0) >> 1)
	var SignedMin = ^int(^uint(0) >> 1)

	var a = uint32(0)
	var b = int(^a)
	var c = 0 - b - 1

	var d = math.MaxUint32
	var e = 0 - d - 1

	var test = '0'
	print(int(rune(test)))

	fmt.Printf("int max is %d\n", SignedMax)
	fmt.Printf("int max is%d %d %d\n", a, b, c)
	fmt.Printf("int max is%d %d\n", d, e)
	print(SignedMin)
}

func ChristmasTree(height int) {
	if height <= 0 {
		return
	}

	for i := 0; i < height; i++ {
		var res []byte
		for j := 0; j < 2 * height - 1; j++ {
			l := height - 2 - i
			r := height  + i
			if j > l && j < r {
				res = append(res, '*')
			}

			res = append(res,' ')
		}
		fmt.Println(string(res))
	}
}

func reverseWords(s string) string {
    var input = []byte(s)

    for i := 0 ; i < len(input); {
        if input[i] == ' ' {
            input = input[1:]
        }else {
            break
        }
    }

    for i := len(input) - 1; i >= 0; i -- {
        if input[i] == ' ' {
            input = input[:len(input)-1]
        }else {
            break
        }
    }

    for i := 1 ; i < len(input); {
        if input[i] == ' ' && input[i-1]== ' ' {
            input = append(input[:i],input[i+1:]...)
        }else {
			i++
		}
    }


    revese(input,0,len(input)-1)

     var start int
    for i := 1; i < len(input); i ++ {
         if input[i] == ' ' {
             revese(input,start,i-1)
         }
         if input[i-1] == ' ' && input[i] != ' ' {
             start = i
         }
     }

    revese(input,start,len(input)-1)

    return string(input)
}


func revese(input []byte,l,r int) {
    for l < r {
        input[l],input[r] = input[r],input[l]
        l++
        r--
    }
}

// 3
// 2

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	total := len(nums1) + len(nums2)
	
	half := total/2
	
	if len(nums1) > len(nums2) {
		nums1,nums2 = nums2,nums1
	}
	
	l := 0
	r := len(nums1) - 1
	
	for {
		m := (l + r) >>1
		j := half - 2 - m
	
		var left1,left2,right1,right2 int
		if m >= 0 {
			left1 = nums1[m]
		} else {
			left1 = math.MinInt64
		}
	
		if m+1 < len(nums1) {
			right1 = nums1[m+1]
		} else {
			right1 = math.MaxInt64
		}
		
		if j >= 0 {
			left2 = nums2[j]
		} else {
			left2 = math.MinInt64
		}
	
		if j+1 < len(nums2) {
			right2 = nums2[j+1]
		} else {
			right2 = math.MaxInt64
		}
	
		if left1 <= right2 && left2 <= right1 {
		if total%2 == 0 {
			return float64( max(left1,left2) + min(right1,right2) )/ 2
		}
			return float64(min(right1,right2))
	
		} else if left1 > right2{
			r = m - 1
	
		} else  {
			l = m + 1
		}
	
	}
}

func min (a,b int) int {
    if a < b {
        return a
    }

    return b
}

func max(a,b int) int {
    if a < b {
        return b
    }
    return a
}

func testfunc(int ) int{
	print(-1>>1)
	print(-2>>1)
	print(-8>>1)
	print(-5>>1)
	print(-5/2)
	return 0 
}

func testcopy(int) {
	var a = []int {1,2,3,4,5}
	var b = []int{7,8,9,10}

	var a1 ='0'
	var b1 = '1'
	var c = b1 - a1
	var d = byte(2 +'0')
	print(c)
	fmt.Printf("value:%c",d)
	copy(a,b)
	print(a)
}
