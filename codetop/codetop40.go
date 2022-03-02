package codetop

//206 reverse linked list
//solution use recurision and itnerative

//146 LRU get is O(1) remove and insert is O(1)
//so we choose hashtable and double linked list

/* 3 idea: use two pointers
do interatively to record max length,use map to deal with repeat problem
when we find repeat:
e.g."12345" 3 comes, we should start new round from 3, so it should be 345 continue,and we need to clean the map of before 3 value also
but map has no sequence
so when i comes before 3, we do not start.like when 2 comes, because the index of 2 is 1,less than what we start point,we just give it up.
*/

//215 sort first O(nlogn)
//use max heap build heap takes O(n) +kO(nlogn)
//use quick select takes 1/2n + 1/4n + ... = 2n O(n) worse case is n*n n2

/* 25 reverse k group
	find the k location

	cut in two parts
	reverse the first part, the right part use reverse function

	join together and return
*/

/* 15 three sum
	sort first

*/

/* 912 quick sort
	partition
	left part quick sort
	right part quick sort
*/

/* 53 max subarray
	dp to solve
	find max subarray end with index value, prev +choose current index or not choose
*/

/* 21 combine two sorted linked list
	just use two pointer to conbine,append the left nodes
*/

/*
	1 two sums
	use map to record if exists,and find the location
*/

/*
	
*/