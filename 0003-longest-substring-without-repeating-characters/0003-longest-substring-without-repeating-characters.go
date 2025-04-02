func lengthOfLongestSubstring(s string) int {
    sMap := make(map[rune]int)
    slice := make([]rune, 0)
    runes := []rune(s)

    var big, check int
    for i, r := range runes {
        if sMap[r] != 0 && check < sMap[r] {
            slice = runes[sMap[r]:i]
            check = sMap[r]
        } 
        slice = append(slice, r)
        sMap[r] = i + 1

        if big < len(slice) {
            big = len(slice) 
        }
    }

    

    return big
}