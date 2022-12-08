/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */


 type Queue struct {
	items []int
}



func Constructor() MyStack {
    return MyStack{}
}


func (this *MyStack) Push(x int)  {
    this.items = append(this.items, x)
}


func (this *MyStack) Pop() int {
	l := len(this.items)-1
	toRemove:= this.items[l]
	s.items = this.items[:l]
	return toRemove 
    
}


func (this *MyStack) Top() int {
    return this.items[0]
}


func (this *MyStack) Empty() bool {
        if len(this.items) == 0 {
			return true
		}
		return false
    
}


