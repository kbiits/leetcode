// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        return firstBadVersionRecursive(1, n); 
    }
    
    int firstBadVersionRecursive(int left, int right) {
        
        if (left == right) {
            return left;
        }

        int mid = left + (right - left) / 2;    
        bool isBadMid = isBadVersion(mid);

        
        if (isBadMid) {
//             go to (left, mid)
            return firstBadVersionRecursive(left, mid);
        } else {
//             go to right
            return firstBadVersionRecursive(mid + 1, right);
        }
    }
};


//  1  2  3   4   5  6  7  
