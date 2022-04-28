package dispatcher

import (
	"bufio"
	"fmt"
	"koneksa.pizza/deliverer"
	"reflect"
	"strings"
	"testing"
)

var (
	path            = "../PizzaDeliveryInput.txt"
	scannerType     = reflect.TypeOf((*bufio.Scanner)(nil))
	stringSliceType = reflect.TypeOf((*[]string)(nil)).Elem()
)

func TestCanOpenDirectionsFile(t *testing.T) {
	file, err := OpenDirections(path)

	if scannerType != reflect.TypeOf(file) || err != nil {
		t.Fatalf("Failed to open directions. Expected *bufio.Scanner. Got path: %v, file descriptor: %v, error:, %v", path, reflect.TypeOf(file), err)
	}
}

func TestCanGetDirections(t *testing.T) {
	file, err := OpenDirections(path)
	directions, err := GetDirections(file)

	if stringSliceType != reflect.TypeOf(directions) || err != nil {
		t.Fatalf("Failed to get directions. Expected string slice. Got path: %v, type: %v, error: %v", path, reflect.TypeOf(directions), err)
	}
}

func TestCanProcessDirections(t *testing.T) {
	deliverers := []*deliverer.Deliverer{
		&deliverer.Deliverer{
			X:          0,
			Y:          0,
			Deliveries: make(map[string]int),
		},
	}

	directions := strings.Split("vv^<>>", "")
	success, direction := ProcessDirections(directions, deliverers)

	fmt.Printf("%v", (*deliverers[0]).X)

	if !success {
		t.Fatalf("Failed to process directions. Bad direction '%v'", direction)
	}
}
