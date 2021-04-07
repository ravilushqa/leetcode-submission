func uniqueMorseRepresentations(words []string) int {
   dict := [26]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	res := 0
	uniq := make(map[string]struct{}, len(words))
	for _, word := range words {
		var morse string
		for _, c := range word {
			morse += dict[c-97]
		}
		if _, found := uniq[morse]; !found {
			uniq[morse] = struct{}{}
			res++
		}
	}
	
	return res 
}
