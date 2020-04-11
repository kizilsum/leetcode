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

// MaxUint constant.
const MaxUint = ^uint(0)

// MinUint constant.
const MinUint = 0

// MaxInt constant.
const MaxInt = int(MaxUint >> 1)

// MinInt constant.
const MinInt = -MaxInt - 1

// MinStack struct.
type MinStack struct {
	stack    []int
	minStack []int
}

// Constructor for MinStack struct.
func Constructor() MinStack {
	m := MinStack{stack: []int{}, minStack: []int{}}
	return m
}

// Push method for MinStack
func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	minLen := len(m.minStack)
	if minLen == 0 || x <= m.minStack[minLen-1] {
		m.minStack = append(m.minStack, x)
	}
}

// Pop method for MinStack
func (m *MinStack) Pop() {
	stackLen := len(m.stack)
	minLen := len(m.minStack)
	if m.stack[stackLen-1] == m.minStack[minLen-1] {
		m.minStack = m.minStack[0 : minLen-1]
	}
	m.stack = m.stack[0 : stackLen-1]
}

// Top method for MinStack
func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

// GetMin method for MinStack
func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
