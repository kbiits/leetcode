type MyStack struct {
    arr []int
}


func Constructor() MyStack {
    return MyStack{
        arr: make([]int, 0, 5),
    }
}


func (this *MyStack) Push(x int)  {
    this.arr = append(this.arr, x)
}


func (this *MyStack) Pop() int {
    returnedValue := this.arr[len(this.arr) - 1]
    this.arr = this.arr[:len(this.arr) - 1]
    return returnedValue
}


func (this *MyStack) Top() int {
    return this.arr[len(this.arr) - 1]
}


func (this *MyStack) Empty() bool {
    return len(this.arr) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
