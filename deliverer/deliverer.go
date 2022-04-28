package deliverer

import (
	"strconv"
)

type Deliverer struct {
	X, Y       int
	Deliveries map[string]int
}

func ParseDirection(direction string, deliverer *Deliverer) bool {
	switch direction {
	case "<":
		applyMovement(-1, 0, deliverer)
	case ">":
		applyMovement(1, 0, deliverer)
	case "^":
		applyMovement(0, 1, deliverer)
	case "v":
		applyMovement(0, -1, deliverer)
	default:
		return false
	}

	return true
}

func applyMovement(x int, y int, deliverer *Deliverer) {
	deliverer.X += x
	deliverer.Y += y
	deliverer.Deliveries[strconv.Itoa(deliverer.X)+","+strconv.Itoa(deliverer.Y)] += 1
}

func AggregateDeliveries(deliverers []*Deliverer) (int, int) {
	homes := make(map[string]int)
	pizzas := 0
	for _, deliverer := range deliverers {
		// All deliverers begin at 0,0 and it counts as a delivery.
		homes["0,0"] += 1
		pizzas += 1

		for coordinates, delivered := range deliverer.Deliveries {
			homes[coordinates] += delivered
			pizzas += delivered
		}
	}
	return len(homes), pizzas
}
