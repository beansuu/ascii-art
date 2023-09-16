package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadBannerFile(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

bannerMap := make(map[string][]string)
scanner := bufio.NewScanner(file)

var currentChar string
for scanner.Scan() {
	line := scanner.Text()
	if string.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
		currentChar = line[1 : len(line)-1]
		bannerMap[currentChar] = []string{}
	} else if currentChar != "" {
		bannerMap[currentChar] = append(bannerMap[currentChar], line)
	}
}
if err := scanner.Err(); err != nil {
	return nil, err
}
return bannerMap, nil
}

func ConvertStringToAscii(input string, bannerMap map[string][]string) string {
	var result []string
	height := len(bannerMap["A"])

	for i := 0; i < height; i++ {
		var line string
		for _, char := range input {
			strChar := string(char)
			if ascii, exists := bannerMap[strChar]; exists {
				line += ascii[i]
			} else {
				line += " "
			}
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

func main() {
	
}