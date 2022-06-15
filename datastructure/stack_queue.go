package datastructure

// 20 https://leetcode.com/problems/valid-parentheses
/*
	idea: create three stack for each parenthese
	if complete just remove both parenthese
	finally check is all stack is empty and clean

	got problem:
	we missed the situation like this ([)]

	how to deal it?
	we still need one stack
	just two ways:
	1.if we meet correct we empty the top of the stack
	2.if we meet wrong parenthese just keep in the stack
	3.finally we check if stack is empty
*/

func isValid(s string) bool {    
    var stack []byte
    stack = append(stack,s[0])
    
    var dict = make(map[byte]byte,3)
    dict[']']='['    
    dict['}']='{'
    dict[')']='('

    for i:=1;i<len(s);i++ {
        if v,ok := dict[s[i]];ok {
            if len(stack)-1 >= 0 && v ==stack[len(stack)-1] {
                stack = stack[:len(stack)-1]
                continue
            }
        }
        stack = append(stack,s[i])
    }
    
    return len(stack) == 0
}

type MyQueue struct {
    stack1 Stack
    stack2 Stack
}


func Constructor() MyQueue {
    return MyQueue{
        stack1:*new(Stack),
        stack2:*new(Stack),
    }
}

func (this *MyQueue) Push(x int)  {
    this.stack1.Push(x)
}


func (this *MyQueue) Pop() int {
    if len(this.stack2) > 0 {
         return this.stack2.Pop()
    }

    for len(this.stack1) > 0 {
        this.stack2.Push(this.stack1.Pop())
    }

    return this.stack2.Pop()
}

func (this *MyQueue) Peek() int {
    if this.stack2.Len() > 0 {
        return this.stack2.Peek()
    }

    for len(this.stack1) > 0 {
        this.stack2.Push(this.stack1.Pop())
    }

    return this.stack2.Peek()
}

func (this *MyQueue) Empty() bool {
    return this.stack2.Len() == 0 && this.stack1.Len() == 0
}

type Stack []int

func(s *Stack) Push(x int) {  
    *s = append(*s,x)
}

func(s *Stack) Pop() int {
    var res int

    if s.Len() > 0 {
        res = (*s)[s.Len()-1]
        *s = (*s)[:s.Len()-1]
    }

    return res
}

func (s *Stack) Peek() int {
    return (*s)[s.Len()-1]
}

func (s *Stack) Len() int {
    return len(*s)
}

func testqu() {
    var a = Constructor()

    a.Push(1)
    a.Push(2)
    a.Peek()
    a.Pop()
    a.Empty()
}