
var MinStack = function() {
    this.stack = [];
    this.minstack = [];
};

/** 
 * @param {number} val
 * @return {void}
 */
MinStack.prototype.push = function(val) {
    let len = this.stack.length;
    this.stack[len] = val;
    // check if minstack is empty or new val is the new minimum
    if (this.minstack.length === 0 || val <= this.minstack.at(-1)) {
        this.minstack[this.minstack.length] = val
    }
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
    let data = this.stack.pop();
    if (this.minstack.at(-1) === data) {
        this.minstack.pop();
    }
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
    return this.stack.at(-1);
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
    return this.minstack.at(-1);
};

/** 
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(val)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
