package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const baseAlphabetExtras = ".,_"
const baseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + baseAlphabetExtras

func sanitize(s string) string {
	// Convert to upper case
	s = strings.ToUpper(s)

	// Remove all non-alphabetic characters except spaces and periods
	var result strings.Builder
	for _, c := range s {
		if c == ' ' {
			result.WriteRune('_')
		} else if (unicode.IsLetter(c) && c >= 'A' && c <= 'Z') || contains(baseAlphabetExtras, c) {
			result.WriteRune(c)
		}
	}
	return result.String()
}

func readAndSanitizeFile(filePath string) string {
	// Read the entire file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Sanitize the file content
	finalString := sanitize(string(data))

	return finalString
}

func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}

func createKeyedAlphabet(key string) string {
	// Create a keyed alphabet
	keyedAlphabet := key
	for i := 0; i < len(baseAlphabet); i++ {
		if !contains(keyedAlphabet, rune(baseAlphabet[i])) {
			keyedAlphabet += string(baseAlphabet[i])
		}
	}
	if len(keyedAlphabet) != len(baseAlphabet) {
		log.Fatal("Alphabet key has repeated characters")
	}
	return keyedAlphabet
}

func getShiftedAlphabet(alphabet string, shift int) string {
	// Shift the alphabet by shift
	shiftedAlphabet := alphabet[shift:] + alphabet[:shift]
	return shiftedAlphabet
}

func getCypherChar(plainChar rune, keyChar rune, alphabet string) rune {
	// Get the cypher character for the given plain character and key character
	plainIndex := strings.IndexRune(alphabet, plainChar)
	keyIndex := strings.IndexRune(alphabet, keyChar)
	cypherChar := alphabet[(plainIndex+keyIndex)%len(alphabet)]
	return rune(cypherChar)
}

func encrypt(plain string, key string, alphabet string) string {
	// Encrypt the plain text using the key and alphabet
	var encryptedText strings.Builder
	for i := 0; i < len(plain); i++ {
		encryptedText.WriteRune(getCypherChar(rune(plain[i]), rune(key[i%len(key)]), alphabet))
	}
	return encryptedText.String()
}

func printAlphabetTable(alphabet string) {
	// fmt.Print(" |", alphabet, "\n")
	// fmt.Print(strings.Repeat("_", len(alphabet)+2), "\n")
	for i := 0; i < len(alphabet); i++ {
		shiftedAlphabet := getShiftedAlphabet(alphabet, i)
		// fmt.Printf("%d|", i)
		// fmt.Printf("%c|", rune(alphabet[i]))
		fmt.Println(shiftedAlphabet)
	}
}

func main() {
	// Read and sanitize the files
	plain := readAndSanitizeFile("plain.txt")
	key := readAndSanitizeFile("key.txt")
	alphakey := readAndSanitizeFile("alphakey.txt")

	// Create a keyed alphabet by moving the key to the front of the alphabet
	alphabet := createKeyedAlphabet(alphakey)

	// Print the keyed alphabet in a table
	printAlphabetTable(alphabet)
	fmt.Println()

	// Repeat the key until it is the same length as the plain text
	encrypted := encrypt(plain, key, alphabet)
	fmt.Println(encrypted)

	// Write the encrypted text to a file
	err := os.WriteFile("encrypted.txt", []byte(encrypted), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
