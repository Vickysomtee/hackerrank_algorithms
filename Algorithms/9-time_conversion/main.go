package main

import "fmt"

// Test case 1
var time = "07:05:45PM"

// Test case 2
// var time = "12:45:54PM"

func timeConversion(s string) {
	meridiemPm := map[string]string{
		"01": "13",
		"02": "14",
		"03": "15",
		"04": "16",
		"05": "17",
		"06": "18",
		"07": "19",
		"08": "20",
		"09": "21",
		"10": "22",
		"11": "23",
		"12": "12",
	}

	meridiemAm := map[string]string{
		"12": "00",
	}

	var militaryHour string
	meridiem := s[len(s)-2 : len(s)]

	if meridiem == "PM" {
		militaryHour = meridiemPm[s[0:2]]
	} else if meridiem == "AM" && s[0:2] == "12" {
		militaryHour = meridiemAm[s[0:2]]
	} else {
		militaryHour = s[0:2]
	}

	militaryTime := militaryHour + s[2:len(s)-2]

	fmt.Println(militaryTime)
}

func main() {
	timeConversion(time)
}
