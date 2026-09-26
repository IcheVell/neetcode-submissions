type MinStack struct {
	entry []int
	min []int
}

func Constructor() MinStack {
	return MinStack{entry : make([]int, 0), min: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.entry = append(this.entry, val)

	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else {
		currMin := this.min[len(this.min) - 1]
		
		if currMin < val {
			this.min = append(this.min, currMin)
		} else {
			this.min = append(this.min, val)
		}
	}
}

func (this *MinStack) Pop() {
	this.entry = this.entry[:len(this.entry) - 1]
	this.min = this.min[:len(this.min) - 1]
}

func (this *MinStack) Top() int {
	return this.entry[len(this.entry) - 1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min) - 1]
}
