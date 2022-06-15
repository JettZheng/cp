package datastructure

import (
	"reflect"
	"testing"
)

var (
	testNodeA  = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	testNodeA1 = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	testNodeB  = &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	testNodeC  = &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	testNodeD  = &ListNode{2, nil}
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
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    2,
			},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, nil}},
				k:    2,
			},
			want: &ListNode{2, &ListNode{1, nil}},
		},
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
				k:    2,
			},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
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

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				headA: testNodeA,
				headB: testNodeB,
			},
			want: testNodeC,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: testNodeA,
				n:    2,
			},
			want: testNodeA1,
		},
		{
			name: "",
			args: args{
				head: testNodeD,
				n:    1,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, nil}},
				n:    1,
			},
			want: &ListNode{1, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicatesIteratively(t *testing.T) {
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
				head: &ListNode{},
			},
			want: &ListNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesIteratively(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicatesIteratively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicatesRecursively(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesRecursively(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicatesRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoSortedListsIteratively(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				l1: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
				l2: &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoSortedListsIteratively(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoSortedListsIteratively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head:  testNodeA,
				left:  1,
				right: 5,
			},
			want: &ListNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reorderList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				head: testNodeA,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList2(tt.args.head)
		})
	}
}
