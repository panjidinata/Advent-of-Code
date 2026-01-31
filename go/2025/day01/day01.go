package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GetNumOfDialAtZero(rotationSeq []string) (int, error) {
	if len(rotationSeq) == 0 {
		return 0, errors.New("rotation sequence is empty")
	}

	currHead := 50
	result := 0

	for _, seq := range rotationSeq {
		if strings.HasPrefix(seq, "L") {
			after, _ := strings.CutPrefix(seq, "L")
			if num, err := strconv.Atoi(after); err == nil {
				currHead = currHead - num
				currHead = currHead % 100

				if currHead == 0 {
					result += 1
				} else if currHead < 0 {
					currHead = currHead + 100
				}

			}
		} else if strings.HasPrefix(seq, "R") {
			after, _ := strings.CutPrefix(seq, "R")
			if num, err := strconv.Atoi(after); err == nil {
				currHead = currHead + num
				currHead = currHead % 100

				if currHead == 0 {
					result += 1
				}
			}
		}
	}

	return result, nil
}

func main() {
	fmt.Println("Hello, World")
}
