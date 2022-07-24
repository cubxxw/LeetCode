func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    cur, num := 0, x
    for num != 0 {
        cur = cur * 10 + num % 10  
        num /= 10
    }
    return cur == x
}

// class Solution {
//     public boolean isPalindrome(int x) {
//         if(x < 0)
//             return false;
//         int cur = 0;
//         int num = x;
//         while(num != 0) {
//             cur = cur * 10 + num % 10;
//             num /= 10;
//         }
//         return cur == x;
//     }
// }