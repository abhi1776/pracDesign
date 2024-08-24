package service

import (
	"bufio"
	"fmt"
	"os"

	"github.com/abhi1776/pracDesign/parkingLot/domain"
)

func ParkingLotService() {
	var cmd string
	reader := bufio.NewReader(os.Stdin)
	parkingLot := domain.NewParkingLot("single")
	for {
		fmt.Print("Enter command: ")
		cmd, _ = reader.ReadString('\n')
		if cmd == "exit\n" {
			return
		}
		fmt.Println("Command entered:", cmd)
		fmt.Println("output:", parkingLot.ProcessCommand(cmd))
	}
}
