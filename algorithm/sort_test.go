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
				nums: &[]int{5,6,1,2,3,4},
				l:    0,
				r:    5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, tt.args.l, tt.args.r)
		})
	}
}
