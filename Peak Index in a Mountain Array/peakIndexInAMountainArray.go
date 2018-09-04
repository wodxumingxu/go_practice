func peakIndexInMountainArray(A []int) int {
    l,r, mid := 0, len(A), 0
    for l < r {
        mid = (l + r) / 2
        if A[mid] < A[mid + 1] {
            l = mid
        } else if A[mid - 1] > A[mid] {
            r = mid
        } else {
            return mid   
        }
    }
    return mid
}
