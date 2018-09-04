import "strings"
import "strconv"
import "fmt"

func subdomainVisits(cpdomains []string) []string {
    ret := make(map[string]int)
    var counts []string
    for _, domain := range cpdomains {
        parsed := strings.Split(domain, " ")
        domainSegs := strings.Split(parsed[1], ".")
        
        for idx := range domainSegs {
            key := strings.Join(domainSegs[idx : ], ".")
            
            count, err := strconv.Atoi(parsed[0])
            if err == nil {
                ret[key] += count   
            }
        }
    }
    
    for key, val := range ret {
        counts = append(counts, strconv.Itoa(val) + " " + key)
    }
    
    return counts
}
