package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the anagram prgram!\n Please enter the first phrase:")

	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')

	fmt.Println("Please enter the second phrase")
	reader = bufio.NewReader(os.Stdin)
	word2, _ := reader.ReadString('\n')

	if isAnagram(formatString(word), formatString(word2)) {
		fmt.Println("The phrases are anagram")
	} else {
		fmt.Println("The phrases are NOT anagram")
	}

}

func formatString(s string) string {
	str := strings.ReplaceAll(s, " ", "")
	str = strings.ToLower(str)
	return str
}

func isAnagram(word, word2 string) bool {
	chars := []rune(word)
	chars2 := []rune(word2)

	// add characters to the map with number of occurrencies
	charsMap := charToMap(chars)

	charsMap2 := charToMap(chars2)

	//compare maps
	//	return reflect.DeepEqual(charsMap, charsMap2)
	return eq(charsMap, charsMap2)
}

func eq(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}

func charToMap(chars []rune) map[string]int {
	charsMap := make(map[string]int)

	for _, j := range chars {
		charsMap[string(j)] = charsMap[string(j)] + 1
	}

	return charsMap

}
