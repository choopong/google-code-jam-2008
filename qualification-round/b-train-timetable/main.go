package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	FIND_NUM_CASE = iota
	FIND_TURNAROUND_TIME
	FIND_NUM_TRIP
	FIND_TRIP_A_TO_B
	FIND_TRIP_B_TO_A
)

type Trip struct {
	From     time.Time
	To       time.Time
	WillTurn bool
}

type ByFromTime []Trip

func (a ByFromTime) Len() int           { return len(a) }
func (a ByFromTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFromTime) Less(i, j int) bool { return a[i].From.Before(a[j].From) }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	numCase := 0
	countCase := 0
	turnaroundTime := 0
	numTripA := 0
	numTripB := 0
	tripAB := []Trip{}
	tripBA := []Trip{}
	countTrainA := 0
	countTrainB := 0
	mode := FIND_NUM_CASE

	endB := func() {

		sort.Sort(ByFromTime(tripAB))
		sort.Sort(ByFromTime(tripBA))
		// log.Printf("%v\n", tripAB)
		// log.Printf("%v\n", tripBA)
		for _, trAB := range tripAB {
			found := false
			for i, trBA := range tripBA {
				turnTime := trBA.To.Add(time.Duration(turnaroundTime) * time.Minute)
				// fmt.Println(turnTime)
				if !trBA.WillTurn && !turnTime.After(trAB.From) {
					trBA.WillTurn = true
					tripBA[i] = trBA
					found = true
					break
				}
			}
			if !found {
				countTrainA++
			}
		}
		for _, trBA := range tripBA {
			found := false
			for i, trAB := range tripAB {
				turnTime := trAB.To.Add(time.Duration(turnaroundTime) * time.Minute)
				if !trAB.WillTurn && !turnTime.After(trBA.From) {
					trAB.WillTurn = true
					tripAB[i] = trAB
					found = true
					break
				}

			}
			if !found {
				countTrainB++
			}
		}
		fmt.Fprintf(os.Stdout, "Case #%v: %v %v\n", countCase, countTrainA, countTrainB)
		if countCase == numCase {
			//Done
			os.Exit(0)
		} else {
			mode = FIND_TURNAROUND_TIME
		}
	}
	endA := func() {
		if numTripB == 0 {
			endB()
		} else {
			mode = FIND_TRIP_B_TO_A
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		switch mode {
		case FIND_NUM_CASE:
			// log.Println("FIND_NUM_CASE")
			if numCase, err = strconv.Atoi(line); err != nil {
				panic(err)
			}
			mode = FIND_TURNAROUND_TIME
		case FIND_TURNAROUND_TIME:
			// log.Println("FIND_TURNAROUND_TIME")
			if turnaroundTime, err = strconv.Atoi(line); err != nil {
				panic(err)
			}
			countCase++
			tripAB = []Trip{}
			tripBA = []Trip{}
			countTrainA = 0
			countTrainB = 0
			mode = FIND_NUM_TRIP
		case FIND_NUM_TRIP:
			// log.Println("FIND_NUM_TRIP")
			fmt.Fscanf(strings.NewReader(line), "%d %d", &numTripA, &numTripB)
			if numTripA != 0 {
				mode = FIND_TRIP_A_TO_B
			} else if numTripB != 0 {
				endA()
			} else {
				endB()
			}
		case FIND_TRIP_A_TO_B:
			// log.Println("FIND_TRIP_A_TO_B")
			tripAB = append(tripAB, Trip{
				From:     newTime(strings.Split(line, " ")[0]),
				To:       newTime(strings.Split(line, " ")[1]),
				WillTurn: false,
			})
			if len(tripAB) == numTripA {
				endA()
			}
		case FIND_TRIP_B_TO_A:
			// log.Println("FIND_TRIP_B_TO_A")
			// fmt.Println(line)
			tripBA = append(tripBA, Trip{
				From:     newTime(strings.Split(line, " ")[0]),
				To:       newTime(strings.Split(line, " ")[1]),
				WillTurn: false,
			})
			if len(tripBA) == numTripB {
				endB()
			}
		}
	}

}

func newTime(text string) time.Time {
	// fmt.Println(text)
	h, _ := strconv.Atoi(strings.Split(text, ":")[0])
	m, _ := strconv.Atoi(strings.Split(text, ":")[1])
	t := time.Date(1980, time.May, 25, h, m, 0, 0, time.UTC)
	return t
}
