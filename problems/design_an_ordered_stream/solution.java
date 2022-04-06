class OrderedStream {

    String[] arr;
    int pointer;
    int n;
    
    public OrderedStream(int n) {
        this.pointer = 0;
        this.arr = new String[n];
        this.n = n;
    }
    
    public List<String> insert(int idKey, String value) {
        List<String> res = new ArrayList<>();

        this.arr[idKey - 1] = value;
        while (this.pointer < n && this.arr[this.pointer] != null) {
            res.add(this.arr[this.pointer]);
            this.pointer++;
        }

        return res;
    }
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * OrderedStream obj = new OrderedStream(n);
 * List<String> param_1 = obj.insert(idKey,value);
 */