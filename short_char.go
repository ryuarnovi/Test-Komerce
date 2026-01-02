package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"example.com/psbbfamily"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'i' || r == 'u' || r == 'e' || r == 'o'
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input one line of words (S) : ")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(input)

	var vowels, consonants []rune

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		if isVowel(r) {
			vowels = append(vowels, r)
		} else {
			consonants = append(consonants, r)
		}
	}

	fmt.Println("Vowel Characters :", string(vowels))
	fmt.Println("Consonant Characters :", string(consonants))

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Input one line of words (S) : ")
	input, _ = reader.ReadString('\n')
	input = strings.ToLower(input)

	var vowels2, consonants2 []rune

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		if isVowel(r) {
			vowels2 = append(vowels2, r)
		} else {
			consonants2 = append(consonants2, r)
		}
	}

	fmt.Println("Vowel Characters :", string(vowels2))
	fmt.Println("Consonant Characters :", string(consonants2))
	psbbfamily.PsbbFamily()
}
