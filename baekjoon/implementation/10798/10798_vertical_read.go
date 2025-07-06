// https://www.acmicpc.net/problem/11720
// 문제: 세로 읽기

package main

import (
	"bufio"
	"os"
	"strings"
)

func solve(wordList []string) string {
	row := len(wordList)
	maxLen := 0
	for _, word := range wordList {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	var result string
	for i := 0; i < maxLen; i++ {
		for j := 0; j < row; j++ {
			if len(wordList[j]) <= i {
				continue
			}
			result += string(wordList[j][i])
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	wordList := make([]string, 5)
	for i := 0; i < 5; i++ {
		input, _ := reader.ReadString('\n')
		wordList[i] = strings.TrimSpace(input)
	}

	result := solve(wordList)
	writer.WriteString(result)
}
