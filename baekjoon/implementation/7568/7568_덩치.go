// https://www.acmicpc.net/problem/14467
// 문제: 덩치
// 유형 : 구현, 브루트포스

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	weightHeightList := make([][]uint8, 0, n)
	for i := 0; i < n; i++ {
		var weight, height uint8
		fmt.Fscanln(reader, &weight, &height)
		weightHeightList = append(weightHeightList, []uint8{weight, height})
	}

	rankingList := make([]string, 0, n)
	for i, wh1 := range weightHeightList {
		ranking := 1
		for j, wh2 := range weightHeightList {
			if i == j {
				continue
			}
			if wh1[0] < wh2[0] && wh1[1] < wh2[1] {
				ranking++
			}
		}
		rankingList = append(rankingList, strconv.Itoa(ranking))
	}

	fmt.Fprintln(writer, strings.Join(rankingList, " "))
}
