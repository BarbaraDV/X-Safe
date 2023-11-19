package main

import (
	_ "gallery/backend/migrations"
	"gallery/backend/pb"
)

func main() {
	pb.SetServeCommand()
	pb.StartServer()
}
