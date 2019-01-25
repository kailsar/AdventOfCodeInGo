package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {  
	data, err := ioutil.ReadFile("aoc1.txt")
	if err != nil {
		fmt.Print(err)
	}
	strdata := string(data)
	splitdata := strings.Split(strdata, "\n")
	total := 0
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
		fmt.Println(total)
	}
}