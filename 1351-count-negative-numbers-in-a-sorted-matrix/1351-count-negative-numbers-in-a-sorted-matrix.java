class Solution {
    public int countNegatives(int[][] grid) {
        if(grid.length == 0 || grid[0].length == 0){
            return 0;
        }
        int count = 0;
        for(int i=0,j=grid[i].length-1; i<grid.length && j>=0;){
            if(grid[i][j] >= 0){
                i++;
            }else{
                count += grid.length - i;
                j--;
            }
        }
    return count;
    } 
}