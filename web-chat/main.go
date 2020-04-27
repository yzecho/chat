package main

import (
	"log"
	"web-chat/route"
)

func main() {
	//s := server.NewServer()
	//log.Fatal(s.Run())
	r := route.Router()
	log.Fatal(r.Run(":8081"))
}
