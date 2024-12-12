package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read_input(file_name string) []string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func decode_input(lines []string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	regex := regexp.MustCompile(pattern)
	result := 0

	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			nb_1, _ := strconv.Atoi(match[1])
			nb_2, _ := strconv.Atoi(match[2])
			result += nb_1 * nb_2
		}
	}

	return result
}

func decode_input_2(lines []string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	regex := regexp.MustCompile(pattern)
	result := 0
	is_active := true

	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if strings.HasPrefix(match[0], "do") {
				is_active = match[0] == "do()"
				continue
			}

			if !is_active {
				continue
			}

			nb_1, _ := strconv.Atoi(match[1])
			nb_2, _ := strconv.Atoi(match[2])
			result += nb_1 * nb_2
		}
	}

	return result
}

func main() {
	fmt.Println("Merry day 3 in code calendar !")
	lines := read_input("input.txt")
	fmt.Printf("The result for part 1 is : %d\n", decode_input(lines))
	fmt.Printf("The result for part 2 is : %d\n", decode_input_2(lines))
}
