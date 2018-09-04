import "strconv"

func isDividingNumber(n int) bool {
    str := strconv.Itoa(n)
    for _, ch := range str  {
        num := int(ch - '0')
        if 0 == num || n % num > 0  {
            return false   
        }
    }
    return true
}

func selfDividingNumbers(left int, right int) []int {
    var ret []int 
    for i := left; i <= right; i = i + 1 {
        if isDividingNumber(i) {
            ret = append(ret, i)
        }
    }
    return ret
}
