package main

import (
	"log"
	"testing"
)

func Test_findConsecutiveSum(t *testing.T) {
	var testCases = []struct {
		arr            []int
		expectedResult int
	}{
		{
			arr:            []int{-2, 2, -2, 7, 8},
			expectedResult: 15,
		},
		{
			arr:            []int{-4, -2, -2, -1, -10},
			expectedResult: -1,
		},
		{
			arr:            []int{1, 1, -100, 3, -3, 2, -7, 2},
			expectedResult: 3,
		},
	}

	for _, testCase := range testCases {
		res := findConsecutiveSum(testCase.arr)
		if res != testCase.expectedResult {
			log.Panicf("Expected %d, Got %d", testCase.expectedResult, res)
		}
	}
}

func Test_findMaxSumInArray(t *testing.T) {
	var testCases = []struct {
		arr            []int
		expectedResult int
	}{
		{
			arr:            []int{1, 2, -2},
			expectedResult: 3,
		},
		{
			arr:            []int{-1, -2, -2},
			expectedResult: -1,
		},
		{
			arr:            []int{1, 2, 3},
			expectedResult: 6,
		},
	}

	for _, testCase := range testCases {
		res := findMaxSumInArray(testCase.arr)
		if res != testCase.expectedResult {
			log.Panicf("Expected %d, Got %d", testCase.expectedResult, res)
		}
	}
}
