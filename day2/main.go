package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	sumOf := []int{}
	for i := range list {
		s := 0
		for j := range list[i] {
			s += list[i][j]
		}
		sumOf = append(sumOf, s)
	}
	sort.Ints(sumOf)
	n, n1, n2 := len(sumOf)-1, len(sumOf)-2, len(sumOf)-3
	fmt.Println(sumOf[n] + sumOf[n1] + sumOf[n2])
}
