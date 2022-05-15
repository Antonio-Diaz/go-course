package main

import "fmt"

type Room struct {
	Name    string
	Cleaned bool
}

func isCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.Cleaned {
			fmt.Println(room.Name, "is cleaned")
		} else {
			fmt.Println(room.Name, "is not cleaned")
		}
	}
}

func main() {
	rooms := [4]Room{
		{Name: "Living Room", Cleaned: false},
		{Name: "Kitchen", Cleaned: true},
		{Name: "Bedroom", Cleaned: false},
		 {Name: "Bathroom", Cleaned: true},
	}
	isCleanliness(rooms)

	fmt.Println("Performing cleaning...")
	rooms[0].Cleaned = true
	rooms[2].Cleaned = true

	isCleanliness(rooms)
}
