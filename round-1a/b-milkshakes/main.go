package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FIND_NUM_CASE = iota
	FIND_NUM_FLAVOR
	FIND_NUM_CUSTOMER
	FIND_CUSTOMER_MS
)

const (
	UNMALTED = 0
	MALTED   = 1
)

type MS struct {
	Flavor int
	Malted int
}

type CustomerMSList struct {
	CustomerID     int
	UnmaltedMSList []MS
	MaltedMS       MS
	rawData        string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numCase := 0
	countCase := 0
	numFlavor := 0
	numCustomer := 0
	mode := FIND_NUM_CASE
	countCustomer := 0
	customerMSMap := map[int]*CustomerMSList{}
	customerHappy := map[int]bool{}
	batch := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		switch mode {
		case FIND_NUM_CASE:
			numCase, _ = strconv.Atoi(line)
			mode = FIND_NUM_FLAVOR
		case FIND_NUM_FLAVOR:
			countCase++
			countCustomer = 0
			customerMSMap = map[int]*CustomerMSList{}
			batch = map[int]int{}
			customerHappy = map[int]bool{}
			numFlavor, _ = strconv.Atoi(line)
			mode = FIND_NUM_CUSTOMER
		case FIND_NUM_CUSTOMER:
			numCustomer, _ = strconv.Atoi(line)
			mode = FIND_CUSTOMER_MS
		case FIND_CUSTOMER_MS:
			countCustomer++
			customerMSMap[countCustomer] = &CustomerMSList{}
			customerMSMap[countCustomer].CustomerID = countCustomer
			customerMSMap[countCustomer].rawData += line
			values := strings.Split(line, " ")
			numMS, _ := strconv.Atoi(values[0])
			for i := 1; i <= numMS; i++ {
				flavor, _ := strconv.Atoi(values[(i*2)-1])
				malted, _ := strconv.Atoi(values[(i * 2)])
				ms := MS{
					Flavor: flavor,
					Malted: malted,
				}
				if malted == MALTED {
					customerMSMap[countCustomer].MaltedMS = ms
				} else {
					customerMSMap[countCustomer].UnmaltedMSList = append(customerMSMap[countCustomer].UnmaltedMSList, ms)
				}
			}

			if countCustomer == numCustomer {
				// Cal here
				for i := 1; i <= numFlavor; i++ {
					batch[i] = -1
				}
				result := ""
				// Each Customer with malted only
				for customerID, customerMSList := range customerMSMap {
					// Has Malted?
					if customerMSList.MaltedMS.Flavor != 0 && len(customerMSList.UnmaltedMSList) == 0 {
						batch[customerMSList.MaltedMS.Flavor] = MALTED
						customerHappy[customerID] = true
						continue
					}
				}
				// Each unhappy Customer with malted as well
				for customerID, customerMSList := range customerMSMap {
					// Has Malted?
					if !customerHappy[customerID] {
						if customerMSList.MaltedMS.Flavor != 0 {
							if batch[customerMSList.MaltedMS.Flavor] == MALTED {
								customerHappy[customerID] = true
								continue
							}
						}
					}
				}

				// Each unhappy Customer with unmalted only one
				for customerID, customerMSList := range customerMSMap {
					if !customerHappy[customerID] {
						if len(customerMSList.UnmaltedMSList) == 1 && customerMSList.MaltedMS.Flavor == 0 {
							if batch[customerMSList.UnmaltedMSList[0].Flavor] == -1 {
								batch[customerMSList.UnmaltedMSList[0].Flavor] = UNMALTED
								customerHappy[customerID] = true
								continue
							}
						}

					}
				}

			CustomerUnmaltedMSListLoop:
				for customerID, customerMSList := range customerMSMap {
					if !customerHappy[customerID] {
						for _, ms := range customerMSList.UnmaltedMSList {
							if batch[ms.Flavor] == UNMALTED || batch[ms.Flavor] == -1 {
								batch[ms.Flavor] = UNMALTED
								customerHappy[customerID] = true
								// Go next customer
								continue CustomerUnmaltedMSListLoop
							}
						}

					}
				}

				for customerID, customerMSList := range customerMSMap {
					if !customerHappy[customerID] {
						if customerMSList.MaltedMS.Flavor != 0 && (batch[customerMSList.MaltedMS.Flavor] == -1 || batch[customerMSList.MaltedMS.Flavor] == MALTED) {
							batch[customerMSList.MaltedMS.Flavor] = MALTED
							customerHappy[customerID] = true
							continue
						}
						result = "IMPOSSIBLE"
						break
					}
				}

				// If not success, try to swap priority
				if result == "IMPOSSIBLE" {
					result = ""
					customerHappy = map[int]bool{}
					for i := 1; i <= numFlavor; i++ {
						batch[i] = -1
					}
					// Each unhappy Customer with malted is first
					for customerID, customerMSList := range customerMSMap {
						if !customerHappy[customerID] {

							if customerMSList.MaltedMS.Flavor != 0 && (batch[customerMSList.MaltedMS.Flavor] == -1 || batch[customerMSList.MaltedMS.Flavor] == MALTED) {
								batch[customerMSList.MaltedMS.Flavor] = MALTED
								customerHappy[customerID] = true
								continue
							}

						}
					}
				CustomerUnmaltedMSListLoop2:
					for customerID, customerMSList := range customerMSMap {
						if !customerHappy[customerID] {
							for _, ms := range customerMSList.UnmaltedMSList {
								if batch[ms.Flavor] == UNMALTED || batch[ms.Flavor] == -1 {
									batch[ms.Flavor] = UNMALTED
									customerHappy[customerID] = true
									// Go next customer
									continue CustomerUnmaltedMSListLoop2
								}
							}
							result = "IMPOSSIBLE"
							break
						}
					}
				}

				if result != "IMPOSSIBLE" {
					for i := 1; i <= numFlavor; i++ {
						malted := batch[i]
						if malted == -1 {
							malted = 0
						}
						if i != 1 {
							result += " "
						}
						result += strconv.Itoa(malted)
					}
				}
				fmt.Fprintf(os.Stdout, "Case #%v: %v\n", countCase, result)

				if countCase == numCase {
					//Done
					os.Exit(0)
				} else {
					mode = FIND_NUM_FLAVOR
				}
			}
		}
	}
}
