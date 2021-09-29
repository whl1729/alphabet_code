package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	codeTable := generateCodeTable()
	guessAlphabetCode(codeTable)
	fmt.Println("Goodbye, Grand Master of Memory~~")
}

func generateCodeTable() map[string]string {
	file, err := os.Open("/home/along/live/memory/code/alphabet_code.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	codeTable := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addToCodeTable(codeTable, scanner.Text())
	}

	return codeTable
}

func addToCodeTable(codeTable map[string]string, line string) {
	if len(line) == 0 || line[0] == '#' {
		return
	}

	words := split(line, " ")
	if len(words) >= 2 {
		codeTable[words[0]] = words[1]
	}
}

func split(line, delimiter string) []string {
	words := strings.Split(line, delimiter)
	result := []string{}
	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(word) > 0 {
			result = append(result, word)
		}
	}
	return result
}

func guessAlphabetCode(codeTable map[string]string) {
	answer := ""
	for {
		alphabet := generateAlphabet()
		fmt.Println("----------------------------------------------------")
		fmt.Printf("Please enter the code for %s: ", alphabet)
		fmt.Scanf("%s", &answer)
		fmt.Printf("\nThe true answer is: %s\n\n", codeTable[alphabet])
		if len(answer) >= 2 && strings.Contains(codeTable[alphabet], answer) {
			fmt.Printf("You're right, Congratulations!\n\n")
		} else {
			fmt.Printf("Sorry, you're wrong.\n\n")
		}

		fmt.Printf("Would you like to play again? Enter 'y' or 'n': ")
		fmt.Scanf("%s", &answer)
		fmt.Println()
		if len(answer) == 0 || unicode.ToLower(rune(answer[0])) != 'y' {
			break
		}
	}
}

func generateAlphabet() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	result := ""
	for i := 0; i < 2; i++ {
		result += string(random.Intn(26) + 97)
	}
	return result
}
