package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./01/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	p1(file)
}

func p1(f *os.File) {
	scanner := bufio.NewScanner(f)
	arr1 := []int{}
	arr2 := []int{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		i1, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			log.Fatal("parsing")
		}

		arr1 = append(arr1, int(i1))

		i2, err := strconv.ParseInt(line[1], 10, 32)
		if err != nil {
			log.Fatal("parsing")
		}

		arr2 = append(arr2, int(i2))
	}

	slices.Sort(arr1)

	slices.Sort(arr2)

	fmt.Println("Vals: ", arr1, " \nVals:2", arr2)

	ret := 0

	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			ret += arr1[i] - arr2[i];
		} else {
			ret += arr2[i] - arr1[i];
		}
	}

	fmt.Println("ANSWER: ", ret);	
}
