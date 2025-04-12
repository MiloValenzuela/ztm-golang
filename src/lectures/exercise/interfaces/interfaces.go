//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import (
	"fmt"
)

type Vehicle struct {
	Type  string
	Model string
}

func directToLift(vehicle Vehicle) {
	switch vehicle.Type {
	case "Motorcycle":
		fmt.Printf("Directing Motorcycle '%s' to small lift\n", vehicle.Model)
	case "Car":
		fmt.Printf("Directing Car '%s' to standard lift\n", vehicle.Model)
	case "Truck":
		fmt.Printf("Directing Truck '%s' to large lift\n", vehicle.Model)
	default:
		fmt.Printf("Unknown vehicle type '%s' for model '%s'\n", vehicle.Type, vehicle.Model)
	}
}

func main() {
	vehicle := []Vehicle{
		{Type: "Motorcycle", Model: "Ducati Monster"},
		{Type: "Car", Model: "Toyota Corolla"},
		{Type: "Truck", Model: "Jeeb Wangler"},
		{Type: "Bicycle", Model: "Speedster 5000"},
	}

	for _, v := range vehicle {
		directToLift(v)
	}
}
