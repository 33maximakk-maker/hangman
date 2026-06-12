package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

var hangmanPics = []string{
    `
  +---+
      |
      |
      |
      |
      |
=========`,
    `
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
    `
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
    `
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
    `
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
    `
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
    `
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

var words = map[string][]string{
    "animals":     {"python", "elephant", "giraffe", "tiger", "zebra", "kangaroo"},
    "tech":        {"computer", "keyboard", "monitor", "algorithm", "function", "variable"},
    "programming": {"python", "javascript", "rust", "golang", "java", "typescript"},
}

func getRandomWord(category string) string {
    if catWords, ok := words[category]; ok {
        return catWords[rand.Intn(len(catWords))]
    }
    var all []string
    for _, w := range words {
        all = append(all, w...)
    }
    return all[rand.Intn(len(all))]
}

func displayHangman(attempts int) {
    fmt.Print(hangmanPics[attempts])
}

func displayWord(word string, guessed map[rune]bool) {
    for _, ch := range word {
        if guessed[ch] {
            fmt.Printf("%c ", ch)
        } else {
            fmt.Print("_ ")
        }
    }
    fmt.Println()
}

func main() {
    rand.Seed(time.Now().UnixNano())
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Welcome to Hangman!")
    fmt.Print("Choose category (animals/tech/programming) or Enter for random: ")
    scanner.Scan()
    category := strings.ToLower(scanner.Text())
    word := getRandomWord(category)
    guessed := make(map[rune]bool)
    attempts := 0
    maxAttempts := len(hangmanPics) - 1

    for attempts < maxAttempts {
        displayHangman(attempts)
        fmt.Println()
        displayWord(word, guessed)
        fmt.Print("Guessed letters: ")
        for k := range guessed {
            fmt.Printf("%c ", k)
        }
        fmt.Println()
        fmt.Printf("Attempts left: %d\n", maxAttempts-attempts)
        fmt.Print("Guess a letter: ")
        scanner.Scan()
        guess := strings.ToLower(scanner.Text())
        if len(guess) != 1 || !(guess[0] >= 'a' && guess[0] <= 'z') {
            fmt.Println("Invalid input. Please enter a single letter.")
            continue
        }
        g := rune(guess[0])
        if guessed[g] {
            fmt.Println("You already guessed that letter.")
            continue
        }
        guessed[g] = true
        if strings.ContainsRune(word, g) {
            fmt.Println("Good guess!")
            allGuessed := true
            for _, ch := range word {
                if !guessed[ch] {
                    allGuessed = false
                    break
                }
            }
            if allGuessed {
                displayHangman(attempts)
                fmt.Println()
                displayWord(word, guessed)
                fmt.Printf("Congratulations! You guessed the word: %s\n", word)
                break
            }
        } else {
            attempts++
            fmt.Println("Wrong guess!")
        }
        if attempts == maxAttempts {
            displayHangman(attempts)
            fmt.Printf("\nYou lost! The word was: %s\n", word)
        }
    }
}
