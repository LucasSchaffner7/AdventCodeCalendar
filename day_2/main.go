package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func read_input(file_name string) [][]int {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reports_list [][]int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var report []int
		numbers := strings.Split(scanner.Text(), " ")
		for _, number_str := range numbers {
			number, _ := strconv.Atoi(number_str)
			report = append(report, number)
		}

		reports_list = append(reports_list, report)
	}

	return reports_list
}

func is_report_safe(report []int) bool {
	is_increasing := report[0]-report[1] < 0
	is_safe := true

	for j := 0; j < len(report)-1; j++ {
		nb_1 := report[j]
		nb_2 := report[j+1]

		if nb_1 == nb_2 || math.Abs(float64(nb_1-nb_2)) > 3 {
			is_safe = false
			break
		}

		if (is_increasing && nb_1-nb_2 > 0) || (!is_increasing && nb_1-nb_2 < 0) {
			is_safe = false
			break
		}
	}

	return is_safe
}

func count_safe_reports(reports_list [][]int) int {
	nb_safe_reports := 0

	for i := 0; i < len(reports_list); i++ {
		if is_report_safe(reports_list[i]) {
			nb_safe_reports++
		}
	}

	return nb_safe_reports
}

func count_safe_reports_2(reports_list [][]int) int {
	nb_safe_reports := 0

	for i := 0; i < len(reports_list); i++ {
		report := reports_list[i]
		if is_report_safe(report) {
			nb_safe_reports++
			continue
		}

		for j := 0; j < len(report); j++ {
			modified_report := make([]int, len(report))
			copy(modified_report, report)
			modified_report = append(modified_report[:j], modified_report[j+1:]...)

			if is_report_safe(modified_report) {
				nb_safe_reports++
				break
			}
		}
	}

	return nb_safe_reports
}

func main() {
	fmt.Println("Merry day 2 in code calendar !")

	reports_list := read_input("input.txt")

	fmt.Printf("The number of safe reports is : %d\n", count_safe_reports(reports_list))
	fmt.Printf("The new number of safe reports is : %d\n", count_safe_reports_2(reports_list))
}
