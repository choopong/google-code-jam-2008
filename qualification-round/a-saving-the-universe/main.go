package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	FIND_NUM_CASE_MODE = iota
	FIND_NUM_SE_MODE
	FIND_SE_MODE
	FIND_NUM_QUERY_MODE
	FIND_QUERY_MODE
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	noCase := 0
	numSearchEngine := 0
	numQuery := 0
	numSwitch := 0
	count := 0
	countCase := 0
	var err error
	seQuery := map[string]bool{}
	mode := FIND_NUM_CASE_MODE
	// lastSE := ""
	for scanner.Scan() {
		line := scanner.Text()
		switch mode {
		case FIND_NUM_CASE_MODE:
			// log.Println("FIND_NUM_CASE_MODE")
			if noCase, err = strconv.Atoi(line); err != nil {
				panic(err)
			}
			mode = FIND_NUM_SE_MODE

		case FIND_NUM_SE_MODE:
			// log.Println("FIND_NUM_SEARCH_ENGINE_MODE")
			if numSearchEngine, err = strconv.Atoi(line); err != nil {
				panic(err)
			}
			countCase++
			count = 0
			numSwitch = 0
			seQuery = map[string]bool{}
			mode = FIND_SE_MODE
		case FIND_SE_MODE:
			// log.Println("FIND_SEARCH_ENGINE_MODE")
			count++
			seQuery[line] = false
			if count == numSearchEngine {
				mode = FIND_NUM_QUERY_MODE
			}
		case FIND_NUM_QUERY_MODE:
			// log.Println("FIND_NUM_QUERY_MODE")
			if numQuery, err = strconv.Atoi(line); err != nil {
				panic(err)
			}
			if numQuery == 0 {
				fmt.Fprintf(os.Stdout, "Case #%v: %v\n", countCase, numSwitch)
				mode = FIND_NUM_SE_MODE
			} else {
				count = 0
				mode = FIND_QUERY_MODE
			}
		case FIND_QUERY_MODE:
			// log.Println("FIND_QUERY_MODE")
			count++
			seQuery[line] = true
			if foundAllSE(seQuery) {
				numSwitch++
				resetSegment(seQuery)
				seQuery[line] = true
			}
			if count == numQuery {
				fmt.Fprintf(os.Stdout, "Case #%v: %v\n", countCase, numSwitch)
				if countCase == noCase {
					//Done
					os.Exit(0)
				} else {
					mode = FIND_NUM_SE_MODE
				}
			}
		}
	}

}

func foundAllSE(seQuery map[string]bool) bool {
	for _, v := range seQuery {
		if !v {
			return false
		}
	}
	return true
}

func resetSegment(seQuery map[string]bool) {
	for k := range seQuery {
		seQuery[k] = false
	}
}
