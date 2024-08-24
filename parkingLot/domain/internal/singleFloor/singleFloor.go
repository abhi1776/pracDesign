package singleFloor

import (
	"strconv"
	"strings"
)

type Slot struct {
	Id     int
	Number string
	Color  string
}

type SingleFloor struct {
	Slots []Slot
}

func (s *SingleFloor) CreateParkingLot(slots string) string {
	slotsSplit := strings.Split(slots, "\n")
	slotsInt, err := strconv.Atoi(slotsSplit[0])
	if err != nil {
		return "Invalid number of slots"
	}
	s.Slots = make([]Slot, slotsInt)
	return "Created a parking lot with " + slots + " slots"
}

func (s *SingleFloor) Park(number, color string) string {
	for i, slot := range s.Slots {
		if slot.Number == "" {
			s.Slots[i] = Slot{Number: number, Color: color}
			return strconv.Itoa(i + 1)
		}
	}
	return "Sorry, parking lot is full"
}
