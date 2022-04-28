package main

import (
	"koneksa.pizza/deliverer"
	"koneksa.pizza/dispatcher"
	"log"
	"os"
	//"github.com/pkg/profile"
)

func main() {
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	deliverers := []*deliverer.Deliverer{
		&deliverer.Deliverer{
			X:          0,
			Y:          0,
			Deliveries: make(map[string]int),
		},
	}

	routes, err := dispatcher.OpenDirections(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	dispatcher.Process(routes, deliverers)

	deliverers = []*deliverer.Deliverer{
		&deliverer.Deliverer{
			X:          0,
			Y:          0,
			Deliveries: make(map[string]int),
		},
		&deliverer.Deliverer{
			X:          0,
			Y:          0,
			Deliveries: make(map[string]int),
		},
	}

	routes, err = dispatcher.OpenDirections(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	dispatcher.Process(routes, deliverers)
}
