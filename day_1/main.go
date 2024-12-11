package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func read_input(file_name string) ([]int, []int) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])
		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func calculate_distance(list_locations_1 []int, list_locations_2 []int) int {
	distance := 0

	for i := 0; i < len(list_locations_1); i++ {
		diff := math.Abs(float64(list_locations_1[i] - list_locations_2[i]))
		distance += int(diff)
	}

	return distance
}

func calculate_similarity(list_locations_1 []int, list_locations_2 []int) int {
	similarity := 0

	for i := 0; i < len(list_locations_1); i++ {
		number := list_locations_1[i]

		appearances := 0
		for j := 0; j < len(list_locations_2); j++ {
			if number < list_locations_2[j] {
				break
			}
			if number == list_locations_2[j] {
				appearances++
			}
		}

		similarity += number * appearances
	}

	return similarity
}

func main() {
	fmt.Println("Merry day 1 in code calendar !")

	list_locations_1, list_locations_2 := read_input("input.txt")

	//list_locations_1 := []int{3, 4, 2, 1, 3, 3}
	//list_locations_2 := []int{4, 3, 5, 3, 9, 3}

	//slices.Sort(list_locations_1)
	//slices.Sort(list_locations_2)

	fmt.Printf("The result of the first part is : %d\n", calculate_distance(list_locations_1, list_locations_2))
	fmt.Printf("The result of the second part is : %d\n", calculate_similarity(list_locations_1, list_locations_2))
}
