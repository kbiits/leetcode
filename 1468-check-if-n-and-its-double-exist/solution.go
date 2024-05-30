func checkIfExist(arr []int) bool {
    memory := map[int]int{}

    for i, v := range arr {
        memory[v] = i
    }

    for i := 0; i < len(arr); i++ {
        if j, ok := memory[arr[i]*2]; ok && j != i {
            return true
        }
    }

    return false
}
