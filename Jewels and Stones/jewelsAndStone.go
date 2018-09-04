func numJewelsInStones(J string, S string) int {
    var nCount int = 0   // define a variable with intialized value
    for _, char := range S {   // iterater over a string
        if strings.ContainsRune(J, char) {  //check if the string contains the rune 
            nCount = nCount + 1
        }
    }
    
    return nCount
}
