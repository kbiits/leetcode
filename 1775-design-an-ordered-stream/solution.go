type OrderedStream struct {
    size int
    arr []string
    ptr int
}


func Constructor(n int) OrderedStream {
    return OrderedStream{
        size: n,
        arr: make([]string, n),
        ptr: 0,
    }
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.arr[idKey - 1] = value
    
    // move the ptr
    beforeMovedPtr := this.ptr

    // while ptr is in index bound and 
    // current value of arr[ptr] not empty string
    for this.ptr < this.size && len(this.arr[this.ptr]) > 0  {
        this.ptr++
    }
    return this.arr[beforeMovedPtr:this.ptr]
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
