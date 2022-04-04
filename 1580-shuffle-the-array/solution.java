class Solution {
    public int[] shuffle(int[] nums, int n) {
        // temp used to store the result array
        int[] temp = new int[2 * n];
        for (int i = 0; i < n; i++) {
            temp[2 * i] = nums[i];
            temp[2 * i + 1] = nums[i + n];
        }

        return temp;
    }
}

// 5 * 1000 + 4 = 50004

// 3
// 0, 3
// 1, 4
// 2, 5

// 0 => 0 / 2 = 0
// 1 => (1 / 2) + 3 = 3
// (0,3)

// 2 => 2 / 2 = 1
// 3 => (3 / 2) + 3 = 4
// (1, 4)

// 4 => 4 / 2 = 2
// 5 => (5 / 2) + 3 = 5
// (2, 5)
