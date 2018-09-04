import "strings"

func uncommonFromSentences(A string, B string) []string {
    ret := make(map[string]int)
    sliceA := strings.Split(A, " ")
    sliceB := strings.Split(B, " ")
    
    for _, wordA := range sliceA {
        ret[wordA] += 1 
    }
    
    for _, wordB := range sliceB {
        ret[wordB] += 1 
    }
    
    var res []string
    for key,val := range ret {
        if val == 1 {
            res = append(res, key)   
        }
    }
    
    return res
}
