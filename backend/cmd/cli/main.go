package main

import (
	"gallery/backend/pb"
	"log"
)

func main() {
	if err := pb.GenerateMigrations(); err != nil {
		log.Fatal(err)
	}
	pb.StartServer()
}
