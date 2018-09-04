import "sort"

func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    
    ret := 0
    for idx, num := range(nums) {
        if 0 == idx % 2 {
            ret += num
        }
    }
    
    return ret
}
