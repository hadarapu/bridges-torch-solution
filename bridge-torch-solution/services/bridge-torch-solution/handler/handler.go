package handler

import (
	"fmt"
	. "git.sqad.io/bridge-torch-solution/services/common"
	"sort"
)

func CalculateQuickestTime(input chan *InputObject, output chan *OutputObject, conf *ConfigInfo) {
	for i := range input {

		bridgeLength := conf.Bridges[i.BridgeId]
		if bridgeLength > 0 {
			time := SubCalc(bridgeLength, i.PersonIdsList, conf)
			fmt.Println(fmt.Sprintf("optimal time taken for people: %v to cross bridge of length %d is %v", i.PersonIdsList, bridgeLength, time))
			output <- &OutputObject{
				BridgeId:     i.BridgeId,
				QuickestTime: time,
			}
		}
	}
}

func SubCalc(bridgeLength int, personsList []int, conf *ConfigInfo) float64{
	timesList :=  getSortedTimeList(bridgeLength, personsList, conf)
	return QuickestTime(timesList)
}

func getSortedTimeList(bridgeLength int, persons []int, conf *ConfigInfo) []float64 {
	speedList := make([]float64, 0, len(persons))

	for _, p := range persons {
		speedList = append(speedList, float64(bridgeLength)/conf.Persons[p])
	}
	sort.Float64s(speedList)
	return speedList
}

func QuickestTime(timesList []float64) float64{
	if len(timesList) == 1 {
		return timesList[0]
	}

	first, second := timesList[0], timesList[1]
	timesList = timesList[2:]

	var finalTime float64
	for len(timesList) >= 2 {
		newLen := len(timesList)
		temp1 := timesList[newLen-2]
		temp2 := timesList[newLen-1]

		timesList = timesList[:newLen-2]

		timeForFirstMethod := (2 * second) + first + temp2
		timeForSecondMethod := (2 * first) + temp1 + temp2

		if timeForFirstMethod < timeForSecondMethod {
			finalTime += timeForFirstMethod
		} else {
			finalTime += timeForSecondMethod
		}
	}

	if len(timesList) == 1 {
		finalTime += first + second + timesList[0]
	} else {
		finalTime += second
	}

	return finalTime

}