/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
package main


type MyQueue struct {
    items []int
    
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.items = append(this.items, x)
}


func (this *MyQueue) Pop() int {
    l := len(this.items)-1
	toRemove := this.items[l]
	this.items = this.items[1:]
	return toRemove
}


func (this *MyQueue) Peek() int {
    return this.items[0]
}


func (this *MyQueue) Empty() bool {
    if this.items == nil {
		return true
	}
	return false
}


