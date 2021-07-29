package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/giteshnxtlvl/cook/pkg/cook"
)

var total = 0

//Initializing with empty string, so loops will run for 1st column
var final = []string{""}
var params map[string]string
var pattern []string
var start = time.Now()

// TODO
// -save [keyword] to save the generated permutations

func appendMode(values []string) {
	tmp := []string{}
	till := len(final)
	if len(final) > len(values) {
		till = len(values)
	}
	for i := 0; i < till; i++ {
		tmp = append(tmp, final[i]+values[i])
	}
	final = tmp
}

func permutationMode(values []string) {
	tmp := []string{}
	for _, t := range final {
		for _, v := range values {
			tmp = append(tmp, t+v)
		}
	}
	final = tmp
}

func repititionOp(value string, array *[]string) bool {
	if strings.Count(value, "*") == 1 {
		s := strings.Split(value, "*")
		input := strings.Split(s[0], ",")
		times, err := strconv.Atoi(s[1])
		if err != nil {
			return false
		}

		for _, v := range input {
			*array = append(*array, strings.Repeat(v, times))
		}
		return true
	}
	if strings.Count(value, "*") == 2 {
		s := strings.Split(value, "**")
		input := strings.Split(s[0], ",")
		till, err := strconv.Atoi(s[1])
		if err != nil {
			return false
		}
		for _, v := range input {
			for i := 0; i < till; i++ {
				*array = append(*array, v)
			}
		}
		return true
	}
	return false
}

func checkParam(p string, array *[]string) bool {
	if val, exists := params[p]; exists {
		if cook.PipeInput(val, array) || cook.RawInput(val, array) || repititionOp(val, array) || cook.ParseFunc(val, array) || cook.ParseFile(p, val, array) || checkMethods(val, array) {
			return true
		}

		*array = append(*array, splitValues(val)...)
		return true
	}
	return false
}

func main() {

	parseInput()

	for columnNum, param := range pattern {

		columnValues := []string{}

		for _, p := range splitValues(param) {
			cook.VPrint(fmt.Sprintf("Param: %s \n", p))
			if cook.RawInput(p, &columnValues) || cook.ParseRanges(p, &columnValues) || repititionOp(p, &columnValues) || checkMethods(p, &columnValues) || checkParam(p, &columnValues) || cook.CheckYaml(p, &columnValues) {
				continue
			}
			columnValues = append(columnValues, p)

		}

		cook.VPrint(fmt.Sprintf("%-40s: %s", "Time after getting values", time.Since(start)))

		if mapval, exists := methodMap[columnNum]; exists {
			tmp := []string{}
			applyMethods(columnValues, mapval, &tmp)
			columnValues = tmp
		}

		if !appendMap[columnNum] || columnNum == 0 {
			permutationMode(columnValues)
		} else {
			appendMode(columnValues)
		}

		cook.VPrint(fmt.Sprintf("%-40s: %s", "Time ApplyColumnCases", time.Since(start)))

		if columnNum >= min {
			print()
		}
	}

	cook.VPrint(fmt.Sprintf("%-40s: %s", "Elapsed Time", time.Since(start)))
	cook.VPrint(fmt.Sprintf("%-40s: %d", "Total words generated", total))
}

func init() {
	log.SetFlags(0)
}
