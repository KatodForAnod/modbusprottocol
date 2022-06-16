package main

import (
	"fmt"
	"github.com/goburrow/serial"
	"log"
	"modbusprottocol/client"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)

	rtuClient := client.BaseClient{}
	conf := serial.Config{
		Address:  "COM3",
		BaudRate: 115200,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  5 * time.Second,
		RS485:    serial.RS485Config{},
	}

	rtuClient.Connect(conf, 0x01)
	fmt.Println(rtuClient.ReadCoils(0x11, 3))
	fmt.Println(rtuClient.WriteSingleCoil(0x11, client.OFF))
}
