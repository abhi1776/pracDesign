package domain

import (
	"strings"

	"github.com/abhi1776/pracDesign/parkingLot/domain/internal/driver"
	"github.com/abhi1776/pracDesign/parkingLot/domain/internal/singleFloor"
)

//type ProviderMux map[string]

type ParkingLot struct {
	d driver.ParkingLotInterface
}

func NewParkingLot(typ string) (p *ParkingLot) {
	switch typ {
	case "single":
		p = &ParkingLot{d: &singleFloor.SingleFloor{}}
	}
	return p
}

func (p *ParkingLot) CreateParkingLot(slots string) string {
	return p.d.CreateParkingLot(slots)
}

func (p *ParkingLot) Park(number, color string) string {
	return p.d.Park(number, color)
}

func (p *ParkingLot) ProcessCommand(cmd string) string {
	if cmd == "" {
		return "Command is empty"
	}

	split := strings.Split(cmd, " ")
	switch split[0] {
	case "create_parking_lot":
		return p.CreateParkingLot(split[1])
	case "park":
		return p.Park(split[1], split[2])
	case "leave":
		return "leaving"
	case "status":
		return "status"
	case "registration_numbers_for_cars_with_colour":
		return "registration_numbers_for_cars_with_colour"
	case "slot_numbers_for_cars_with_colour":
		return "slot_numbers_for_cars_with_colour"
	case "slot_number_for_registration_number":
		return "slot_number_for_registration_number"
	default:
		return "Invalid command"
	}
}
