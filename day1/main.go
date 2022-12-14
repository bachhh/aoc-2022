package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var list = [][]int{}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	cur := []int{}
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			list = append(list, cur)
			cur = []int{} // reset
		} else {
			i, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			cur = append(cur, i)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	maxS := 0
	for i := range list {
		s := 0
		for j := range list[i] {
			s += list[i][j]
		}
		if s > maxS {
			maxS = s
		}
	}
	fmt.Println(maxS)
}
