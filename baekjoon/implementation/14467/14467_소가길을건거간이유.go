// https://www.acmicpc.net/problem/14467
// 문제: 소가 길을 건너간 이유
// 유형 : 구현

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

	var n int
	_, err := fmt.Fscanln(reader, &n)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var count uint8
	cowMoveList := make(map[uint8]uint8)
	for i := 0; i < n; i++ {
		var cow, land uint8
		_, err = fmt.Fscanln(reader, &cow, &land)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if _, ok := cowMoveList[cow]; !ok {
			cowMoveList[cow] = land
		} else {
			if cowMoveList[cow] != land {
				count++
				cowMoveList[cow] = land
			}
		}
	}

	_, err = fmt.Fprintln(writer, count)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
