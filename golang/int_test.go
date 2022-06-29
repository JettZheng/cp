package golang

import (
	"testing"
)

func Test_testMaxInt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testMaxInt()
		})
	}
}

func TestChristmasTree(t *testing.T) {
	type args struct {
		height int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				height: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChristmasTree(tt.args.height)
		})
	}
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "a good   example",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testfunc(t *testing.T) {
	type args struct {
		in0 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testfunc(tt.args.in0); got != tt.want {
				t.Errorf("testfunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testcopy(t *testing.T) {
	type args struct {
		in0 int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testcopy(tt.args.in0)
		})
	}
}
