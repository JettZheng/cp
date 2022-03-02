package algorithm

import "testing"

func Test_quickSort(t *testing.T) {
	type args struct {
		nums *[]int
		l    int
		r    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				nums: &[]int{8,4,5,7,6,5,6},
				l:    0,
				r:    6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, tt.args.l, tt.args.r)
		})
	}
}
