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

const (
	smallLift = iota
	standardLift
	largeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcycles string
type Car string
type Truck string

func (m Motorcycles) String() string {
	return fmt.Sprintf("Motorcycles: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motorcycles) PickLift() Lift {
	return smallLift
}

func (c Car) PickLift() Lift {
	return standardLift
}

func (t Truck) PickLift() Lift {
	return largeLift
}

func sendToLift(v LiftPicker) {
	switch v.PickLift() {
	case smallLift:
		fmt.Printf("Send %v to small lift", v)
	case standardLift:
		fmt.Printf("Send %v to standard lift", v)
	case largeLift:
		fmt.Printf("Send %v to large lift", v)
	}
}

func main() {
	vehicles := []LiftPicker{
		Motorcycles("Honda"),
		Car("Ford"),
		Truck("Dodge"),
	}

	for _, v := range vehicles {
		sendToLift(v)
	}
}
