package datastructure

import (
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val:   0,
					Left:  &TreeNode{
						Val:   1,
						Left:  	nil,
						Right:	nil,
					},
					Right: nil,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
