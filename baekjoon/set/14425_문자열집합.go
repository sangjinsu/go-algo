// https://www.acmicpc.net/problem/14425
// 문제: 문자열 집합
// 영어 이름: String Set
// 난이도: 실버

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	wordMap := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscanln(reader, &word)
		wordMap[word] = struct{}{}
	}

	var count uint
	for i := 0; i < m; i++ {
		var word string
		fmt.Fscanln(reader, &word)
		if _, ok := wordMap[word]; ok {
			count++
		}
	}

	fmt.Fprintln(writer, count)

}
