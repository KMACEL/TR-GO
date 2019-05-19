package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	randomMinValue  int
	randomMaxValue  int
	sampleDataCount int
	trueDataValue   int
	falseDataValue  int
)

func init() {
	randomMinValue = 1
	randomMaxValue = 51
	sampleDataCount = 15
	trueDataValue = 8
	falseDataValue = 7

	if randomSub := randomMaxValue - randomMinValue; randomSub <= sampleDataCount ||
		randomMaxValue < randomMinValue ||
		(trueDataValue+falseDataValue) != sampleDataCount {
		panic("Variable Check Please!!!")
	}
}

func main() {
begin:
	file, err := os.Open("TamVeri.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var counter int

	dataList := make(map[int][]string)
	for scanner.Scan() {
		getData := strings.Split(scanner.Text(), ",")
		dataList[counter] = getData
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var getRandom = func() int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(randomMaxValue-randomMinValue) + randomMinValue
	}

	var listControl = func(list []int, control int) bool {
		for _, value := range list {
			if value == control {
				return true
			}
		}
		return false
	}

	selectedData := make(map[int][]string)
	randomNumberList := make([]int, sampleDataCount)
	var trueData int
	var falseData int

	for i := 0; i < sampleDataCount; i++ {
	retry:
		randomValue := getRandom()
		if !listControl(randomNumberList, randomValue) {
			randomNumberList[i] = randomValue
		} else {
			goto retry
		}

		if getData := dataList[i]; getData[len(getData)-1] == "Hayır" &&
			falseData < falseDataValue {
			falseData++
			selectedData[i] = dataList[randomValue]
		} else if trueData < trueDataValue {
			trueData++
			selectedData[i] = dataList[randomValue]
		}
	}

	if falseData != falseDataValue || trueData != trueDataValue {
		log.Println("Selected Data Error, return begin...")
		file.Close()
		goto begin
	}

	findS := make([]string, len(selectedData[0]))

	for i := 0; i < len(findS); i++ {
		findS[i] = "0"
	}

	fmt.Println(selectedData)

	for i, value := range selectedData {
		if i == 0 {
			continue
		}
		if value[len(value)-1] == "Hayır" {
			continue
		}

		for i, col := range value {

			if findS[i] == "?" {
				continue
			}
			if findS[i] == "0" {
				findS[i] = col
				continue
			}
			if findS[i] != col {
				findS[i] = "?"
			}
		}
		fmt.Println(findS[1:len(findS)-1], value[1:len(findS)-1])
	}

	fmt.Println("Sonuç : ", findS[1:len(findS)-1])

	fmt.Println("######################################################")

	adayP := make([]string, len(selectedData[0]))

	for i := 0; i < len(adayP); i++ {
		adayP[i] = "0"
	}

	adayN := make([]string, len(selectedData[0]))
	for i := 0; i < len(adayN); i++ {
		adayN[i] = "?"
	}

	for _, value := range selectedData {
		if value[len(value)-1] == "Evet" {
			for i, col := range value {

				if adayP[i] == "?" {
					continue
				}
				if adayP[i] == "0" {
					adayP[i] = col
					continue
				}
				if adayP[i] != col {
					adayP[i] = "?"
				}
			}
			//fmt.Println(adayP[1:len(adayP)-1], value)
		}
		for i, valueElem := range value {
			if adayP[i] == "?" {
				adayN[i] = "?"
				continue
			} else if adayP[i] == valueElem {
				adayN[i] = "?"
				continue
			} else if adayP[i] == "0" {
				adayN[i] = valueElem
			} else {
				adayN[i] = valueElem
			}
		}

	}
	fmt.Println(adayP[1:len(adayP)-1], adayN[1:len(adayN)-1])
}
