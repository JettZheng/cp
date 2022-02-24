package datastructure

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		input map[int]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				input: map[int]int{1:10000,10:1000,100:99,1000:10,10000:1},
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
