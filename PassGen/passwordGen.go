package main

import (
	"fmt"
	"os"
	"sync"
)

type PasswordGenerator struct {
	possibleCombination int
	combinationType     int
	special             string
	numeric             string
	alphabet            string
	getCharacters       string
	numThreads          int
}

func NewPasswordGenerator(possibleCombination, combinationType, numThreads int) *PasswordGenerator {
	return &PasswordGenerator{
		possibleCombination: possibleCombination,
		combinationType:     combinationType,
		special:             `!"#$%&'()*+,-. /:;?@[]^_{|}~`,
		numeric:             `0123456789`,
		alphabet:            `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`,
		numThreads:          numThreads,
	}
}

func (pg *PasswordGenerator) generateGetCharacters() {
	switch pg.combinationType {
	case 1:
		pg.getCharacters = pg.numeric + pg.alphabet
	case 2:
		pg.getCharacters = pg.numeric
	case 3:
		pg.getCharacters = pg.alphabet
	case 4:
		pg.getCharacters = pg.special
	case 5:
		pg.getCharacters = pg.special + pg.numeric
	case 6:
		pg.getCharacters = pg.special + pg.numeric + pg.alphabet
	default:
		panic("Invalid combination_type")
	}
}

func (pg *PasswordGenerator) generatePassword(start, end int, output chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < end; i++ {
		password := make([]byte, pg.possibleCombination)
		for j := 0; j < pg.possibleCombination; j++ {
			password[j] = pg.getCharacters[i%len(pg.getCharacters)]
			i /= len(pg.getCharacters)
		}
		output <- string(password)
	}
}

func (pg *PasswordGenerator) generatePasswords() {
	pg.generateGetCharacters()
	passwords := make(chan string, pg.possibleCombination)
	var wg sync.WaitGroup

	for i := 0; i < pg.numThreads; i++ {
		start := i * pg.possibleCombination / pg.numThreads
		end := (i + 1) * pg.possibleCombination / pg.numThreads
		wg.Add(1)
		go pg.generatePassword(start, end, passwords, &wg)
	}

	go func() {
		wg.Wait()
		close(passwords)
	}()

	fileName := "password_list.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for password := range passwords {
		fmt.Fprintf(file, "%s\n", password)
		fmt.Printf("Possible combination: %s\n", password)
	}
}

func main() {
	var possibleCombination, combinationType, numThreads int

	fmt.Print("How many password combinations do you want to create? Exp(3): ")
	fmt.Scan(&possibleCombination)

	fmt.Print("Enter combination type (1-6): ")
	fmt.Scan(&combinationType)

	fmt.Print("How many threads to use? ")
	fmt.Scan(&numThreads)

	generator := NewPasswordGenerator(possibleCombination, combinationType, numThreads)
	generator.generatePasswords()
}
