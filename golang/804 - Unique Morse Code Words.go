func uniqueMorseRepresentations(words []string) int {
    morses := [] string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    transformations := make(map[string]bool)
    for _,v := range words {
        var  morse string
        for _,s := range v {
            morse += morses[s - 'a']
        }
        transformations[morse] = true
    }
    return len(transformations)
}