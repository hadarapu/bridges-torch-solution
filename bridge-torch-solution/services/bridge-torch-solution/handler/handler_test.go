package handler

import (
	"fmt"
	"git.sqad.io/bridge-torch-solution/services/common"
	"testing"
)

func TestQuickestTime(t *testing.T) {
	var times []float64 =  []float64{1, 2, 4, 5, 10, 20, 50}
	fmt.Println(QuickestTime(times))
}

func TestSubCalc(t *testing.T) {
	bridges := make(map[int]int)
	bridges[0] = 100

	persons := make(map[int]float64)
	persons[0] = 5
	persons[1] = 2
	persons[2] = 1
	persons[3] = 20
	persons[4] = 10
	persons[5] = 4
	persons[6] = 100
	persons[7] = 50

	problem := make(map[int][]int)
	p := []int{0,2,3,5,6}
	problem[0] = p

	Cfg := &common.ConfigInfo{
		Persons: persons,
		Bridges: bridges,
		Problem: problem,
	}

	fmt.Println(SubCalc(200, p, Cfg))
}