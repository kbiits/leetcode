func replaceElements(arr []int) []int {
    n := len(arr)

    maxEl := -1
    for i := n - 1; i >= 0; i-- {
        temp := arr[i]
        arr[i] = maxEl

        if temp > maxEl {
            maxEl = temp
        }
    }
    
    return arr
}
