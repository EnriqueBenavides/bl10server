package command

import (
	"bl10server/util"
	"log"
)

func ProcessLogin(content []byte) {
	log.Println("IMEI ", util.BytesToInt(content[1:8]))
	log.Println("Sequence number: ", util.BytesToInt(content[13:14]))
}