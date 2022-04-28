package deliverer

import (
	//"fmt"
	"os"
	"strconv"
	"testing"
)

var (
	deliveries = make(map[string]int)
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCanParseAndApplyDirectionLeft(t *testing.T) {
	deliverer := Deliverer{
		X:          0,
		Y:          0,
		Deliveries: deliveries,
	}

	success := ParseDirection("<", &deliverer)

	if !success {
		t.Fatalf("Failed to parse '<' character.")
	}

	if deliverer.X != -1 || deliverer.Y != 0 || deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)] != 1 {
		t.Fatalf("Failed to apply left movement. Got X: %v, Y: %v, v: %v", deliverer.X, deliverer.Y, deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)])
	}
}

func TestCanParseAndApplyDirectionRight(t *testing.T) {
	deliverer := Deliverer{
		X:          0,
		Y:          0,
		Deliveries: deliveries,
	}

	success := ParseDirection(">", &deliverer)

	if !success {
		t.Fatalf("Failed to parse '>' character.")
	}

	if deliverer.X != 1 || deliverer.Y != 0 || deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)] != 1 {
		t.Fatalf("Failed to apply right movement. Got X: %v, Y: %v, v: %v", deliverer.X, deliverer.Y, deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)])
	}
}

func TestCanParseAndApplyDirectionUp(t *testing.T) {
	deliverer := Deliverer{
		X:          0,
		Y:          0,
		Deliveries: deliveries,
	}

	success := ParseDirection("^", &deliverer)

	if !success {
		t.Fatalf("Failed to parse '^' character.")
	}

	if deliverer.X != 0 || deliverer.Y != 1 || deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)] != 1 {
		t.Fatalf("Failed to apply up movement. Got X: %v, Y: %v, v: %v", deliverer.X, deliverer.Y, deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)])
	}
}

func TestCanParseAndApplyDirectionDown(t *testing.T) {
	deliverer := Deliverer{
		X:          0,
		Y:          0,
		Deliveries: deliveries,
	}

	success := ParseDirection("v", &deliverer)

	if !success {
		t.Fatalf("Failed to parse 'v' character.")
	}

	if deliverer.X != 0 || deliverer.Y != -1 || deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)] != 1 {
		t.Fatalf("Failed to apply down movement. Got X: %v, Y: %v, v: %v", deliverer.X, deliverer.Y, deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)])
	}
}

func TestGetUniqueDeliveries(t *testing.T) {
	deliveries = make(map[string]int)

	deliverer := Deliverer{
		X:          0,
		Y:          0,
		Deliveries: deliveries,
	}

	ParseDirection("<", &deliverer)
	ParseDirection(">", &deliverer)
	ParseDirection("^", &deliverer)
	ParseDirection("v", &deliverer)

	total := len(deliverer.Deliveries)
	if 3 != total {
		t.Fatalf("failed to get the expected total, got %v", total)
	}
}

func TestAggregateDeliveries(t *testing.T) {
	deliverer_a := Deliverer{
		X:          0,
		Y:          0,
		Deliveries: make(map[string]int),
	}

	ParseDirection("<", &deliverer_a)
	ParseDirection(">", &deliverer_a)
	ParseDirection("^", &deliverer_a)
	ParseDirection("v", &deliverer_a)

	deliverer_b := Deliverer{
		X:          0,
		Y:          0,
		Deliveries: make(map[string]int),
	}

	ParseDirection("<", &deliverer_b)
	ParseDirection(">", &deliverer_b)
	ParseDirection("^", &deliverer_b)
	ParseDirection("v", &deliverer_b)
	ParseDirection("v", &deliverer_b)
	ParseDirection("v", &deliverer_b)

	deliverers := []*Deliverer{&deliverer_a, &deliverer_b}

	homes, _ := AggregateDeliveries(deliverers)

	if 5 != homes {
		t.Fatalf("failed to get the expected total, got %v", homes)
	}
}
