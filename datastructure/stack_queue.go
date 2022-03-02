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