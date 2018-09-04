import ("fmt"
       "bytes")
       
func uniqueMorseRepresentations(words []string) int {
    
    mors := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    
    morsMap := make(map[string] int)
    for _, word := range words {
        
        var buffer string
        
        for _, c := range word {
            buffer += mors[c - 'a']
        }
        morsMap[buffer] += 1
    }
    
    return len(morsMap)   
}
