package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
)

func main() {  
	data, err := ioutil.ReadFile("aoc1.txt")
	if err != nil {
		fmt.Print(err)
	}
	strdata := string(data)
	splitdata := strings.Split(strdata, "\n")
	total := 0
	var values []int
	found := false
	for found == false {
	
		for _, item := range splitdata {
			value := string(item[1:])
			numvalue, err := strconv.Atoi(value)
			if err != nil { fmt.Print(err)}
			if string(item[0]) == "+" {
				total = total + numvalue
			}
			if string(item[0]) == "-" {
				total = total - numvalue
			}
			if checkValues(total, values) == true {
				fmt.Println(total)
				found = true
				os.Exit(0)
			}
			values = append(values, total)
		}
	}
}
func checkValues(value int, valueSlice []int) bool {
	for _, number := range valueSlice {
		if number == value {
			return true
		}
	}
	return false
}