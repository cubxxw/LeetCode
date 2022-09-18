class Solution {
    public boolean isAnagram(String s, String t) {
     //对数组进行排序，如果排序后的数组是一样的，说明√
        if(s.length()!=t.length()){
             return false;    
        }
        //相等继续判断，先排序
        char[] str1 = s.toCharArray();
        char[] str2 = t.toCharArray();
        Arrays.sort(str1);
        Arrays.sort(str2);
        // if(Arrays.equals(str1,str2)) {
        //     return true;
        // }else{
        //     return false;
        // }
        return Arrays.equals(str1, str2);
    }
}