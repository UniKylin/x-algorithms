// link		: https://leetcode-cn.com/problems/climbing-stairs/
// Author	: Kylin
// Date		: 2021-08-28
class Solution {
    public int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }
        int f1 = 1, f2 = 2, f3 = 3;

        for (int i = 3; i < n + 1; i++) {
            f3 = f1 + f2;
            f1 = f2;
            f2 =f3;
        }

        return f3;
    }
}