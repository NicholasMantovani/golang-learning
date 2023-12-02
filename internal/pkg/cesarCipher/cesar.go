package cesarcipher

import (
	"fmt"
	"strings"
)


var letters []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "z"}

func ExecuteCesarCipher() {
	result := doCipher("test", 5)

	fmt.Println("The encrypted input is ", result)
}


func doCipher(input string, key int) string{

	result := ""
	
	for _, char := range input {
		for letterIndex, letter := range letters {
			if strings.EqualFold(string(char), string(letter)) {
				letterIndex = letterIndex + key
				if letterIndex >= len(letters) {
					letterIndex = letterIndex - len(letters)
				}
				result += letters[letterIndex]
				break; 
			}
		}
	}
	
	return result
}

