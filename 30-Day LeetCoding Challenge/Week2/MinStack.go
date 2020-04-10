/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/

package week2

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{stack: []int{}, minStack: []int{}}
	return m
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	minLen := len(this.minStack)
	if minLen == 0 || x <= this.minStack[minLen-1] {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	stackLen := len(this.stack)
	minLen := len(this.minStack)
	if this.stack[stackLen-1] == this.minStack[minLen-1] {
		this.minStack = this.minStack[0 : minLen-1]
	}
	this.stack = this.stack[0 : stackLen-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
