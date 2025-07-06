// https://www.acmicpc.net/problem/14467
// 문제: 전구
// 유형 : 구현

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(bulbs []string, command, cond1, cond2 int) ([]string, error) {
	switch command {
	case 1:
		bulb := strconv.Itoa(cond2)
		bulbs[cond1-1] = bulb
	case 2:
		for i := cond1 - 1; i < cond2; i++ {
			if bulbs[i] == "0" {
				bulbs[i] = "1"
			} else {
				bulbs[i] = "0"
			}
		}
	case 3:
		for i := cond1 - 1; i < cond2; i++ {
			bulbs[i] = "0"
		}
	case 4:
		for i := cond1 - 1; i < cond2; i++ {
			bulbs[i] = "1"
		}
	default:
		return nil, fmt.Errorf("command error")
	}

	return bulbs, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	bulbs := make([]string, 0, n)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	inputs := strings.Split(input, " ")
	for _, bulb := range inputs {
		bulbs = append(bulbs, bulb)
	}

	var command, cond1, cond2 int
	var err error
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &command, &cond1, &cond2)
		bulbs, err = solve(bulbs, command, cond1, cond2)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	fmt.Fprintln(writer, strings.Join(bulbs, " "))

}
