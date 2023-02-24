package algorithm

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{0, 1},
				target: 20,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchI(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchRecursively(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 6},
				target: 5,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearchRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchLower(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 2, 2, 6},
				target: 6,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchLower(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearchLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchUpper(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 2, 2, 6},
				target: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchUpper(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearchUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hsearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr:    []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14},
				target: 200,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hsearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("hsearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchII(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{0, 1},
				target: 9,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchII(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearchII() = %v, want %v", got, tt.want)
			}
		})
	}
}
