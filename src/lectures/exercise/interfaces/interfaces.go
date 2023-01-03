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

import "fmt"

type Vehicle interface {
	// interface usually gets -er suffix
	VehicleType() string
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) VehicleType() string {
	return "Motorcycle"
}

func (c Car) VehicleType() string {
	return "Car"
}

func (t Truck) VehicleType() string {
	return "Truck"
}

func allocateLift(v Vehicle) string {
	if _, ok := v.(Motorcycle); ok {
		return "small lift"
	} else if _, ok := v.(Car); ok {
		return "standard lift"
	} else if _, ok := v.(Truck); ok {
		return "large lift"
	}
	return ""
}

func allocateLifts(vehicles []Vehicle) {
	for _, vehicle := range vehicles {
		fmt.Printf("Vehicle --  %v \n", vehicle)
		fmt.Println(vehicle.VehicleType())
		fmt.Println(allocateLift(vehicle))
	}
}

func main() {
	vehicles := []Vehicle{Motorcycle("Kawasaki"), Car("Nissan"), Truck("Isuzu")}
	allocateLifts(vehicles)
}
