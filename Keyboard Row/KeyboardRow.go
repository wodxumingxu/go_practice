func findWords(words []string) []string {
    ret := make([]string, 0)
    chrTbl := map[byte]int {
        'Q': 0,'W': 0,'E': 0,'R': 0,'T': 0,'Y': 0,'U': 0,'I': 0,'O': 0,'P': 0,
        'q': 0,'w': 0,'e': 0,'r': 0,'t': 0,'y': 0,'u': 0,'i': 0,'o': 0,'p': 0,
        'A': 1,'S': 1,'D': 1,'F': 1,'G': 1,'H': 1,'J': 1,'K': 1,'L': 1,
        'a': 1,'s': 1,'d': 1,'f': 1,'g': 1,'h': 1,'j': 1,'k': 1,'l': 1,
        'Z': 2,'X': 2,'C': 2,'V': 2,'B': 2,'N': 2,'M': 2,
        'z': 2,'x': 2,'c': 2,'v': 2,'b': 2,'n': 2,'m': 2}
    
    for _, word := range words {
        
        row := chrTbl[ word[0] ]
        bAligned := true
        for _, ch := range word {
            if row != chrTbl[byte(ch)] {
                bAligned = false
            }
        }
        
        if bAligned {
            ret = append(ret, word)
        }
    }
    
    return ret
}
