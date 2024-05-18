type SubrectangleQueries struct {
    arr [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
    return SubrectangleQueries{rectangle}
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    lenRow := len(this.arr)
    lenCol := len(this.arr[0])

    for row := row1; row < lenRow && row <= row2; row++ {
        for col := col1; col < lenCol && col <= col2; col++ {
            this.arr[row][col] = newValue
        }
    }
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.arr[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
