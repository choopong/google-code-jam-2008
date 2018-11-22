package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	FIND_NUM_CASE = iota
	FIND_NUM_INT
	FIND_VECTOR_X
	FIND_VECTOR_Y
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numCase := 0
	countCase := 0
	numInt := 0
	vectorX := []string{}
	vectorY := []string{}
	mode := FIND_NUM_CASE
	for scanner.Scan() {
		line := scanner.Text()
		switch mode {
		case FIND_NUM_CASE:
			numCase, _ = strconv.Atoi(line)
			mode = FIND_NUM_INT
		case FIND_NUM_INT:
			countCase++
			numInt, _ = strconv.Atoi(line)
			mode = FIND_VECTOR_X
		case FIND_VECTOR_X:
			vectorX = strings.Split(line, " ")
			mode = FIND_VECTOR_Y
		case FIND_VECTOR_Y:
			vectorY = strings.Split(line, " ")
			xInts := []int{}
			yInts := []int{}
			for i := 0; i < numInt; i++ {
				x, _ := strconv.Atoi(vectorX[i])
				xInts = append(xInts, x)
				y, _ := strconv.Atoi(vectorY[i])
				yInts = append(yInts, y)
			}
			sort.Ints(xInts)
			sort.Ints(yInts)
			// log.Println(xInts, yInts)
			result := 0
			for i, x := range xInts {
				y := yInts[numInt-1-i]
				result += x * y
			}
			fmt.Fprintf(os.Stdout, "Case #%v: %v\n", countCase, result)
			if countCase == numCase {
				//Done
				os.Exit(0)
			} else {
				mode = FIND_NUM_INT
			}
		}
	}
}
