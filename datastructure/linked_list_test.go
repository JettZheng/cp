package datastructure

import (
	"reflect"
	"testing"
)

func Test_reverseListIteratively(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			},
			want: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListIteratively(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListIteratively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListRecursively(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			},
			want: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListRecursively(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4,&ListNode{5,nil}}}}},
				k:    2,
			},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3,&ListNode{5,nil}}}}},
		},
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, nil}},
				k:    2,
			},
			want:  &ListNode{2, &ListNode{1, nil}},
		},
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3,&ListNode{4,nil}}}},
				k:    2,
			},
			want:  &ListNode{2, &ListNode{1, &ListNode{4,&ListNode{3,nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
