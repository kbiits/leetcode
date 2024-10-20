func moveZeroes(nums []int)  {

    ptrToZero := -1
    for i, v := range nums {
        if v == 0 {
            ptrToZero = i
            break
        }
    }

    if ptrToZero == -1 {
        return
    }

    ptr := ptrToZero + 1
    for ptr < len(nums) {
        if nums[ptr] == 0 {
            ptr++
        } else {
            // fmt.Printf("ptr: %d, ptrtozero: %d, nums %v\n", ptr, ptrToZero, nums)
            swap(ptr, ptrToZero, nums)
            // fmt.Println(nums)
            ptrToZero++
            ptr++
        }
    }

}

func swap(i, j int, nums []int) {
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}
