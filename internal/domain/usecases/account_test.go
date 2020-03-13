package usecases

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCaseType struct {
	testData       string
	expectedResult string
}

var testCases = []TestCaseType{
	{
		testData:       "password",
		expectedResult: "affa1d444373d54a2579e4d84be48cc5",
	},
	{
		testData:       "01ozad123",
		expectedResult: "0ee06a5deeda94981d3662b341072d4f",
	},
	{
		testData:       "%%$!@Ddasd",
		expectedResult: "6035594b73fa18fb0bbf7e673187a1f8",
	},
}

func TestHashPassword(t *testing.T) {
	for _, testCase := range testCases {
		assert.Equal(t, HashPassword(testCase.testData), testCase.expectedResult)
	}
}

func TestCheckPassword(t *testing.T) {
	for _, testCase := range testCases {
		assert.Equal(t, CheckPassword(testCase.testData, testCase.expectedResult), true)
	}
}
