package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	pass, err := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pass)
}

type Config struct {
	Types  []string
	Seed   int64
	Length uint
}

func run() (string, error) {

	withLower := flag.Bool("l", false, "it adds lowercase carachters")
	withUpper := flag.Bool("u", false, "it adds uppercase carachters")
	withNumber := flag.Bool("n", false, "it adds numbers")
	withExtra := flag.Bool("e", false, "it adds special carachters")
	passLength := flag.Int("size", 10, "the password size the default is 10")

	flag.Parse()

	selected := []string{}

	if *withLower {
		selected = append(selected, "lower")
	}

	if *withUpper {
		selected = append(selected, "upper")
	}

	if *withNumber {
		selected = append(selected, "numbers")
	}

	if *withExtra {
		selected = append(selected, "extra")
	}

	if len(selected) == 0 {
		selected = []string{"lower", "upper", "numbers", "extra"}
	}

	if *passLength > 100 {
		return fmt.Sprint("the password size can't be greater then 100"), nil
	}

	chars := map[string]string{
		"lower":   "abcdefghijklmnopqrstuvwxyz",
		"upper":   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"numbers": "0123456789",
		"extra":   "!@#$%^&*?",
	}

	config := &Config{
		Types:  selected,
		Seed:   time.Now().UnixNano(),
		Length: uint(*passLength),
	}

	return GenPassword(chars, config)
}

func GenPassword(characters map[string]string, config *Config) (string, error) {

	if len(config.Types) <= 0 {
		return "", nil
	}

	for _, t := range config.Types {
		if _, ok := characters[t]; !ok {
			return "", fmt.Errorf("the type %s is not present on characters map", t)
		}
	}

	rand.Seed(config.Seed)
	pass := strings.Builder{}

	for i := 0; i < int(config.Length); i++ {
		typeIndex := rand.Intn(len(config.Types))
		charType := config.Types[typeIndex]

		chars := characters[charType]
		charIndex := rand.Intn(len(chars))

		pass.WriteString(string(chars[charIndex]))
	}

	return pass.String(), nil
}
