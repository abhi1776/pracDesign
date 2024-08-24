package driver

type ParkingLotInterface interface {
	CreateParkingLot(slots string) string
	Park(number, color string) string
}
