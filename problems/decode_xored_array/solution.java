class Solution {
    public int[] decode(int[] encoded, int first) {
        int[] res = new int[encoded.length + 1];
        int temp = first;
        res[0] = temp;

        for (int i = 0; i < encoded.length; i++) {
            temp = encoded[i] ^ temp;
            res[i + 1] = temp;
        }

        return res;
    }
}