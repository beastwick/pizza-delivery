package dispatcher

import (
	"bufio"
	"fmt"
	"koneksa.pizza/deliverer"
	"os"
	"strconv"
	"strings"
)

func OpenDirections(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)

	if err != nil {
		//log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
	}

	return scanner, err
}

func GetDirections(scanner *bufio.Scanner) ([]string, error) {
	err := scanner.Err()

	if err != nil {
		//log.Fatal(err)
	}

	return strings.Split(scanner.Text(), ""), err
}

func ProcessDirections(directions []string, deliverers []*deliverer.Deliverer) (bool, string) {
	avail := len(deliverers)
	avail_now := 0

	for _, direction := range directions {
		success := deliverer.ParseDirection(direction, deliverers[avail_now])

		if !success {
			return false, direction
		}

		avail_now += 1

		if avail_now == avail {
			avail_now = 0
		}
	}

	return true, ""
}

func Process(scanner *bufio.Scanner, deliverers []*deliverer.Deliverer) {
	route := 1
	for scanner.Scan() {
		directions, _ := GetDirections(scanner)
		ProcessDirections(directions, deliverers)

		visited, pizzas := deliverer.AggregateDeliveries(deliverers)

		fmt.Printf("Route: %v \n", route)
		fmt.Printf("Directions: %v \n", directions)

		fmt.Printf(strconv.Itoa(len(deliverers))+" deliverer(s) visited %v homes and delivered %v pizzas. \n\n", visited, pizzas)

		ResetDeliverers(deliverers)

		route += 1
	}
}

func ResetDeliverers(deliverers []*deliverer.Deliverer) {
	for _, deliverer := range deliverers {
		(*deliverer).X = 0
		(*deliverer).Y = 0
		(*deliverer).Deliveries = make(map[string]int)
	}
}
