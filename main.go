package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const SymbolSize = 8

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ascii-art <text> [style_file]")
		os.Exit(1)
	}

	style := "standard.txt"
	if len(os.Args) > 2 {
		style = os.Args[2]
	}

	symbolMap, err := makeSymbolMap(style)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	printArt(os.Args[1], symbolMap)
}

func makeSymbolMap(style string) (map[rune][]string, error) {
	symbolMap := make(map[rune][]string)
	file, err := ioutil.ReadFile(style)
	if err != nil {
		return nil, err
	}

	artSymbols := strings.Split(string(file), "\n")
	for symbol := ' '; symbol <= '~'; symbol++ {
		j, slice := int(symbol-' '), make([]string, SymbolSize)
		for i := range slice {
			slice[i] = artSymbols[j*(SymbolSize+1)+i+1]
		}
		symbolMap[symbol] = slice
	}

	return symbolMap, nil
}

func printArt(input string, symbolMap map[rune][]string) {
	for _, line := range strings.Split(input, "\\n") {
		for i := 0; i < SymbolSize; i++ {
			for _, symbol := range line {
				fmt.Print(symbolMap[symbol][i])
			}
			fmt.Println()
		}
		if line == "" {
			fmt.Println()
		}
	}
}
