package main

import (
	"log"

	"github.com/mesg-foundation/core/x/xsignal"
	"github.com/mesg-foundation/service-google-sheets/sheets"
)

func main() {
	s, err := sheets.New()
	if err != nil {
		log.Fatal(err)
	}

	// start the service.
	go func() {
		log.Println("google-sheets service has been started")

		if err := s.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for interrupt and gracefully shutdown the service.
	<-xsignal.WaitForInterrupt()

	log.Println("shutting down...")

	if err := s.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("shutdown")
}
