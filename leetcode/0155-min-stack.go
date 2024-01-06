package main

type MinStack struct {
	_stack []int
	_min   int
}

func Constructor() MinStack {
	return MinStack{
		_stack: make([]int, 0),
		_min:   2147483647,
	}

}

func (this *MinStack) Push(val int) {
	this._stack = append(this._stack, val)
	this._min = min(this._min, val)
}

func (this *MinStack) Pop() {
	if this._min == this._stack[len(this._stack)-1] {
		if len(this._stack) > 1 {
			this._min = this._stack[0]
			for i := 1; i < len(this._stack)-1; i++ {
				this._min = min(this._min, this._stack[i])
			}

		} else {
			this._min = 2147483647
		}
	}
	this._stack = this._stack[:len(this._stack)-1]
}

func (this *MinStack) Top() int {
	return this._stack[len(this._stack)-1]
}

func (this *MinStack) GetMin() int {
	return this._min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
