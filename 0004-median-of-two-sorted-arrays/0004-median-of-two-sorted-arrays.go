func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    arr := append(nums1, nums2...)
    sort.Ints(arr)

    if len(arr)%2 == 1 {
        return float64(arr[len(arr)/2])
    } else {
        n1 := arr[len(arr)/2-1]
        n2 := arr[len(arr)/2]
        return float64(n1 + n2) / 2.0
    }
}