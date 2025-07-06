// https://www.acmicpc.net/problem/1620
// 문제: 나는야 포켓몬 마스터 이다솜
// 영어 이름: I’m a Pokémon Master (Idasom)
// 난이도: 실버4

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(nameList []string, nameMap map[string]string, input string) string {
	if number, err := strconv.Atoi(input); err == nil {
		return nameList[number]
	}
	return nameMap[input]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	inputs := strings.Split(input, " ")
	n, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	m, err := strconv.Atoi(strings.TrimRight(inputs[1], "\n"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	nameList := make([]string, 0, n)
	nameList = append(nameList, "")
	nameMap := make(map[string]string, n)

	for i := 1; i < n+1; i++ {
		name, nameReadErr := reader.ReadString('\n')
		if nameReadErr != nil {
			fmt.Fprintln(os.Stderr, nameReadErr)
		}

		name = strings.TrimRight(name, "\n ")

		nameList = append(nameList, name)

		index := strconv.Itoa(i)
		nameMap[name] = index
	}

	result := make([]string, 0, m)
	for i := 0; i < m; i++ {
		question, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		question = strings.TrimRight(question, "\n")
		result = append(result, solve(nameList, nameMap, question))
	}

	fmt.Fprintln(writer, strings.Join(result, "\n"))

}
