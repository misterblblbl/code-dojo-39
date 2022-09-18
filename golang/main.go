package main

import (
	"fmt"
	"strings"
)

func main() {
	printSong()
}

func printSong() (int, error) {
	songParts := []part{
		{
			"fly",
			".",
			"I don't know why she swallowed a fly - perhaps she'll die!",
			false,
		},
		{
			"spider",
			";",
			"That wriggled and wiggled and tickled inside her.",
			true,
		},
		{
			"bird",
			";",
			"How absurd to swallow a bird.",
			true,
		},
		{
			"cat",
			";",
			"Fancy that to swallow a cat!",
			true,
		},
		{
			"dog",
			";",
			"What a hog, to swallow a dog!",
			true,
		},
		{
			"cow",
			";",
			"I don't know how she swallowed a cow!",
			true,
		},
		{
			"horse",
			"...",
			"...She's dead, of course!",
			false,
		},
	}
	return fmt.Printf(generateSong(songParts))
}

func generateSong(songParts []part) string {
	song := make([]string, 0)
	animalsDone := make([]string, 0, len(songParts))

	for _, sp := range songParts {
		animalsDone = append(animalsDone, sp.animal)
		chorus := generateChorus(animalsDone)
		song = append(song, sp.combineLines(chorus)...)
	}

	return joinLines(song)
}

type part struct {
	animal         string
	punctuation    string
	animalLine     string
	requiresChorus bool
}

func (p *part) generateVerse() string {
	return fmt.Sprintf("There was an old lady who swallowed a %s%s", p.animal, p.punctuation)
}

func (p *part) combineLines(chorus []string) []string {
	start := p.generateVerse()

	part := []string{
		start,
		p.animalLine,
	}

	if !p.requiresChorus {
		return part
	}

	return append(part, chorus...)
}

func joinLines(song []string) string {
	return strings.Join(song, "\n")
}

func generateChorus(animals []string) []string {
	lastLine := "I don't know why she swallowed a fly - perhaps she'll die!"

	if len(animals) < 1 {
		return nil
	}

	if len(animals) < 2 {
		return []string{lastLine}
	}

	lines := make([]string, 0)

	for i := 1; i < len(animals); i++ {
		prevAnimal := animals[i-1]
		currAnimal := animals[i]
		punctuation := ","

		if i == 1 {
			punctuation = ";"
		}

		line := fmt.Sprintf("She swallowed the %s to catch the %s%s", currAnimal, prevAnimal, punctuation)
		lines = append(lines, line)
	}

	return append(reverseList(lines), lastLine)
}

func reverseList(list []string) []string {
	reversed := make([]string, len(list))

	for i := len(list) - 1; i >= 0; i-- {
		reversed[len(list)-1-i] = list[i]
	}

	return reversed
}
