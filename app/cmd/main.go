package main

import (
	"fance/app/deliveries"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	deliveries.RunRestServer()
}
