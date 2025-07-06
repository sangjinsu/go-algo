// https://www.acmicpc.net/problem/1157
// 문제: 단어 공부

package main

import (
	"bufio"
	"os"
	"strings"
)

func solve(word string) rune {
	letters := strings.ToUpper(strings.TrimSpace(word))

	result := make([]rune, 0)
	letterMap := make(map[rune]uint)

	for _, letter := range letters {
		if _, ok := letterMap[letter]; ok {
			letterMap[letter]++
		} else {
			letterMap[letter] = 1
		}
	}

	var maxCount uint = 0
	for _, count := range letterMap {
		if maxCount < count {
			maxCount = count
		}
	}

	for letter, count := range letterMap {
		if count == maxCount {
			result = append(result, letter)
		}
	}

	if len(result) > 1 {
		return '?'
	}

	return result[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, _ := reader.ReadString('\n')

	result := solve(input)

	writer.WriteRune(result)
}
