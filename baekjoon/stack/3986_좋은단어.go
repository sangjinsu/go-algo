// https://www.acmicpc.net/problem/3986
// 문제: 좋은 단어
// 영어 이름: Good Word
// 난이도: 실버4

package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	items []string
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *stack) Len() int {
	return len(s.items)
}

func (s *stack) Flush() {
	s.items = make([]string, 0)
}

func (s *stack) Top() string {
	return s.items[len(s.items)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var count int

	for i := 0; i < n; i++ {
		stack := newStack()
		var word string
		fmt.Fscanln(reader, &word)
		for _, letter := range word {
			strLetter := string(letter)
			if stack.IsEmpty() || stack.Top() != strLetter {
				stack.Push(strLetter)
			} else if stack.Top() == strLetter {
				stack.Pop()
			}
		}

		if stack.IsEmpty() {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}
