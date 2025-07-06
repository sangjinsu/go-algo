// https://www.acmicpc.net/problem/11720
// 문제: 숫자의 합
package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_, _ = reader.ReadString('\n')
	input, _ := reader.ReadString('\n')

	numbers := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		numbers[i], _ = strconv.Atoi(string(input[i]))
	}

	result := 0
	for _, number := range numbers {
		result += number
	}

	writer.WriteString(strconv.Itoa(result))

}
